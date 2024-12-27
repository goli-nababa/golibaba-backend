package port

import (
	"bank_service/internal/services/payment/domain"
	txDomain "bank_service/internal/services/transaction/domain"
	"context"
)

type PaymentRepository interface {
	Create(ctx context.Context, payment *domain.Payment) error
	Update(ctx context.Context, payment *domain.Payment) error
	GetByID(ctx context.Context, id string) (*domain.Payment, error)
	GetByTransactionID(ctx context.Context, txID txDomain.TransactionID) (*domain.Payment, error)
	GetByReferenceID(ctx context.Context, referenceID string) (*domain.Payment, error)
	GetByStatus(ctx context.Context, status domain.PaymentStatus) ([]*domain.Payment, error)
}
