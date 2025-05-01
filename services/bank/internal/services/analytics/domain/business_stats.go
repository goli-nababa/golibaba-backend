package domain

import (
	"bank_service/internal/common/types"
	"time"
)

type BusinessStats struct {
	TotalRevenue      *types.Money
	TotalCommission   *types.Money
	TotalPayouts      *types.Money
	TransactionCount  int64
	SuccessfulTxCount int64
	FailedTxCount     int64
	AverageOrderValue *types.Money
	CommissionRate    float64
	Period            struct {
		Start time.Time
		End   time.Time
	}
}
