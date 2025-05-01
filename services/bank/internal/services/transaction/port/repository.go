package port

import (
	"bank_service/internal/services/transaction/domain"
	walletDomain "bank_service/internal/services/wallet/domain"
	"context"
	"time"
)

type TransactionRepository interface {
	Create(ctx context.Context, tx *domain.Transaction) error
	Update(ctx context.Context, tx *domain.Transaction) error
	FindByID(ctx context.Context, id domain.TransactionID) (*domain.Transaction, error)
	FindByWallet(ctx context.Context, walletID walletDomain.WalletID, filter *TransactionFilter) ([]*domain.Transaction, error)
	GetVersion(ctx context.Context, id domain.TransactionID) (int, error)
	FindByBusinessIDAndDateRange(ctx context.Context, businessID uint64, startDate, endDate time.Time) ([]*domain.Transaction, error)
	FindByReferenceID(ctx context.Context, referenceID string) (*domain.Transaction, error)
}

type LockProvider interface {
	AcquireLock(ctx context.Context, key string) (bool, error)
	ReleaseLock(ctx context.Context, key string) error
}

type TransactionPublisher interface {
	PublishTransactionCreated(ctx context.Context, tx *domain.Transaction) error
	PublishTransactionStatusChanged(ctx context.Context, tx *domain.Transaction, oldStatus domain.TransactionStatus) error
}
