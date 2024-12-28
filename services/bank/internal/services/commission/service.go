package service

import (
	businessDomain "bank_service/internal/services/business/domain"
	"bank_service/internal/services/commission/domain"
	"bank_service/internal/services/commission/port"
	txDomain "bank_service/internal/services/transaction/domain"
	walletPort "bank_service/internal/services/wallet/port"
	"bank_service/pkg/cache"
	"context"
	"errors"
	"fmt"
	"time"
)

type commissionService struct {
	repo         port.CommissionRepository
	rateProvider port.CommissionRateProvider
	walletSvc    walletPort.Service
	cache        cache.Provider
}

func NewCommissionService(
	repo port.CommissionRepository,
	rateProvider port.CommissionRateProvider,
	walletSvc walletPort.Service,
	cache cache.Provider,
) port.Service {
	return &commissionService{
		repo:         repo,
		rateProvider: rateProvider,
		walletSvc:    walletSvc,
		cache:        cache,
	}
}

func (s *commissionService) CalculateCommission(ctx context.Context, tx *txDomain.Transaction, businessType businessDomain.BusinessType) (*domain.Commission, error) {

	rate, err := s.rateProvider.GetRate(ctx, businessType)
	if err != nil {
		return nil, fmt.Errorf("failed to get commission rate: %w", err)
	}

	commission, err := domain.NewCommission(tx, rate, businessType)
	if err != nil {
		return nil, fmt.Errorf("failed to create commission: %w", err)
	}

	if err := s.repo.Create(ctx, commission); err != nil {
		return nil, fmt.Errorf("failed to save commission: %w", err)
	}

	return commission, nil
}

func (s *commissionService) ProcessCommission(ctx context.Context, commissionID string) error {
	lockKey := fmt.Sprintf("commission_lock:%s", commissionID)
	locked, err := s.cache.GetLock(ctx, lockKey, time.Second*30)
	if err != nil {
		return fmt.Errorf("failed to acquire lock: %w", err)
	}
	if !locked {
		return errors.New("concurrent commission processing")
	}
	defer s.cache.ReleaseLock(ctx, lockKey)

	commission, err := s.repo.GetByID(ctx, commissionID)
	if err != nil {
		return fmt.Errorf("failed to get commission: %w", err)
	}

	if commission.Status != domain.PaymentStatusPending {
		return errors.New("commission is not in pending status")
	}

	err = s.walletSvc.Transfer(
		ctx,
		commission.RecipientID,
		businessDomain.CENTRAL_WALLET_ID,
		commission.Amount,
	)
	if err != nil {
		commission.Status = domain.PaymentStatusFailed
		s.repo.Update(ctx, commission)
		return fmt.Errorf("failed to transfer commission: %w", err)
	}

	commission.Status = domain.PaymentStatusProcessed
	now := time.Now()
	commission.PaidAt = &now
	if err := s.repo.Update(ctx, commission); err != nil {
		return fmt.Errorf("failed to update commission: %w", err)
	}

	return nil
}

func (s *commissionService) GetCommission(ctx context.Context, id string) (*domain.Commission, error) {
	commission, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get commission: %w", err)
	}
	return commission, nil
}

func (s *commissionService) GetPendingCommissions(ctx context.Context) ([]*domain.Commission, error) {
	return s.repo.FindByStatus(ctx, domain.PaymentStatusPending)
}

func (s *commissionService) GetFailedCommissions(ctx context.Context) ([]*domain.Commission, error) {
	return s.repo.FindByStatus(ctx, domain.PaymentStatusFailed)
}

func (s *commissionService) RetryFailedCommissions(ctx context.Context) error {
	failedCommissions, err := s.GetFailedCommissions(ctx)
	if err != nil {
		return fmt.Errorf("failed to get failed commissions: %w", err)
	}

	for _, commission := range failedCommissions {
		commission.Status = domain.PaymentStatusPending
		if err := s.repo.Update(ctx, commission); err != nil {
			continue
		}

		if err := s.ProcessCommission(ctx, commission.ID); err != nil {
			continue
		}
	}

	return nil
}
