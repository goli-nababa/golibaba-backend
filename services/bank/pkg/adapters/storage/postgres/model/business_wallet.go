package model

import (
	"encoding/json"
	"time"
)

type BusinessWalletModel struct {
	WalletModel
	BusinessID        uint64  `gorm:"index;not null"`
	BusinessType      string  `gorm:"type:business_type;not null"`
	CommissionRate    float64 `gorm:"type:decimal(5,4);not null"`
	PayoutSchedule    string  `gorm:"type:varchar(20)"`
	LastPayoutDate    *time.Time
	MinPayoutAmount   int64
	MinPayoutCurrency string          `gorm:"type:varchar(3)"`
	BankInfo          json.RawMessage `gorm:"type:jsonb"`

	tableName struct{} `gorm:"table:business_wallets"`
}
