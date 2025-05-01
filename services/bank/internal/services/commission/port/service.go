package port

import (
	businessWalletDomain "bank_service/internal/services/business/domain"
	"bank_service/internal/services/commission/domain"
	txDomain "bank_service/internal/services/transaction/domain"
	"context"
)

type Service interface {
	CalculateCommission(ctx context.Context, tx *txDomain.Transaction, businessType businessWalletDomain.BusinessType) (*domain.Commission, error)
	ProcessCommission(ctx context.Context, commissionID string) error
	GetCommission(ctx context.Context, id string) (*domain.Commission, error)

	GetPendingCommissions(ctx context.Context) ([]*domain.Commission, error)
	GetFailedCommissions(ctx context.Context) ([]*domain.Commission, error)
	RetryFailedCommissions(ctx context.Context) error
}

type CommissionRateProvider interface {
	GetRate(ctx context.Context, businessType businessWalletDomain.BusinessType) (float64, error)
}
