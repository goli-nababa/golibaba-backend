package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	walletDomain "bank_service/internal/services/wallet/domain"
	walletPort "bank_service/internal/services/wallet/port"
	"bank_service/pkg/cache"
	"bank_service/pkg/logging"
)

const (
	walletCachePrefix = "wallet"
	defaultTTL        = 24 * time.Hour
)

type RedisWalletCache struct {
	provider cache.Provider
	logger   logging.Logger
}

func NewWalletCache(provider cache.Provider, logger logging.Logger) walletPort.WalletCache {
	return &RedisWalletCache{
		provider: provider,
		logger:   logger,
	}
}

func (c *RedisWalletCache) Get(ctx context.Context, id walletDomain.WalletID) (*walletDomain.Wallet, error) {
	key := c.makeKey(id)

	data, err := c.provider.Get(ctx, key)
	if err != nil {
		if err == cache.ErrCacheMiss {
			c.logger.Debug(logging.Redis, "Cache", "Cache miss for wallet", map[logging.ExtraKey]interface{}{
				"wallet_id": id,
			})
			return nil, nil
		}
		c.logger.Error(logging.Redis, "Cache", "Failed to get wallet from cache", map[logging.ExtraKey]interface{}{
			"wallet_id": id,
			"error":     err.Error(),
		})
		return nil, fmt.Errorf("failed to get wallet from cache: %w", err)
	}

	var wallet walletDomain.Wallet
	if err := json.Unmarshal(data, &wallet); err != nil {
		c.logger.Error(logging.Redis, "Cache", "Failed to unmarshal wallet data", map[logging.ExtraKey]interface{}{
			"wallet_id": id,
			"error":     err.Error(),
		})
		return nil, fmt.Errorf("failed to unmarshal wallet data: %w", err)
	}

	c.logger.Debug(logging.Redis, "Cache", "Cache hit for wallet", map[logging.ExtraKey]interface{}{
		"wallet_id": id,
	})
	return &wallet, nil
}

func (c *RedisWalletCache) Set(ctx context.Context, wallet *walletDomain.Wallet, ttl time.Duration) error {
	if wallet == nil {
		return fmt.Errorf("wallet cannot be nil")
	}

	key := c.makeKey(wallet.ID)

	data, err := json.Marshal(wallet)
	if err != nil {
		c.logger.Error(logging.Redis, "Cache", "Failed to marshal wallet data", map[logging.ExtraKey]interface{}{
			"wallet_id": wallet.ID,
			"error":     err.Error(),
		})
		return fmt.Errorf("failed to marshal wallet data: %w", err)
	}

	if ttl <= 0 {
		ttl = defaultTTL
	}

	if err := c.provider.Set(ctx, key, ttl, data); err != nil {
		c.logger.Error(logging.Redis, "Cache", "Failed to set wallet in cache", map[logging.ExtraKey]interface{}{
			"wallet_id": wallet.ID,
			"error":     err.Error(),
		})
		return fmt.Errorf("failed to set wallet in cache: %w", err)
	}

	c.logger.Debug(logging.Redis, "Cache", "Successfully cached wallet", map[logging.ExtraKey]interface{}{
		"wallet_id": wallet.ID,
		"ttl":       ttl.String(),
	})
	return nil
}

func (c *RedisWalletCache) Delete(ctx context.Context, id walletDomain.WalletID) error {
	key := c.makeKey(id)

	if err := c.provider.Del(ctx, key); err != nil {
		c.logger.Error(logging.Redis, "Cache", "Failed to delete wallet from cache", map[logging.ExtraKey]interface{}{
			"wallet_id": id,
			"error":     err.Error(),
		})
		return fmt.Errorf("failed to delete wallet from cache: %w", err)
	}

	c.logger.Debug(logging.Redis, "Cache", "Successfully deleted wallet from cache", map[logging.ExtraKey]interface{}{
		"wallet_id": id,
	})
	return nil
}

func (c *RedisWalletCache) makeKey(id walletDomain.WalletID) string {
	return fmt.Sprintf("%s:%s", walletCachePrefix, id)
}
