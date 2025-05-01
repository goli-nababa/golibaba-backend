package port

import (
	moneyDomain "bank_service/internal/common/types"
	walletDomain "bank_service/internal/services/wallet/domain"
	"context"
)

type Service interface {
	CreateWallet(ctx context.Context, ownerID uint64, walletType walletDomain.WalletType, currency string) (*walletDomain.Wallet, error)
	GetWallet(ctx context.Context, id walletDomain.WalletID) (*walletDomain.Wallet, error)
	GetWalletsByUser(ctx context.Context, userID uint64) ([]*walletDomain.Wallet, error)
	Credit(ctx context.Context, id walletDomain.WalletID, amount *moneyDomain.Money) error
	Debit(ctx context.Context, id walletDomain.WalletID, amount *moneyDomain.Money) error
	Transfer(ctx context.Context, fromID walletDomain.WalletID, toID walletDomain.WalletID, amount *moneyDomain.Money) error
	UpdateWalletStatus(ctx context.Context, id walletDomain.WalletID, status walletDomain.WalletStatus) error
}
