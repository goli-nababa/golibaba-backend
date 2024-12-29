package repository

import (
	"bank_service/internal/services/transaction/domain"
	"bank_service/internal/services/transaction/port"
	walletDomain "bank_service/internal/services/wallet/domain"
	"bank_service/pkg/adapters/storage/postgres/mapper"

	"bank_service/pkg/adapters/storage/postgres/model"

	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type transactionRepository struct {
	BaseRepository
}

func NewTransactionRepository(db *gorm.DB) port.TransactionRepository {
	return &transactionRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

func (r *transactionRepository) FindByReferenceID(ctx context.Context, referenceID string) (*domain.Transaction, error) {
	var transactionModel model.TransactionModel

	result := r.DB(ctx).
		Where("reference_id = ?", "https://sandbox.zarinpal.com/pg/StartPay/"+referenceID).
		Order("created_at DESC").
		First(&transactionModel)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("transaction not found for reference ID: %s", referenceID)
		}
		return nil, fmt.Errorf("failed to get transaction by reference ID: %w", result.Error)
	}

	return mapper.ToDomain(&transactionModel)
}
func (r *transactionRepository) FindByWallet(ctx context.Context, walletID walletDomain.WalletID, filter *port.TransactionFilter) ([]*domain.Transaction, error) {
	query := r.DB(ctx).Model(&model.TransactionModel{})

	// Add base wallet filter
	query = query.Where("from_wallet_id = ? OR to_wallet_id = ?", string(walletID), string(walletID))

	// Apply status filter
	if filter != nil && len(filter.Status) > 0 {
		statuses := make([]string, len(filter.Status))
		for i, status := range filter.Status {
			statuses[i] = string(status)
		}
		query = query.Where("status IN ?", statuses)
	}

	// Apply type filter
	if filter != nil && len(filter.Types) > 0 {
		types := make([]string, len(filter.Types))
		for i, txType := range filter.Types {
			types[i] = string(txType)
		}
		query = query.Where("type IN ?", types)
	}

	// Apply date range filter
	if filter != nil && !filter.From.IsZero() {
		query = query.Where("created_at >= ?", filter.From)
	}
	if filter != nil && !filter.To.IsZero() {
		query = query.Where("created_at <= ?", filter.To)
	}

	var models []model.TransactionModel
	if err := query.Find(&models).Error; err != nil {
		return nil, fmt.Errorf("failed to find transactions: %w", err)
	}

	transactions := make([]*domain.Transaction, len(models))
	for i, m := range models {
		tx, err := mapper.ToDomain(&m)
		if err != nil {
			return nil, fmt.Errorf("failed to convert model to domain: %w", err)
		}
		transactions[i] = tx
	}

	return transactions, nil
}

func (r *transactionRepository) FindByStatus(ctx context.Context, status domain.TransactionStatus) ([]*domain.Transaction, error) {
	var models []model.TransactionModel
	if err := r.DB(ctx).Where("status = ?", string(status)).Find(&models).Error; err != nil {
		return nil, fmt.Errorf("failed to find transactions by status: %w", err)
	}

	transactions := make([]*domain.Transaction, len(models))
	for i, m := range models {
		tx, err := mapper.ToDomain(&m)
		if err != nil {
			return nil, fmt.Errorf("failed to convert model to domain: %w", err)
		}
		transactions[i] = tx
	}

	return transactions, nil
}

func (r *transactionRepository) FindByBatchID(ctx context.Context, batchID string) ([]*domain.Transaction, error) {
	var models []model.TransactionModel
	if err := r.DB(ctx).Where("batch_id = ?", batchID).Find(&models).Error; err != nil {
		return nil, fmt.Errorf("failed to find transactions by batch ID: %w", err)
	}

	transactions := make([]*domain.Transaction, len(models))
	for i, m := range models {
		tx, err := mapper.ToDomain(&m)
		if err != nil {
			return nil, fmt.Errorf("failed to convert model to domain: %w", err)
		}
		transactions[i] = tx
	}

	return transactions, nil
}

