package domain

import (
	"bank_service/internal/common/types"
	businessWalletDomain "bank_service/internal/services/business/domain"
	transactionDomain "bank_service/internal/services/transaction/domain"
	walletDomain "bank_service/internal/services/wallet/domain"
	"errors"
	"github.com/google/uuid"
	"time"
)

type PaymentStatus string

const (
	PaymentStatusPending    PaymentStatus = "pending"
	PaymentStatusProcessing PaymentStatus = "processing"
	PaymentStatusProcessed  PaymentStatus = "processed"
	PaymentStatusFailed     PaymentStatus = "failed"
)

type Commission struct {
	ID            string
	TransactionID transactionDomain.TransactionID
	Amount        *types.Money
	Rate          float64
	RecipientID   walletDomain.WalletID
	BusinessType  businessWalletDomain.BusinessType
	Status        PaymentStatus
	CreatedAt     time.Time
	PaidAt        *time.Time
	Description   string
}

func NewCommission(tx *transactionDomain.Transaction, rate float64, businessType businessWalletDomain.BusinessType) (*Commission, error) {
	if rate < 0 || rate > 1 {
		return nil, errors.New("invalid commission rate")
	}

	commissionAmount, err := calculateCommissionAmount(tx.Amount, rate)
	if err != nil {
		return nil, err
	}

	return &Commission{
		ID:            uuid.New().String(),
		TransactionID: tx.ID,
		Amount:        commissionAmount,
		Rate:          rate,
		RecipientID:   tx.ToWalletID,
		BusinessType:  businessType,
		Status:        PaymentStatusPending,
		CreatedAt:     time.Now(),
	}, nil
}

func calculateCommissionAmount(baseAmount *types.Money, rate float64) (*types.Money, error) {
	commissionFloat := float64(baseAmount.Amount) * rate
	return types.NewMoney(commissionFloat, baseAmount.Currency)
}
