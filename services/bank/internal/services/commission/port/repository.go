package port

import (
	"bank_service/internal/services/commission/domain"
	txDomain "bank_service/internal/services/transaction/domain"
	"context"
	"time"
)

type CommissionRepository interface {
	Create(ctx context.Context, commission *domain.Commission) error
	Update(ctx context.Context, commission *domain.Commission) error
	GetByID(ctx context.Context, id string) (*domain.Commission, error)
	GetByTransactionID(ctx context.Context, txID txDomain.TransactionID) (*domain.Commission, error)
	FindByBusinessIDAndDateRange(ctx context.Context, businessID uint64, startDate, endDate time.Time) ([]*domain.Commission, error)
	FindByStatus(ctx context.Context, status domain.PaymentStatus) ([]*domain.Commission, error)
}