func (r *transactionRepository) FindByDateRange(ctx context.Context, start, end time.Time) ([]*domain.Transaction, error) {
	var models []model.TransactionModel
	if err := r.DB(ctx).
		Where("created_at BETWEEN ? AND ?", start, end).
		Find(&models).Error; err != nil {
		return nil, fmt.Errorf("failed to find transactions by date range: %w", err)
	}

	transactions := make([]*domain.Transaction, len(models))
	for i, m := range models {
		tx, err := mapper.ToDomain(&m)
		if err != nil {
			return nil, fmt.Errorf("failed to convert model to domain: %w", err)
		}
		transactions[i] = tx
	}

	return transactions, nil
}

func (r *transactionRepository) SaveBatch(ctx context.Context, txs []*domain.Transaction) error {

	tx := r.DB(ctx)

	for _, transaction := range txs {
		model := mapper.ToModel(transaction)
		if err := tx.Create(model).Error; err != nil {

			return fmt.Errorf("failed to save transaction %s: %w", transaction.ID, err)
		}
	}

	return nil
}

func (r *transactionRepository) FindByBusinessIDAndDateRange(ctx context.Context, businessID uint64, startDate, endDate time.Time) ([]*domain.Transaction, error) {
	var models []model.TransactionModel
	if err := r.DB(ctx).
		Where("(from_wallet_id = ? OR to_wallet_id = ?) AND created_at BETWEEN ? AND ?",
			fmt.Sprintf("business_%d", businessID),
			fmt.Sprintf("business_%d", businessID),
			startDate,
			endDate).
		Find(&models).Error; err != nil {
		return nil, fmt.Errorf("failed to find business transactions: %w", err)
	}

	transactions := make([]*domain.Transaction, len(models))
	for i, m := range models {
		tx, err := mapper.ToDomain(&m)
		if err != nil {
			return nil, fmt.Errorf("failed to convert model to domain: %w", err)
		}
		transactions[i] = tx
	}

	return transactions, nil
}

func (r *transactionRepository) Create(ctx context.Context, tx *domain.Transaction) error {
	transactionModel := mapper.ToModel(tx)

	result := r.DB(ctx).Create(transactionModel)
	if result.Error != nil {
		return fmt.Errorf("failed to create transaction: %w", result.Error)
	}

	return nil
}

func (r *transactionRepository) Update(ctx context.Context, tx *domain.Transaction) error {
	transactionModel := mapper.ToModel(tx)

	result := r.DB(ctx).Model(&transactionModel).
		Where("id = ? AND version = ?", transactionModel.ID, transactionModel.Version).
		Updates(map[string]interface{}{
			"from_wallet_id": transactionModel.FromWalletID,
			"to_wallet_id":   transactionModel.ToWalletID,
			"amount":         transactionModel.Amount,
			"currency":       transactionModel.Currency,
			"type":           transactionModel.Type,
			"status":         transactionModel.Status,
			"description":    transactionModel.Description,
			"reference_id":   transactionModel.ReferenceID,
			"failure_reason": transactionModel.FailureReason,
			"metadata":       transactionModel.Metadata,
			"status_history": transactionModel.StatusHistory,
			"completed_at":   transactionModel.CompletedAt,
			"updated_at":     transactionModel.UpdatedAt,
			"version":        transactionModel.Version + 1,
		})

	if result.Error != nil {
		return fmt.Errorf("failed to update transaction: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("transaction not found or version mismatch")
	}

	// If update was successful, increment version in the domain model
	tx.Version++

	return nil
}

func (r *transactionRepository) GetByID(ctx context.Context, id domain.TransactionID) (*domain.Transaction, error) {
	var transactionModel model.TransactionModel

	result := r.DB(ctx).Where("id = ?", string(id)).First(&transactionModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("transaction not found: %s", id)
		}
		return nil, fmt.Errorf("failed to get transaction: %w", result.Error)
	}

	return mapper.ToDomain(&transactionModel)
}

func (r *transactionRepository) Delete(ctx context.Context, id domain.TransactionID) error {
	result := r.DB(ctx).Delete(&model.TransactionModel{}, "id = ?", string(id))
	if result.Error != nil {
		return fmt.Errorf("failed to delete transaction: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("transaction not found: %s", id)
	}

	return nil
}

func (r *transactionRepository) GetByStatus(ctx context.Context, status domain.TransactionStatus) ([]*domain.Transaction, error) {
	var models []model.TransactionModel

	result := r.DB(ctx).Where("status = ?", string(status)).Find(&models)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to get transactions by status: %w", result.Error)
	}

	transactions := make([]*domain.Transaction, len(models))
	for i, m := range models {
		tx, err := mapper.ToDomain(&m)
		if err != nil {
			return nil, fmt.Errorf("failed to convert model to domain: %w", err)
		}
		transactions[i] = tx
	}

	return transactions, nil
}

