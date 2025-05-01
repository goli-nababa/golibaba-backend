package domain

import (
	moneyDomain "bank_service/internal/common/types"
	"errors"
	"github.com/google/uuid"
	"time"
)

type ReportType string
type ReportStatus string
type TimeGranularity string

const (
	ReportTypeRevenue     ReportType = "revenue"
	ReportTypeCommission  ReportType = "commission"
	ReportTypeRefund      ReportType = "refund"
	ReportTypeTransaction ReportType = "transaction"

	ReportStatusPending   ReportStatus = "pending"
	ReportStatusGenerated ReportStatus = "generated"
	ReportStatusFailed    ReportStatus = "failed"

	GranularityHourly  TimeGranularity = "hourly"
	GranularityDaily   TimeGranularity = "daily"
	GranularityWeekly  TimeGranularity = "weekly"
	GranularityMonthly TimeGranularity = "monthly"
)

type FinancialReport struct {
	ID          string
	ReportType  ReportType
	BusinessID  uint64
	StartDate   time.Time
	EndDate     time.Time
	Granularity TimeGranularity
	Status      ReportStatus
	Metrics     map[string]*moneyDomain.Money
	GeneratedAt time.Time
	GeneratedBy uint64
	Format      ReportFormat
	DataURL     string
}

func NewFinancialReport(
	reportType ReportType,
	businessID uint64,
	startDate, endDate time.Time,
	granularity TimeGranularity,
) (*FinancialReport, error) {
	if startDate.After(endDate) {
		return nil, errors.New("start date must be before end date")
	}

	return &FinancialReport{
		ID:          uuid.New().String(),
		ReportType:  reportType,
		BusinessID:  businessID,
		StartDate:   startDate,
		EndDate:     endDate,
		Granularity: granularity,
		Status:      ReportStatusPending,
		GeneratedAt: time.Now(),
		Metrics:     make(map[string]*moneyDomain.Money),
	}, nil
}

func (r *FinancialReport) AddMetric(name string, value *moneyDomain.Money) {
	if r.Metrics == nil {
		r.Metrics = make(map[string]*moneyDomain.Money)
	}
	r.Metrics[name] = value
}

func (r *FinancialReport) SetStatus(status ReportStatus) {
	r.Status = status
}
