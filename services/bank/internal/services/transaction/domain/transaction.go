package domain

import (
	"bank_service/internal/common/types"
	"errors"
	"time"

	walletDomain "bank_service/internal/services/wallet/domain"
	"github.com/google/uuid"
)

type TransactionID string

type TransactionType string

const (
	TransactionTypeDeposit    TransactionType = "deposit"
	TransactionTypeWithdrawal TransactionType = "withdrawal"
	TransactionTypeTransfer   TransactionType = "transfer"
	TransactionTypePayment    TransactionType = "payment"
	TransactionTypeRefund     TransactionType = "refund"
	TransactionTypeCommission TransactionType = "commission"
)

type TransactionStatus string

const (
	TransactionStatusPending    TransactionStatus = "pending"
	TransactionStatusProcessing TransactionStatus = "processing"
	TransactionStatusSuccess    TransactionStatus = "success"
	TransactionStatusFailed     TransactionStatus = "failed"
	TransactionStatusCancelled  TransactionStatus = "cancelled"
	TransactionStatusRefunded   TransactionStatus = "refunded"
)

var (
	ErrInvalidAmount    = errors.New("invalid amount")
	ErrInvalidStatus    = errors.New("invalid transaction status")
	ErrInvalidWallet    = errors.New("invalid wallet")
	ErrTransactionFinal = errors.New("transaction is already in a final state")
)

type Transaction struct {
	ID            TransactionID
	FromWalletID  walletDomain.WalletID
	ToWalletID    walletDomain.WalletID
	Amount        *types.Money
	Type          TransactionType
	Status        TransactionStatus
	Description   string
	ReferenceID   string
	FailureReason string
	Metadata      map[string]interface{}
	StatusHistory []StatusChange
	CreatedAt     time.Time
	UpdatedAt     time.Time
	CompletedAt   *time.Time
	Version       int
}

type StatusChange struct {
	FromStatus TransactionStatus
	ToStatus   TransactionStatus
	Reason     string
	ChangedAt  time.Time
}

func NewTransaction(
	from walletDomain.WalletID,
	to walletDomain.WalletID,
	amount *types.Money,
	txType TransactionType,
	desc string,
) (*Transaction, error) {
	if amount.Amount <= 0 {
		return nil, ErrInvalidAmount
	}
	now := time.Now()
	return &Transaction{
		ID:           TransactionID(uuid.New().String()),
		FromWalletID: from,
		ToWalletID:   to,
		Amount:       amount,
		Type:         txType,
		Status:       TransactionStatusPending,
		Description:  desc,
		CreatedAt:    now,
		UpdatedAt:    now,
		Version:      1,
	}, nil
}

func (t *Transaction) Complete() error {
	if t.Status != TransactionStatusProcessing {
		return ErrInvalidStatus
	}
	now := time.Now()
	t.Status = TransactionStatusSuccess
	t.CompletedAt = &now
	t.UpdatedAt = now
	return nil
}

func (t *Transaction) Process() error {
	if t.Status != TransactionStatusPending {
		return ErrInvalidStatus
	}
	t.Status = TransactionStatusProcessing
	t.UpdatedAt = time.Now()
	return nil
}

func (t *Transaction) Fail(reason string) error {
	if t.isFinalStatus(t.Status) {
		return ErrTransactionFinal
	}
	t.Status = TransactionStatusFailed
	t.FailureReason = reason
	t.UpdatedAt = time.Now()
	return nil
}

func (t *Transaction) Cancel(reason string) error {
	if t.Status != TransactionStatusPending {
		return ErrInvalidStatus
	}
	t.Status = TransactionStatusCancelled
	t.FailureReason = reason
	t.UpdatedAt = time.Now()
	return nil
}

func (t *Transaction) isFinalStatus(status TransactionStatus) bool {
	return status == TransactionStatusSuccess ||
		status == TransactionStatusFailed ||
		status == TransactionStatusCancelled ||
		status == TransactionStatusRefunded
}