func (r *transactionRepository) GetByWalletID(ctx context.Context, walletID walletDomain.WalletID) ([]*domain.Transaction, error) {
	var models []model.TransactionModel

	result := r.DB(ctx).
		Where("from_wallet_id = ? OR to_wallet_id = ?", string(walletID), string(walletID)).
		Order("created_at DESC").
		Find(&models)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to get wallet transactions: %w", result.Error)
	}

	transactions := make([]*domain.Transaction, len(models))
	for i, m := range models {
		tx, err := mapper.ToDomain(&m)
		if err != nil {
			return nil, fmt.Errorf("failed to convert model to domain: %w", err)
		}
		transactions[i] = tx
	}

	return transactions, nil
}

func (r *transactionRepository) GetByDateRange(ctx context.Context, start, end time.Time) ([]*domain.Transaction, error) {
	var models []model.TransactionModel

	result := r.DB(ctx).
		Where("created_at BETWEEN ? AND ?", start, end).
		Order("created_at DESC").
		Find(&models)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to get transactions by date range: %w", result.Error)
	}

	transactions := make([]*domain.Transaction, len(models))
	for i, m := range models {
		tx, err := mapper.ToDomain(&m)
		if err != nil {
			return nil, fmt.Errorf("failed to convert model to domain: %w", err)
		}
		transactions[i] = tx
	}

	return transactions, nil
}

func (r *transactionRepository) GetByType(ctx context.Context, txType domain.TransactionType) ([]*domain.Transaction, error) {
	var models []model.TransactionModel

	result := r.DB(ctx).
		Where("type = ?", string(txType)).
		Order("created_at DESC").
		Find(&models)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to get transactions by type: %w", result.Error)
	}

	transactions := make([]*domain.Transaction, len(models))
	for i, m := range models {
		tx, err := mapper.ToDomain(&m)
		if err != nil {
			return nil, fmt.Errorf("failed to convert model to domain: %w", err)
		}
		transactions[i] = tx
	}

	return transactions, nil
}

func (r *transactionRepository) GetPendingTransactions(ctx context.Context) ([]*domain.Transaction, error) {
	return r.GetByStatus(ctx, domain.TransactionStatusPending)
}

func (r *transactionRepository) GetFailedTransactions(ctx context.Context) ([]*domain.Transaction, error) {
	return r.GetByStatus(ctx, domain.TransactionStatusFailed)
}

func (r *transactionRepository) GetWalletTransactions(ctx context.Context, walletID walletDomain.WalletID, limit, offset int) ([]*domain.Transaction, error) {
	var models []model.TransactionModel

	result := r.DB(ctx).
		Where("from_wallet_id = ? OR to_wallet_id = ?", string(walletID), string(walletID)).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&models)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to get wallet transactions: %w", result.Error)
	}

	transactions := make([]*domain.Transaction, len(models))
	for i, m := range models {
		tx, err := mapper.ToDomain(&m)
		if err != nil {
			return nil, fmt.Errorf("failed to convert model to domain: %w", err)
		}
		transactions[i] = tx
	}

	return transactions, nil
}

func (r *transactionRepository) GetTransactionsByReference(ctx context.Context, referenceID string) ([]*domain.Transaction, error) {
	var models []model.TransactionModel

	result := r.DB(ctx).
		Where("reference_id = ?", referenceID).
		Find(&models)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to get transactions by reference: %w", result.Error)
	}

	transactions := make([]*domain.Transaction, len(models))
	for i, m := range models {
		tx, err := mapper.ToDomain(&m)
		if err != nil {
			return nil, fmt.Errorf("failed to convert model to domain: %w", err)
		}
		transactions[i] = tx
	}

	return transactions, nil
}

func (r *transactionRepository) CountByStatus(ctx context.Context, status domain.TransactionStatus) (int64, error) {
	var count int64

	result := r.DB(ctx).
		Model(&model.TransactionModel{}).
		Where("status = ?", string(status)).
		Count(&count)

	if result.Error != nil {
		return 0, fmt.Errorf("failed to count transactions by status: %w", result.Error)
	}

	return count, nil
}

