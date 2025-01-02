package domain

import (
	"bank_service/internal/common/types"
	domain2 "bank_service/internal/services/transaction/domain"
	"time"
)

type CommissionEntry struct {
	TransactionID domain2.TransactionID
	Amount        *types.Money
	Rate          float64
	Status        string
	ProcessedAt   time.Time
}
