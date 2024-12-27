package repository

import (
	"bank_service/pkg/adapters/storage/postgres/mapper"
	"bank_service/pkg/adapters/storage/postgres/model"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"

	"bank_service/internal/services/wallet/domain"
)

type WalletRepository struct {
	BaseRepository
}

func (r *WalletRepository) Save(ctx context.Context, wallet *domain.Wallet) error {
	walletModel := mapper.WalletDomainToModel(wallet)
	return r.DB(ctx).Create(walletModel).Error
}

func (r *WalletRepository) FindByUserID(ctx context.Context, userID uint64) ([]*domain.Wallet, error) {
	var models []model.WalletModel
	err := r.DB(ctx).Where("user_id = ?", userID).Find(&models).Error
	if err != nil {
		return nil, err
	}

	wallets := make([]*domain.Wallet, len(models))
	for i, m := range models {
		wallets[i] = m.ToDomain()
	}
	return wallets, nil
}

func NewPostgresWalletRepository(db *gorm.DB) *WalletRepository {
	return &WalletRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

func (r *WalletRepository) Create(ctx context.Context, wallet *domain.Wallet) error {
	return r.DB(ctx).Create(mapper.WalletDomainToModel(wallet)).Error
}

func (r *WalletRepository) Update(ctx context.Context, wallet *domain.Wallet) error {
	walletModel := mapper.WalletDomainToModel(wallet)

	result := r.DB(ctx).
		Model(walletModel).
		Where("id = ? AND version = ?", wallet.ID, wallet.Version).
		Updates(map[string]interface{}{
			"user_id":    walletModel.UserID,
			"balance":    walletModel.Balance,
			"currency":   walletModel.Currency,
			"status":     walletModel.Status,
			"updated_at": walletModel.UpdatedAt,
			"version":    wallet.Version + 1,
		})

	if result.Error != nil {
		return fmt.Errorf("failed to update wallet: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("wallet not found or version mismatch")
	}

	wallet.Version++
	return nil
}

func (r *WalletRepository) FindByID(ctx context.Context, id domain.WalletID) (*domain.Wallet, error) {
	var walletModel model.WalletModel
	err := r.DB(ctx).
		Where("id = ?", id).
		First(&walletModel).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return walletModel.ToDomain(), nil
}

func (r *WalletRepository) GetVersion(ctx context.Context, id domain.WalletID) (int, error) {
	var walletModel model.WalletModel
	err := r.DB(ctx).
		Select("version").
		Where("id = ?", id).
		First(&walletModel).
		Error
	if err != nil {
		return 0, err
	}
	return walletModel.Version, nil
}
