package domain

import (
	moneyDomain "bank_service/internal/common/types"
	txDomain "bank_service/internal/services/transaction/domain"
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	ErrInvalidPaymentStatus   = errors.New("invalid payment status")
	ErrInvalidGatewayResponse = errors.New("invalid gateway response")
	ErrPaymentNotFound        = errors.New("payment not found")
)

type PaymentStatus string

const (
	PaymentStatusPending    PaymentStatus = "pending"
	PaymentStatusProcessing PaymentStatus = "processing"
	PaymentStatusSuccessful PaymentStatus = "successful"
	PaymentStatusFailed     PaymentStatus = "failed"
	PaymentStatusCancelled  PaymentStatus = "cancelled"
	PaymentStatusRefunded   PaymentStatus = "refunded"
)

type PaymentType string

const (
	PaymentTypeGateway  PaymentType = "gateway"
	PaymentTypeWallet   PaymentType = "wallet"
	PaymentTypeTransfer PaymentType = "transfer"
)

type Payment struct {
	ID            string
	TransactionID txDomain.TransactionID
	Amount        *moneyDomain.Money
	Type          PaymentType
	Status        PaymentStatus
	ReferenceID   string // Gateway reference ID
	GatewayName   string // Name of the payment gateway used
	Description   string
	Metadata      map[string]interface{}
	CreatedAt     time.Time
	UpdatedAt     time.Time
	CompletedAt   *time.Time
	FailureReason string
}

func NewPayment(txID txDomain.TransactionID, amount *moneyDomain.Money, paymentType PaymentType) (*Payment, error) {
	if amount.Amount <= 0 {
		return nil, moneyDomain.ErrInvalidAmount
	}

	now := time.Now()
	return &Payment{
		ID:            uuid.New().String(),
		TransactionID: txID,
		Amount:        amount,
		Type:          paymentType,
		Status:        PaymentStatusPending,
		CreatedAt:     now,
		UpdatedAt:     now,
		Metadata:      make(map[string]interface{}),
	}, nil
}

func (p *Payment) Complete(referenceID string) error {
	if p.Status != PaymentStatusPending && p.Status != PaymentStatusProcessing {
		return ErrInvalidPaymentStatus
	}

	now := time.Now()
	p.Status = PaymentStatusSuccessful
	p.ReferenceID = referenceID
	p.CompletedAt = &now
	p.UpdatedAt = now
	return nil
}

func (p *Payment) Fail(reason string) error {
	if p.Status != PaymentStatusPending && p.Status != PaymentStatusProcessing {
		return ErrInvalidPaymentStatus
	}

	p.Status = PaymentStatusFailed
	p.FailureReason = reason
	p.UpdatedAt = time.Now()
	return nil
}

func (p *Payment) Cancel(reason string) error {
	if p.Status != PaymentStatusPending {
		return ErrInvalidPaymentStatus
	}

	p.Status = PaymentStatusCancelled
	p.FailureReason = reason
	p.UpdatedAt = time.Now()
	return nil
}

func (p *Payment) Refund() error {
	if p.Status != PaymentStatusSuccessful {
		return ErrInvalidPaymentStatus
	}

	p.Status = PaymentStatusRefunded
	p.UpdatedAt = time.Now()
	return nil
}
