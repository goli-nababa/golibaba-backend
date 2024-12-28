package model

import (
	"encoding/json"
	"time"
)

type AnalyticsModel struct {
	ID          string          `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	BusinessID  uint64          `gorm:"index"`
	StartDate   time.Time       `gorm:"index"`
	EndDate     time.Time       `gorm:"index"`
	Metrics     json.RawMessage `gorm:"type:jsonb"`
	Trends      json.RawMessage `gorm:"type:jsonb"`
	Comparisons json.RawMessage `gorm:"type:jsonb"`
	GeneratedAt time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (m *AnalyticsModel) TableName() string {
	return "analytics_reports"
}
