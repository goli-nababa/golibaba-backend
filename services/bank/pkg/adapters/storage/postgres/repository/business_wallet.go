package repository

import (
	moneyDomain "bank_service/internal/common/types"
	"bank_service/internal/services/business/domain"
	walletDomain "bank_service/internal/services/wallet/domain"
	"bank_service/pkg/adapters/storage/postgres/model"
	"context"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
)

type BusinessRepository struct {
	BaseRepository
}

func NewBusinessRepository(db *gorm.DB) *BusinessRepository {
	return &BusinessRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

func (r *BusinessRepository) Save(ctx context.Context, wallet *domain.BusinessWallet) error {
	if err := r.DB(ctx).Create(&model.WalletModel{
		ID:        string(wallet.ID),
		UserID:    wallet.UserID,
		Balance:   wallet.Balance.Amount,
		Currency:  wallet.Balance.Currency,
		Status:    string(wallet.Status),
		CreatedAt: wallet.CreatedAt,
		UpdatedAt: wallet.UpdatedAt,
		Version:   wallet.Version,
	}).Error; err != nil {
		return fmt.Errorf("failed to save base wallet: %w", err)
	}

	bankInfo, err := json.Marshal(wallet.BankInfo)
	if err != nil {
		return fmt.Errorf("failed to marshal bank info: %w", err)
	}

	model := &model.BusinessWalletModel{
		WalletModel: model.WalletModel{
			ID: string(wallet.ID),
		},
		BusinessID:     wallet.BusinessID,
		BusinessType:   string(wallet.BusinessType),
		CommissionRate: wallet.CommissionRate,
		PayoutSchedule: wallet.PayoutSchedule,
		LastPayoutDate: wallet.LastPayoutDate,
		BankInfo:       bankInfo,
	}

	if wallet.MinimumPayout != nil {
		model.MinPayoutAmount = wallet.MinimumPayout.Amount
		model.MinPayoutCurrency = wallet.MinimumPayout.Currency
	}

	if err := r.DB(ctx).Create(model).Error; err != nil {
		return fmt.Errorf("failed to save business wallet: %w", err)
	}

	return nil
}

func (r *BusinessRepository) Update(ctx context.Context, wallet *domain.BusinessWallet) error {
	if err := r.DB(ctx).Model(&model.WalletModel{}).
		Where("id = ? AND version = ?", wallet.ID, wallet.Version).
		Updates(map[string]interface{}{
			"balance":    wallet.Balance.Amount,
			"currency":   wallet.Balance.Currency,
			"status":     string(wallet.Status),
			"updated_at": wallet.UpdatedAt,
			"version":    wallet.Version + 1,
		}).Error; err != nil {
		return fmt.Errorf("failed to update base wallet: %w", err)
	}

	bankInfo, err := json.Marshal(wallet.BankInfo)
	if err != nil {
		return fmt.Errorf("failed to marshal bank info: %w", err)
	}

	updates := map[string]interface{}{
		"business_type":    string(wallet.BusinessType),
		"commission_rate":  wallet.CommissionRate,
		"payout_schedule":  wallet.PayoutSchedule,
		"last_payout_date": wallet.LastPayoutDate,
		"bank_info":        bankInfo,
	}

	if wallet.MinimumPayout != nil {
		updates["min_payout_amount"] = wallet.MinimumPayout.Amount
		updates["min_payout_currency"] = wallet.MinimumPayout.Currency
	}

	if err := r.DB(ctx).Model(&model.BusinessWalletModel{}).
		Where("id = ?", wallet.ID).
		Updates(updates).Error; err != nil {
		return fmt.Errorf("failed to update business wallet: %w", err)
	}

	wallet.Version++
	return nil
}

func (r *BusinessRepository) GetByWalletID(ctx context.Context, walletID walletDomain.WalletID) (*domain.BusinessWallet, error) {
	var baseModel model.WalletModel
	var businessModel model.BusinessWalletModel

	if err := r.DB(ctx).First(&baseModel, "id = ?", string(walletID)).Error; err != nil {
		return nil, fmt.Errorf("failed to get base wallet: %w", err)
	}

	if err := r.DB(ctx).First(&businessModel, "id = ?", string(walletID)).Error; err != nil {
		return nil, fmt.Errorf("failed to get business wallet: %w", err)
	}

	return r.modelsToDomain(&baseModel, &businessModel)
}

func (r *BusinessRepository) GetByBusinessID(ctx context.Context, businessID uint64) (*domain.BusinessWallet, error) {
	var baseModel model.WalletModel
	var businessModel model.BusinessWalletModel

	if err := r.DB(ctx).Joins("JOIN business_wallets ON wallets.id = business_wallets.id").
		Where("business_wallets.business_id = ?", businessID).
		First(&baseModel).Error; err != nil {
		return nil, fmt.Errorf("failed to get wallet: %w", err)
	}

	if err := r.DB(ctx).Where("business_id = ?", businessID).
		First(&businessModel).Error; err != nil {
		return nil, fmt.Errorf("failed to get business wallet: %w", err)
	}

	return r.modelsToDomain(&baseModel, &businessModel)
}

func (r *BusinessRepository) ListBusinessWallets(ctx context.Context, filter *domain.BusinessWalletFilter) ([]*domain.BusinessWallet, error) {
	query := r.DB(ctx).Model(&model.BusinessWalletModel{}).
		Joins("JOIN wallets ON business_wallets.id = wallets.id")

	if filter != nil {
		if filter.BusinessType != nil {
			query = query.Where("business_type = ?", string(*filter.BusinessType))
		}
		if filter.Status != nil {
			query = query.Where("status = ?", string(*filter.Status))
		}
		if filter.MinBalance != nil {
			query = query.Where("balance >= ?", filter.MinBalance.Amount)
		}
		if filter.MaxBalance != nil {
			query = query.Where("balance <= ?", filter.MaxBalance.Amount)
		}
		if filter.CreatedAfter != nil {
			query = query.Where("wallets.created_at >= ?", filter.CreatedAfter)
		}
		if filter.CreatedBefore != nil {
			query = query.Where("wallets.created_at <= ?", filter.CreatedBefore)
		}
		if filter.Limit > 0 {
			query = query.Limit(filter.Limit)
		}
		if filter.Offset > 0 {
			query = query.Offset(filter.Offset)
		}
	}

	var businessModels []model.BusinessWalletModel
	if err := query.Find(&businessModels).Error; err != nil {
		return nil, fmt.Errorf("failed to get business wallets: %w", err)
	}

	var walletIDs []string
	for _, bw := range businessModels {
		walletIDs = append(walletIDs, bw.ID)
	}

	var baseModels []model.WalletModel
	if err := r.DB(ctx).Where("id IN ?", walletIDs).Find(&baseModels).Error; err != nil {
		return nil, fmt.Errorf("failed to get base wallets: %w", err)
	}

	wallets := make([]*domain.BusinessWallet, len(businessModels))
	for i := range businessModels {
		wallet, err := r.modelsToDomain(&baseModels[i], &businessModels[i])
		if err != nil {
			return nil, err
		}
		wallets[i] = wallet
	}

	return wallets, nil
}

func (r *BusinessRepository) GetVersion(ctx context.Context, walletID walletDomain.WalletID) (int, error) {
	var version int
	err := r.DB(ctx).Model(&model.WalletModel{}).
		Select("version").
		Where("id = ?", string(walletID)).
		First(&version).Error
	if err != nil {
		return 0, fmt.Errorf("failed to get version: %w", err)
	}
	return version, nil
}

func (r *BusinessRepository) GetCentralWallet(ctx context.Context) (*walletDomain.Wallet, error) {
	var model model.WalletModel
	if err := r.DB(ctx).First(&model, "id = ?", domain.CENTRAL_WALLET_ID).Error; err != nil {
		return nil, fmt.Errorf("failed to get central wallet: %w", err)
	}

	balance, err := moneyDomain.NewMoney(float64(model.Balance), model.Currency)
	if err != nil {
		return nil, fmt.Errorf("invalid balance: %w", err)
	}

	return &walletDomain.Wallet{
		ID:        walletDomain.WalletID(model.ID),
		UserID:    model.UserID,
		Balance:   balance,
		Status:    walletDomain.WalletStatus(model.Status),
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		Version:   model.Version,
	}, nil
}

func (r *BusinessRepository) Delete(ctx context.Context, walletID walletDomain.WalletID) error {
	if err := r.DB(ctx).Delete(&model.BusinessWalletModel{}, "id = ?", string(walletID)).Error; err != nil {
		return fmt.Errorf("failed to delete business wallet: %w", err)
	}

	if err := r.DB(ctx).Delete(&model.WalletModel{}, "id = ?", string(walletID)).Error; err != nil {
		return fmt.Errorf("failed to delete base wallet: %w", err)
	}

	return nil
}

func (r *BusinessRepository) modelsToDomain(baseModel *model.WalletModel, businessModel *model.BusinessWalletModel) (*domain.BusinessWallet, error) {
	balance, err := moneyDomain.NewMoney(float64(baseModel.Balance), baseModel.Currency)
	if err != nil {
		return nil, fmt.Errorf("invalid balance: %w", err)
	}

	var bankInfo *domain.BankAccountInfo
	if businessModel.BankInfo != nil {
		if err := json.Unmarshal(businessModel.BankInfo, &bankInfo); err != nil {
			return nil, fmt.Errorf("failed to unmarshal bank info: %w", err)
		}
	}

	baseWallet := &walletDomain.Wallet{
		ID:        walletDomain.WalletID(baseModel.ID),
		UserID:    baseModel.UserID,
		Balance:   balance,
		Status:    walletDomain.WalletStatus(baseModel.Status),
		CreatedAt: baseModel.CreatedAt,
		UpdatedAt: baseModel.UpdatedAt,
		Version:   baseModel.Version,
	}

	var minPayout *moneyDomain.Money
	if businessModel.MinPayoutAmount > 0 {
		minPayout, err = moneyDomain.NewMoney(float64(businessModel.MinPayoutAmount), businessModel.MinPayoutCurrency)
		if err != nil {
			return nil, fmt.Errorf("invalid minimum payout: %w", err)
		}
	}

	return &domain.BusinessWallet{
		Wallet:         baseWallet,
		BusinessID:     businessModel.BusinessID,
		BusinessType:   domain.BusinessType(businessModel.BusinessType),
		CommissionRate: businessModel.CommissionRate,
		PayoutSchedule: businessModel.PayoutSchedule,
		LastPayoutDate: businessModel.LastPayoutDate,
		MinimumPayout:  minPayout,
		BankInfo:       bankInfo,
	}, nil
}
