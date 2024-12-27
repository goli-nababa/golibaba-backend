package service

import (
	moneyDomain "bank_service/internal/common/types"
	walletDomain "bank_service/internal/services/wallet/domain"
	"bank_service/internal/services/wallet/port"
	"context"
	"fmt"
	"time"
)

type walletService struct {
	repo     port.WalletRepository
	cache    port.WalletCache
	locker   port.WalletLocker
	eventPub port.WalletEventPublisher
}

func NewWalletService(
	repo port.WalletRepository,
	cache port.WalletCache,
	locker port.WalletLocker,
	eventPub port.WalletEventPublisher,
) port.Service {
	return &walletService{
		repo:     repo,
		cache:    cache,
		locker:   locker,
		eventPub: eventPub,
	}
}

func (s *walletService) CreateWallet(ctx context.Context, ownerID uint64, walletType walletDomain.WalletType, currency string) (*walletDomain.Wallet, error) {
	wallet, err := walletDomain.NewWallet(ownerID, walletType, currency)
	if err != nil {
		return nil, fmt.Errorf("failed to create wallet: %w", err)
	}

	if err := s.repo.Save(ctx, wallet); err != nil {
		return nil, fmt.Errorf("failed to save wallet: %w", err)
	}

	if err := s.cache.Set(ctx, wallet, 24*time.Hour); err != nil {
		// Log error but don't fail
	}

	event := walletDomain.WalletCreatedEvent{Wallet: wallet}
	if err := s.eventPub.PublishEvent(ctx, event); err != nil {
		// Log error but don't fail
	}

	return wallet, nil
}

func (s *walletService) GetWallet(ctx context.Context, id walletDomain.WalletID) (*walletDomain.Wallet, error) {
	wallet, err := s.cache.Get(ctx, id)
	if err == nil {
		return wallet, nil
	}

	wallet, err = s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get wallet: %w", err)
	}

	if err := s.cache.Set(ctx, wallet, 24*time.Hour); err != nil {
		// Log error but don't fail
	}

	return wallet, nil
}

func (s *walletService) GetWalletsByUser(ctx context.Context, userID uint64) ([]*walletDomain.Wallet, error) {
	return s.repo.FindByUserID(ctx, userID)
}

func (s *walletService) Credit(ctx context.Context, id walletDomain.WalletID, amount *moneyDomain.Money) error {
	locked, err := s.locker.AcquireLock(ctx, id, 30*time.Second)
	if err != nil {
		return fmt.Errorf("failed to acquire lock: %w", err)
	}
	if !locked {
		return walletDomain.ErrConcurrentOperation
	}
	defer s.locker.ReleaseLock(ctx, id)

	wallet, err := s.getWalletForUpdate(ctx, id)
	if err != nil {
		return err
	}

	if err := wallet.Credit(amount); err != nil {
		return fmt.Errorf("failed to credit wallet: %w", err)
	}

	if err := s.saveWalletChanges(ctx, wallet); err != nil {
		return err
	}

	event := walletDomain.WalletBalanceChangedEvent{
		WalletID:     id,
		ChangeAmount: amount,
		ChangeType:   "credit",
	}
	if err := s.eventPub.PublishEvent(ctx, event); err != nil {
		// Log error but don't fail
	}

	return nil
}

func (s *walletService) Debit(ctx context.Context, id walletDomain.WalletID, amount *moneyDomain.Money) error {
	locked, err := s.locker.AcquireLock(ctx, id, 30*time.Second)
	if err != nil {
		return fmt.Errorf("failed to acquire lock: %w", err)
	}
	if !locked {
		return walletDomain.ErrConcurrentOperation
	}
	defer s.locker.ReleaseLock(ctx, id)

	wallet, err := s.getWalletForUpdate(ctx, id)
	if err != nil {
		return err
	}

	if err := wallet.Debit(amount); err != nil {
		return fmt.Errorf("failed to debit wallet: %w", err)
	}

	if err := s.saveWalletChanges(ctx, wallet); err != nil {
		return err
	}

	event := walletDomain.WalletBalanceChangedEvent{
		WalletID:     id,
		ChangeAmount: amount,
		ChangeType:   "debit",
	}
	if err := s.eventPub.PublishEvent(ctx, event); err != nil {
		// Log error but don't fail
	}

	return nil
}

