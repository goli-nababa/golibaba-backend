package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type TransactionModel struct {
	ID            string         `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	FromWalletID  string         `gorm:"index;not null"`
	ToWalletID    string         `gorm:"index;not null"`
	Amount        int64          `gorm:"not null"`
	Currency      string         `gorm:"type:varchar(3);not null"`
	Type          string         `gorm:"type:transaction_type;not null"`
	Status        string         `gorm:"type:transaction_status;not null"`
	Description   string         `gorm:"type:text"`
	ReferenceID   string         `gorm:"index"`
	FailureReason string         `gorm:"type:text"`
	Metadata      JSONMap        `gorm:"type:jsonb"`
	StatusHistory []StatusChange `gorm:"type:jsonb"`
	BatchID       string         `gorm:"index"`
	CreatedAt     time.Time      `gorm:"not null"`
	UpdatedAt     time.Time      `gorm:"not null"`
	CompletedAt   *time.Time
	Version       int `gorm:"not null;default:1"`
}

type StatusChange struct {
	FromStatus string    `json:"from_status"`
	ToStatus   string    `json:"to_status"`
	Reason     string    `json:"reason"`
	ChangedAt  time.Time `json:"changed_at"`
}

type JSONMap map[string]interface{}

func (m JSONMap) Value() (driver.Value, error) {
	if m == nil {
		return nil, nil
	}
	return json.Marshal(m)
}

func (m *JSONMap) Scan(value interface{}) error {
	if value == nil {
		*m = nil
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal JSONB value: %v", value)
	}

	result := make(JSONMap)
	err := json.Unmarshal(bytes, &result)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSONB value: %v", err)
	}

	*m = result
	return nil
}

func (TransactionModel) TableName() string {
	return "transactions"
}
