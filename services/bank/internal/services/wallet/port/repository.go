package port

import (
	"bank_service/internal/services/wallet/domain"
	"context"
)

type WalletRepository interface {
	Save(ctx context.Context, wallet *domain.Wallet) error
	Update(ctx context.Context, wallet *domain.Wallet) error
	FindByID(ctx context.Context, id domain.WalletID) (*domain.Wallet, error)
	FindByUserID(ctx context.Context, userID uint64) ([]*domain.Wallet, error)
	GetVersion(ctx context.Context, id domain.WalletID) (int, error)
}

type LockProvider interface {
	AcquireLock(ctx context.Context, key string) (bool, error)
	ReleaseLock(ctx context.Context, key string) error
}
