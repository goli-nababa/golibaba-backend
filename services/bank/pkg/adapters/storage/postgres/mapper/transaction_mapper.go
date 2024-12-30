package mapper

import (
	"bank_service/internal/common/types"
	domainTx "bank_service/internal/services/transaction/domain"
	walletDomain "bank_service/internal/services/wallet/domain"
	"bank_service/pkg/adapters/storage/postgres/model"
)

func ToModel(tx *domainTx.Transaction) *model.TransactionModel {
	return &model.TransactionModel{
		ID:            string(tx.ID),
		FromWalletID:  string(tx.FromWalletID),
		ToWalletID:    string(tx.ToWalletID),
		Amount:        tx.Amount.Amount,
		Currency:      tx.Amount.Currency,
		Type:          string(tx.Type),
		Status:        string(tx.Status),
		Description:   tx.Description,
		ReferenceID:   tx.ReferenceID,
		FailureReason: tx.FailureReason,
		Metadata:      model.JSONMap(tx.Metadata),
		StatusHistory: ToModelStatusChanges(tx.StatusHistory),
		CreatedAt:     tx.CreatedAt,
		UpdatedAt:     tx.UpdatedAt,
		CompletedAt:   tx.CompletedAt,
		Version:       tx.Version,
	}
}

func ToDomain(m *model.TransactionModel) (*domainTx.Transaction, error) {
	amount, err := types.NewMoney(float64(m.Amount), m.Currency)
	if err != nil {
		return nil, err
	}

	return &domainTx.Transaction{
		ID:            domainTx.TransactionID(m.ID),
		FromWalletID:  walletDomain.WalletID(m.FromWalletID),
		ToWalletID:    walletDomain.WalletID(m.ToWalletID),
		Amount:        amount,
		Type:          domainTx.TransactionType(m.Type),
		Status:        domainTx.TransactionStatus(m.Status),
		Description:   m.Description,
		ReferenceID:   m.ReferenceID,
		FailureReason: m.FailureReason,
		Metadata:      map[string]interface{}(m.Metadata),
		StatusHistory: ToDomainStatusChanges(m.StatusHistory),
		CreatedAt:     m.CreatedAt,
		UpdatedAt:     m.UpdatedAt,
		CompletedAt:   m.CompletedAt,
		Version:       m.Version,
	}, nil
}

func ToModelStatusChanges(changes []domainTx.StatusChange) []model.StatusChange {
	result := make([]model.StatusChange, len(changes))
	for i, change := range changes {
		result[i] = model.StatusChange{
			FromStatus: string(change.FromStatus),
			ToStatus:   string(change.ToStatus),
			Reason:     change.Reason,
			ChangedAt:  change.ChangedAt,
		}
	}
	return result
}

func ToDomainStatusChanges(changes []model.StatusChange) []domainTx.StatusChange {
	result := make([]domainTx.StatusChange, len(changes))
	for i, change := range changes {
		result[i] = domainTx.StatusChange{
			FromStatus: domainTx.TransactionStatus(change.FromStatus),
			ToStatus:   domainTx.TransactionStatus(change.ToStatus),
			Reason:     change.Reason,
			ChangedAt:  change.ChangedAt,
		}
	}
	return result
}
