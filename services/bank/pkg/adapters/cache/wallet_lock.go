package cache

import (
	walletDomain "bank_service/internal/services/wallet/domain"
	walletPort "bank_service/internal/services/wallet/port"
	"bank_service/pkg/cache"
	"context"
	"fmt"
	"time"
)

type RedisWalletLocker struct {
	provider cache.Provider
}

func NewRedisWalletLocker(provider cache.Provider) walletPort.WalletLocker {
	return &RedisWalletLocker{
		provider: provider,
	}
}

func (l *RedisWalletLocker) AcquireLock(ctx context.Context, id walletDomain.WalletID, ttl time.Duration) (bool, error) {
	lockKey := fmt.Sprintf("wallet_lock:%s", id)
	return l.provider.GetLock(ctx, lockKey, ttl)
}

func (l *RedisWalletLocker) ReleaseLock(ctx context.Context, id walletDomain.WalletID) error {
	lockKey := fmt.Sprintf("wallet_lock:%s", id)
	return l.provider.ReleaseLock(ctx, lockKey)
}
