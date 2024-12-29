package port

import (
	"bank_service/internal/services/financial_report/domain"
	"context"
	"time"
)

type Service interface {
	GenerateDailyReport(ctx context.Context, businessID uint64, date time.Time) (*domain.FinancialReport, error)
	GenerateMonthlyReport(ctx context.Context, businessID uint64, year int, month time.Month) (*domain.FinancialReport, error)
	GenerateCustomReport(ctx context.Context, businessID uint64, startDate, endDate time.Time) (*domain.FinancialReport, error)
	GetReportByID(ctx context.Context, reportID string) (*domain.FinancialReport, error)
	ExportReport(ctx context.Context, reportID string, format domain.ReportFormat) ([]byte, error)
}

type ReportRepository interface {
	Save(ctx context.Context, report *domain.FinancialReport) error
	GetByID(ctx context.Context, id string) (*domain.FinancialReport, error)
	GetByBusinessID(ctx context.Context, businessID uint64, startDate, endDate time.Time) ([]*domain.FinancialReport, error)
	Update(ctx context.Context, report *domain.FinancialReport) error
}
