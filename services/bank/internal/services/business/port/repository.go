package port

import (
	"bank_service/internal/services/business/domain"
	walletDomain "bank_service/internal/services/wallet/domain"
	"context"
)

type BusinessRepository interface {
	Save(ctx context.Context, wallet *domain.BusinessWallet) error
	Update(ctx context.Context, wallet *domain.BusinessWallet) error
	GetByWalletID(ctx context.Context, walletID walletDomain.WalletID) (*domain.BusinessWallet, error)
	GetByBusinessID(ctx context.Context, businessID uint64) (*domain.BusinessWallet, error)
	GetCentralWallet(ctx context.Context) (*walletDomain.Wallet, error)
	Delete(ctx context.Context, walletID walletDomain.WalletID) error

	ListBusinessWallets(ctx context.Context, filter *domain.BusinessWalletFilter) ([]*domain.BusinessWallet, error)
	GetVersion(ctx context.Context, walletID walletDomain.WalletID) (int, error)
}
