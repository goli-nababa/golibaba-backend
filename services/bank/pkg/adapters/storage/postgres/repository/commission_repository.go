package repository

import (
	moneyDomain "bank_service/internal/common/types"
	"context"

	"fmt"
	"time"

	businessDomain "bank_service/internal/services/business/domain"
	"bank_service/internal/services/commission/domain"
	txDomain "bank_service/internal/services/transaction/domain"
	walletDomain "bank_service/internal/services/wallet/domain"
	model2 "bank_service/pkg/adapters/storage/postgres/model"
	"gorm.io/gorm"
)

type CommissionRepository struct {
	BaseRepository
}

func NewCommissionRepository(db *gorm.DB) *CommissionRepository {
	return &CommissionRepository{

		BaseRepository: NewBaseRepository(db),
	}
}

func (r *CommissionRepository) Create(ctx context.Context, commission *domain.Commission) error {
	model := &model2.CommissionModel{
		ID:            commission.ID,
		TransactionID: string(commission.TransactionID),
		Amount:        commission.Amount.Amount,
		Currency:      commission.Amount.Currency,
		Rate:          commission.Rate,
		RecipientID:   string(commission.RecipientID),
		BusinessType:  string(commission.BusinessType),
		Status:        string(commission.Status),
		CreatedAt:     commission.CreatedAt,
		PaidAt:        commission.PaidAt,
		Description:   commission.Description,
	}

	return r.DB(ctx).Create(model).Error
}

func (r *CommissionRepository) Update(ctx context.Context, commission *domain.Commission) error {
	model := &model2.CommissionModel{
		ID:            commission.ID,
		TransactionID: string(commission.TransactionID),
		Amount:        commission.Amount.Amount,
		Currency:      commission.Amount.Currency,
		Rate:          commission.Rate,
		RecipientID:   string(commission.RecipientID),
		BusinessType:  string(commission.BusinessType),
		Status:        string(commission.Status),
		PaidAt:        commission.PaidAt,
		Description:   commission.Description,
	}

	return r.DB(ctx).Model(&model2.CommissionModel{}).Where("id = ?", commission.ID).Updates(model).Error
}

func (r *CommissionRepository) GetByID(ctx context.Context, id string) (*domain.Commission, error) {
	var model model2.CommissionModel
	if err := r.DB(ctx).First(&model, "id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("failed to get commission: %w", err)
	}

	return modelToDomain(&model)
}

func (r *CommissionRepository) GetByTransactionID(ctx context.Context, txID txDomain.TransactionID) (*domain.Commission, error) {
	var model model2.CommissionModel
	if err := r.DB(ctx).First(&model, "transaction_id = ?", string(txID)).Error; err != nil {
		return nil, fmt.Errorf("failed to get commission by transaction: %w", err)
	}

	return modelToDomain(&model)
}

func (r *CommissionRepository) FindByBusinessIDAndDateRange(ctx context.Context, businessID uint64, startDate, endDate time.Time) ([]*domain.Commission, error) {
	var models []model2.CommissionModel
	if err := r.DB(ctx).
		Where("recipient_id = ? AND created_at BETWEEN ? AND ?",
			fmt.Sprintf("business_%d", businessID), startDate, endDate).
		Find(&models).Error; err != nil {
		return nil, fmt.Errorf("failed to get business commissions: %w", err)
	}

	commissions := make([]*domain.Commission, len(models))
	for i, model := range models {
		commission, err := modelToDomain(&model)
		if err != nil {
			return nil, err
		}
		commissions[i] = commission
	}

	return commissions, nil
}

func (r *CommissionRepository) FindByStatus(ctx context.Context, status domain.PaymentStatus) ([]*domain.Commission, error) {
	var models []model2.CommissionModel
	if err := r.DB(ctx).Where("status = ?", string(status)).Find(&models).Error; err != nil {
		return nil, fmt.Errorf("failed to get commissions by status: %w", err)
	}

	commissions := make([]*domain.Commission, len(models))
	for i, model := range models {
		commission, err := modelToDomain(&model)
		if err != nil {
			return nil, err
		}
		commissions[i] = commission
	}

	return commissions, nil
}

func modelToDomain(model *model2.CommissionModel) (*domain.Commission, error) {
	amount, err := moneyDomain.NewMoney(float64(model.Amount), model.Currency)
	if err != nil {
		return nil, fmt.Errorf("invalid amount in commission model: %w", err)
	}

	return &domain.Commission{
		ID:            model.ID,
		TransactionID: txDomain.TransactionID(model.TransactionID),
		Amount:        amount,
		Rate:          model.Rate,
		RecipientID:   walletDomain.WalletID(model.RecipientID),
		BusinessType:  businessDomain.BusinessType(model.BusinessType),
		Status:        domain.PaymentStatus(model.Status),
		CreatedAt:     model.CreatedAt,
		PaidAt:        model.PaidAt,
		Description:   model.Description,
	}, nil
}
