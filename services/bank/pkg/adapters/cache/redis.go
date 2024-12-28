package cache

import (
	walletDomain "bank_service/internal/services/wallet/domain"
	"context"
	"errors"
	"fmt"
	"time"

	c "bank_service/pkg/cache"

	"github.com/redis/go-redis/v9"
)

type redisCacheAdapter struct {
	client *redis.Client
}

func NewRedisProvider(redisAddr string) c.Provider {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "",
		DB:       0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}

	return &redisCacheAdapter{
		client: rdb,
	}
}

func (r *redisCacheAdapter) Set(ctx context.Context, key string, ttl time.Duration, data []byte) error {
	return r.client.Set(ctx, key, data, ttl).Err()
}

func (r *redisCacheAdapter) Get(ctx context.Context, key string) ([]byte, error) {
	data, err := r.client.Get(ctx, key).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, c.ErrCacheMiss
		}
		return nil, err
	}

	return data, nil
}

func (r *redisCacheAdapter) Del(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

func (r *redisCacheAdapter) GetLock(ctx context.Context, key string, ttl time.Duration) (bool, error) {
	ok, err := r.client.SetNX(ctx, key, "locked", ttl).Result()
	if err != nil {
		return false, fmt.Errorf("failed to acquire lock for key %s: %w", key, err)
	}
	return ok, nil
}

func (r *redisCacheAdapter) ReleaseLock(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

func (r *redisCacheAdapter) Delete(ctx context.Context, id walletDomain.WalletID) error {
	//TODO implement me
	panic("implement me")
}
