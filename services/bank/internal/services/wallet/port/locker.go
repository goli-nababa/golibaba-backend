package port

import (
	walletDomain "bank_service/internal/services/wallet/domain"
	"context"
	"time"
)

type WalletLocker interface {
	AcquireLock(ctx context.Context, id walletDomain.WalletID, ttl time.Duration) (bool, error)
	ReleaseLock(ctx context.Context, id walletDomain.WalletID) error
}
