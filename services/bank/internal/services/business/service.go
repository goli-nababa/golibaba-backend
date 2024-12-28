package business

import (
	"bank_service/internal/services/analytics/domain"
	"context"
	"fmt"
	"time"

	"bank_service/internal/common/types"
	analyticsPort "bank_service/internal/services/analytics/port"
	businessDomain "bank_service/internal/services/business/domain"
	"bank_service/internal/services/business/port"
	commissionPort "bank_service/internal/services/commission/port"
	paymentPort "bank_service/internal/services/payment/port"
	txDomain "bank_service/internal/services/transaction/domain"
	walletDomain "bank_service/internal/services/wallet/domain"
	"bank_service/pkg/cache"
)

type businessService struct {
	businessRepo  port.BusinessRepository
	commissionSvc commissionPort.Service
	paymentSvc    paymentPort.Service
	analyticsSvc  analyticsPort.AnalyticsService
	cache         cache.Provider
}

func NewBusinessService(
	businessRepo port.BusinessRepository,
	commissionSvc commissionPort.Service,
	paymentSvc paymentPort.Service,
	analyticsSvc analyticsPort.AnalyticsService,
	cache cache.Provider,
) port.Service {
	return &businessService{
		businessRepo:  businessRepo,
		commissionSvc: commissionSvc,
		paymentSvc:    paymentSvc,
		analyticsSvc:  analyticsSvc,
		cache:         cache,
	}
}

func (s *businessService) CreateBusinessWallet(ctx context.Context, businessID uint64, businessType businessDomain.BusinessType, currency string) (*businessDomain.BusinessWallet, error) {
	wallet, err := businessDomain.NewBusinessWallet(businessID, businessType, currency)
	if err != nil {
		return nil, fmt.Errorf("failed to create business wallet: %w", err)
	}

	if err := s.businessRepo.Save(ctx, wallet); err != nil {
		return nil, fmt.Errorf("failed to save business wallet: %w", err)
	}

	return wallet, nil
}

func (s *businessService) GetBusinessWallet(ctx context.Context, walletID walletDomain.WalletID) (*businessDomain.BusinessWallet, error) {

	wallet, err := s.businessRepo.GetByWalletID(ctx, walletID)
	if err != nil {
		return nil, fmt.Errorf("failed to get business wallet: %w", err)
	}

	return wallet, nil
}

func (s *businessService) UpdateBusinessWallet(ctx context.Context, wallet *businessDomain.BusinessWallet) error {
	if err := s.businessRepo.Update(ctx, wallet); err != nil {
		return fmt.Errorf("failed to update business wallet: %w", err)
	}

	return nil
}

func (s *businessService) ProcessCommission(ctx context.Context, tx *txDomain.Transaction) error {
	business, err := s.GetBusinessWallet(ctx, tx.ToWalletID)
	if err != nil {
		return fmt.Errorf("failed to get business wallet: %w", err)
	}

	commission, err := s.commissionSvc.CalculateCommission(ctx, tx, business.BusinessType)
	if err != nil {
		return fmt.Errorf("failed to calculate commission: %w", err)
	}

	if err := s.commissionSvc.ProcessCommission(ctx, commission.ID); err != nil {
		return fmt.Errorf("failed to process commission: %w", err)
	}

	return nil
}

func (s *businessService) RequestPayout(ctx context.Context, walletID walletDomain.WalletID, amount *types.Money) error {
	wallet, err := s.GetBusinessWallet(ctx, walletID)
	if err != nil {
		return err
	}

	if wallet.MinimumPayout != nil && amount.LessThan(wallet.MinimumPayout) {
		return fmt.Errorf("payout amount is below minimum threshold of %s", wallet.MinimumPayout)
	}

	tx, err := txDomain.NewTransaction(
		wallet.ID,
		businessDomain.CENTRAL_WALLET_ID,
		amount,
		txDomain.TransactionTypeWithdrawal,
		"Business payout",
	)
	if err != nil {
		return err
	}

	_, err = s.paymentSvc.ProcessPayment(ctx, tx)
	return err
}

func (s *businessService) SetPayoutSchedule(ctx context.Context, walletID walletDomain.WalletID, schedule string) error {
	wallet, err := s.GetBusinessWallet(ctx, walletID)
	if err != nil {
		return err
	}

	wallet.PayoutSchedule = schedule
	return s.UpdateBusinessWallet(ctx, wallet)
}

