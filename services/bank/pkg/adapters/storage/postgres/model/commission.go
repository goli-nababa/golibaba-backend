package model

import "time"

type CommissionModel struct {
	ID            string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	TransactionID string    `gorm:"index;not null"`
	Amount        int64     `gorm:"not null"`
	Currency      string    `gorm:"type:varchar(3);not null"`
	Rate          float64   `gorm:"type:decimal(5,4);not null"`
	RecipientID   string    `gorm:"index;not null"`
	BusinessType  string    `gorm:"type:business_type;not null"`
	Status        string    `gorm:"type:varchar(20);not null"`
	CreatedAt     time.Time `gorm:"not null"`
	PaidAt        *time.Time
	Description   string  `gorm:"type:text"`
	Metadata      JSONMap `gorm:"type:jsonb"`
}

func (m *CommissionModel) TableName() string {
	return "commission_models"
}
