package cache

import (
	transactionPort "bank_service/internal/services/transaction/port"
	"bank_service/pkg/cache"
	"context"
	"time"
)

const (
	defaultLockTTL = 30 * time.Second
)

type RedisTransactionLocker struct {
	provider cache.Provider
}

func NewTransactionLocker(provider cache.Provider) transactionPort.LockProvider {
	return &RedisTransactionLocker{
		provider: provider,
	}
}

func (l *RedisTransactionLocker) AcquireLock(ctx context.Context, key string) (bool, error) {

	return l.provider.GetLock(ctx, key, defaultLockTTL)
}

func (l *RedisTransactionLocker) ReleaseLock(ctx context.Context, key string) error {
	return l.provider.ReleaseLock(ctx, key)
}
