package domain

import (
	"bank_service/internal/common/types"
	"time"
)

// PayoutEntry represents a single payout transaction
type PayoutEntry struct {
	ID            string
	Amount        *types.Money
	Status        string
	RequestedAt   time.Time
	ProcessedAt   *time.Time
	BankInfo      *BankAccountInfo
	ReferenceID   string
	FailureReason string
}
