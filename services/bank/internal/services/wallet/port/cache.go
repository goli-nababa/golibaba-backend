package port

import (
	walletDomain "bank_service/internal/services/wallet/domain"
	"context"
	"time"
)

type WalletCache interface {
	Get(ctx context.Context, id walletDomain.WalletID) (*walletDomain.Wallet, error)
	Set(ctx context.Context, wallet *walletDomain.Wallet, ttl time.Duration) error
	Delete(ctx context.Context, id walletDomain.WalletID) error
}
