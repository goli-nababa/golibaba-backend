package service

import (
	"context"
	"fmt"

	"bank_service/internal/services/transaction/domain"
	"bank_service/internal/services/transaction/port"
)

type transactionService struct {
	repo      port.TransactionRepository
	locker    port.LockProvider
	publisher port.TransactionPublisher
}

func NewTransactionService(repo port.TransactionRepository, locker port.LockProvider, publisher port.TransactionPublisher) port.TransactionService {
	return &transactionService{
		repo:      repo,
		locker:    locker,
		publisher: publisher,
	}
}

func (s *transactionService) CreateTransaction(ctx context.Context, req *port.TransactionRequest) (*domain.Transaction, error) {
	tx, err := domain.NewTransaction(req.FromWalletID, req.ToWalletID, req.Amount, req.Type, req.Description)
	if err != nil {
		return nil, err
	}
	if err := s.repo.Create(ctx, tx); err != nil {
		return nil, fmt.Errorf("failed to create transaction: %w", err)
	}
	if err := s.publisher.PublishTransactionCreated(ctx, tx); err != nil {
		// Log error but don't fail
	}
	return tx, nil
}

func (s *transactionService) ProcessTransaction(ctx context.Context, id domain.TransactionID) error {
	tx, err := s.getTransactionForUpdate(ctx, id)
	if err != nil {
		return err
	}

	if err := tx.Process(); err != nil {
		return err
	}

	if err := tx.Complete(); err != nil {
		return err
	}
	if err := s.repo.Update(ctx, tx); err != nil {
		return err
	}
	if err := s.publisher.PublishTransactionStatusChanged(ctx, tx, domain.TransactionStatusProcessing); err != nil {
		// Log error but don't fail
	}
	return nil
}

func (s *transactionService) CancelTransaction(ctx context.Context, id domain.TransactionID, reason string) error {
	tx, err := s.getTransactionForUpdate(ctx, id)
	if err != nil {
		return err
	}
	if err := tx.Cancel(reason); err != nil {
		return err
	}
	if err := s.repo.Update(ctx, tx); err != nil {
		return err
	}
	if err := s.publisher.PublishTransactionStatusChanged(ctx, tx, domain.TransactionStatusPending); err != nil {
		// Log error but don't fail
	}
	return nil
}

func (s *transactionService) GetTransaction(ctx context.Context, id domain.TransactionID) (*domain.Transaction, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *transactionService) ListTransactions(ctx context.Context, filter *port.TransactionFilter) ([]*domain.Transaction, error) {
	return s.repo.FindByWallet(ctx, filter.WalletID, (*port.TransactionFilter)(filter))
}

func (s *transactionService) getTransactionForUpdate(ctx context.Context, id domain.TransactionID) (*domain.Transaction, error) {
	lockKey := fmt.Sprintf("transaction:%s", id)
	locked, err := s.locker.AcquireLock(ctx, lockKey)
	if err != nil {
		return nil, fmt.Errorf("failed to acquire lock: %w", err)
	}
	if !locked {
		return nil, fmt.Errorf("transaction is locked by another process")
	}
	defer s.locker.ReleaseLock(ctx, lockKey)

	tx, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	version, err := s.repo.GetVersion(ctx, id)
	if err != nil {
		return nil, err
	}
	if tx.Version != version {
		return nil, fmt.Errorf("transaction has been modified")
	}
	return tx, nil
}