func (s *walletService) Transfer(ctx context.Context, fromID walletDomain.WalletID, toID walletDomain.WalletID, amount *moneyDomain.Money) error {
	lockID1, lockID2 := fromID, toID
	if string(toID) < string(fromID) {
		lockID1, lockID2 = toID, fromID
	}

	locked1, err := s.locker.AcquireLock(ctx, lockID1, 30*time.Second)
	if err != nil || !locked1 {
		return fmt.Errorf("failed to acquire first lock: %w", err)
	}
	defer s.locker.ReleaseLock(ctx, lockID1)

	locked2, err := s.locker.AcquireLock(ctx, lockID2, 30*time.Second)
	if err != nil || !locked2 {
		return fmt.Errorf("failed to acquire second lock: %w", err)
	}
	defer s.locker.ReleaseLock(ctx, lockID2)

	sourceWallet, err := s.getWalletForUpdate(ctx, fromID)
	if err != nil {
		return fmt.Errorf("failed to get source wallet: %w", err)
	}

	destWallet, err := s.getWalletForUpdate(ctx, toID)
	if err != nil {
		return fmt.Errorf("failed to get destination wallet: %w", err)
	}

	if err := sourceWallet.Debit(amount); err != nil {
		return fmt.Errorf("failed to debit source wallet: %w", err)
	}

	if err := destWallet.Credit(amount); err != nil {
		if rollbackErr := sourceWallet.Credit(amount); rollbackErr != nil {
			return fmt.Errorf("critical error: debit succeeded but credit and rollback failed. Original error: %v, Rollback error: %v",
				err, rollbackErr)
		}
		return fmt.Errorf("failed to credit destination wallet: %w", err)
	}

	if err := s.saveWalletChanges(ctx, sourceWallet); err != nil {
		return err
	}

	if err := s.saveWalletChanges(ctx, destWallet); err != nil {
		return fmt.Errorf("critical error: source wallet updated but destination update failed: %w", err)
	}

	if err := s.eventPub.PublishEvent(ctx, walletDomain.WalletBalanceChangedEvent{
		WalletID:     fromID,
		ChangeAmount: amount,
		ChangeType:   "debit",
	}); err != nil {
		// Log error but don't fail
	}

	if err := s.eventPub.PublishEvent(ctx, walletDomain.WalletBalanceChangedEvent{
		WalletID:     toID,
		ChangeAmount: amount,
		ChangeType:   "credit",
	}); err != nil {
		// Log error but don't fail
	}

	return nil
}

func (s *walletService) UpdateWalletStatus(ctx context.Context, id walletDomain.WalletID, status walletDomain.WalletStatus) error {
	wallet, err := s.getWalletForUpdate(ctx, id)
	if err != nil {
		return err
	}

	wallet.Status = status
	return s.saveWalletChanges(ctx, wallet)
}

func (s *walletService) getWalletForUpdate(ctx context.Context, id walletDomain.WalletID) (*walletDomain.Wallet, error) {
	wallet, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get wallet: %w", err)
	}
	if wallet.DailyTransactions == nil {
		wallet.DailyTransactions = make(map[time.Time]*moneyDomain.Money)
	}
	if wallet.MonthlyTransactions == nil {
		wallet.MonthlyTransactions = make(map[time.Time]*moneyDomain.Money)
	}
	return wallet, nil
}

func (s *walletService) saveWalletChanges(ctx context.Context, wallet *walletDomain.Wallet) error {
	currentVersion, err := s.repo.GetVersion(ctx, wallet.ID)
	if err != nil {
		return fmt.Errorf("failed to get current version: %w", err)
	}

	if currentVersion != wallet.Version {
		return walletDomain.ErrConcurrentModification
	}

	if err := s.repo.Update(ctx, wallet); err != nil {
		return fmt.Errorf("failed to update wallet: %w", err)
	}

	if err := s.cache.Set(ctx, wallet, 24*time.Hour); err != nil {
		// Log error but don't fail
	}

	return nil
}
