package cache

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
)

type redisCacheAdapter struct {
	client *redis.Client
}

func NewRedisProvider(redisAddr string, username, password string, db int) Provider {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Username: username,
		Password: password,
		DB:       db,
	})

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
			return nil, ErrCacheMiss
		}
		return nil, err
	}

	return data, nil
}

func (r *redisCacheAdapter) Del(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

func (r *redisCacheAdapter) Exists(ctx context.Context, key string) (bool, error) {
	result, err := r.client.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return result > 0, nil
}