func (s *businessService) SetMinimumPayout(ctx context.Context, walletID walletDomain.WalletID, amount *types.Money) error {
	wallet, err := s.GetBusinessWallet(ctx, walletID)
	if err != nil {
		return err
	}

	wallet.MinimumPayout = amount
	return s.UpdateBusinessWallet(ctx, wallet)
}

func (s *businessService) UpdateBankInfo(ctx context.Context, walletID walletDomain.WalletID, bankInfo *businessDomain.BankAccountInfo) error {
	wallet, err := s.GetBusinessWallet(ctx, walletID)
	if err != nil {
		return err
	}

	wallet.BankInfo = bankInfo
	return s.UpdateBusinessWallet(ctx, wallet)
}

func (s *businessService) GetBankInfo(ctx context.Context, walletID walletDomain.WalletID) (*businessDomain.BankAccountInfo, error) {
	wallet, err := s.GetBusinessWallet(ctx, walletID)
	if err != nil {
		return nil, err
	}
	return wallet.BankInfo, nil
}

func (s *businessService) SetCommissionRate(ctx context.Context, walletID walletDomain.WalletID, rate float64) error {
	if rate < 0 || rate > 1 {
		return fmt.Errorf("invalid commission rate: must be between 0 and 1")
	}

	wallet, err := s.GetBusinessWallet(ctx, walletID)
	if err != nil {
		return err
	}

	wallet.CommissionRate = rate
	return s.UpdateBusinessWallet(ctx, wallet)
}

func (s *businessService) GetCommissionRate(ctx context.Context, walletID walletDomain.WalletID) (float64, error) {
	wallet, err := s.GetBusinessWallet(ctx, walletID)
	if err != nil {
		return 0, err
	}
	return wallet.CommissionRate, nil
}

func (s *businessService) GetBusinessStats(ctx context.Context, businessID uint64, startDate, endDate time.Time) (*domain.BusinessStats, error) {
	return s.analyticsSvc.GenerateBusinessStats(ctx, businessID, startDate, endDate)
}

func (s *businessService) GetCommissionHistory(ctx context.Context, businessID uint64, startDate, endDate time.Time) ([]*domain.CommissionEntry, error) {
	return s.analyticsSvc.GetCommissionHistory(ctx, businessID, startDate, endDate)
}

func (s *businessService) GetPayoutHistory(ctx context.Context, walletID walletDomain.WalletID, startDate, endDate time.Time) ([]*businessDomain.PayoutEntry, error) {
	return s.analyticsSvc.GetPayoutHistory(ctx, walletID, startDate, endDate)
}

func (s *businessService) ProcessAutomaticPayouts(ctx context.Context) error {
	businesses, err := s.businessRepo.ListBusinessWallets(ctx, &businessDomain.BusinessWalletFilter{})
	if err != nil {
		return err
	}

	for _, business := range businesses {
		if business.PayoutSchedule == "" {
			continue
		}

		if err := s.processBusinessPayout(ctx, business); err != nil {
			continue
		}
	}

	return nil
}

func (s *businessService) processBusinessPayout(ctx context.Context, wallet *businessDomain.BusinessWallet) error {
	if wallet.LastPayoutDate != nil {
		nextPayout := calculateNextPayoutDate(wallet.LastPayoutDate, wallet.PayoutSchedule)
		if time.Now().Before(nextPayout) {
			return nil
		}
	}

	balance := wallet.Balance

	if wallet.MinimumPayout != nil && balance.LessThan(wallet.MinimumPayout) {
		return nil
	}

	if err := s.RequestPayout(ctx, wallet.ID, balance); err != nil {
		return err
	}

	now := time.Now()
	wallet.LastPayoutDate = &now
	return s.UpdateBusinessWallet(ctx, wallet)
}

func calculateNextPayoutDate(lastPayout *time.Time, schedule string) time.Time {
	switch schedule {
	case "daily":
		return lastPayout.Add(24 * time.Hour)
	case "weekly":
		return lastPayout.Add(7 * 24 * time.Hour)
	case "monthly":
		return lastPayout.AddDate(0, 1, 0)
	default:
		return time.Now()
	}
}