func (r *transactionRepository) GetTransactionHistory(ctx context.Context, id domain.TransactionID) ([]domain.StatusChange, error) {
	var transactionModel model.TransactionModel

	result := r.DB(ctx).
		Select("status_history").
		Where("id = ?", string(id)).
		First(&transactionModel)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("transaction not found: %s", id)
		}
		return nil, fmt.Errorf("failed to get transaction history: %w", result.Error)
	}

	return mapper.ToDomainStatusChanges(transactionModel.StatusHistory), nil
}

func (r *transactionRepository) CreateBatch(ctx context.Context, txs []*domain.Transaction) error {
	models := make([]model.TransactionModel, len(txs))
	for i, tx := range txs {
		models[i] = *mapper.ToModel(tx)
	}

	result := r.DB(ctx).Create(&models)
	if result.Error != nil {
		return fmt.Errorf("failed to create transactions in batch: %w", result.Error)
	}

	return nil
}

func (r *transactionRepository) UpdateBatch(ctx context.Context, txs []*domain.Transaction) error {
	tx := r.DB(ctx)

	for _, transaction := range txs {
		data := transaction
		tx.Model(&transaction).
			Where("id = ? AND version = ?", data.ID, data.Version-1).
			Updates(data)
	}

	return nil
}

func (r *transactionRepository) GetVersion(ctx context.Context, id domain.TransactionID) (int, error) {
	var transactionModel model.TransactionModel

	result := r.DB(ctx).
		Select("version").
		Where("id = ?", string(id)).
		First(&transactionModel)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return 0, fmt.Errorf("transaction not found: %s", id)
		}
		return 0, fmt.Errorf("failed to get transaction version: %w", result.Error)
	}

	return transactionModel.Version, nil
}

func (r *transactionRepository) UpdateWithVersion(ctx context.Context, tx *domain.Transaction, version int) error {
	data := mapper.ToModel(tx)

	result := r.DB(ctx).Model(&data).
		Where("id = ? AND version = ?", data.ID, version).
		Updates(data)

	if result.Error != nil {
		return fmt.Errorf("failed to update transaction: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("transaction not found or version mismatch")
	}

	return nil
}

func (r *transactionRepository) GetByWalletAndStatus(ctx context.Context, walletID walletDomain.WalletID, status domain.TransactionStatus) ([]*domain.Transaction, error) {
	var models []model.TransactionModel

	result := r.DB(ctx).
		Where("(from_wallet_id = ? OR to_wallet_id = ?) AND status = ?",
			string(walletID),
			string(walletID),
			string(status)).
		Order("created_at DESC").
		Find(&models)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to get wallet transactions by status: %w", result.Error)
	}

	transactions := make([]*domain.Transaction, len(models))
	for i, m := range models {
		tx, err := mapper.ToDomain(&m)
		if err != nil {
			return nil, fmt.Errorf("failed to convert model to domain: %w", err)
		}
		transactions[i] = tx
	}

	return transactions, nil
}

func (r *transactionRepository) Save(ctx context.Context, tx *domain.Transaction) error {
	transactionModel := mapper.ToModel(tx)
	result := r.DB(ctx).Create(transactionModel)
	if result.Error != nil {
		return fmt.Errorf("failed to save transaction: %w", result.Error)
	}
	return nil
}

func (r *transactionRepository) FindByID(ctx context.Context, id domain.TransactionID) (*domain.Transaction, error) {
	var model model.TransactionModel
	result := r.DB(ctx).Where("id = ?", id).First(&model)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("transaction not found: %s", id)
		}
		return nil, fmt.Errorf("failed to find transaction: %w", result.Error)
	}
	return mapper.ToDomain(&model)
}

func (r *transactionRepository) FindByWalletID(ctx context.Context, walletID string) ([]*domain.Transaction, error) {
	var models []model.TransactionModel
	result := r.DB(ctx).
		Where("from_wallet_id = ? OR to_wallet_id = ?", walletID, walletID).
		Find(&models)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to find transactions: %w", result.Error)
	}

	transactions := make([]*domain.Transaction, len(models))
	for i, model := range models {
		tx, err := mapper.ToDomain(&model)
		if err != nil {
			return nil, err
		}
		transactions[i] = tx
	}
	return transactions, nil
}
