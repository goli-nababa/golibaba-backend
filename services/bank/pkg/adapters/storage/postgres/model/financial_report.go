package model

import (
	"encoding/json"
	"time"
)

type FinancialReportModel struct {
	ID          string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	ReportType  string    `gorm:"index"`
	BusinessID  uint64    `gorm:"index"`
	StartDate   time.Time `gorm:"index"`
	EndDate     time.Time `gorm:"index"`
	Granularity string
	Status      string          `gorm:"index"`
	Metrics     json.RawMessage `gorm:"type:jsonb"`
	GeneratedAt time.Time
	GeneratedBy uint64
	Format      string
	DataURL     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (m *FinancialReportModel) TableName() string {
	return "financial_reports"
}
