package port

import (
	"bank_service/internal/services/analytics/domain"
	"context"
	"time"
)

type AnalyticsRepository interface {
	SaveReport(ctx context.Context, report *domain.AnalyticsReport) error
	GetReport(ctx context.Context, reportID string) (*domain.AnalyticsReport, error)
	GetBusinessReports(ctx context.Context, businessID uint64, startDate, endDate time.Time) ([]*domain.AnalyticsReport, error)
	DeleteReport(ctx context.Context, reportID string) error
	DeleteBusinessReports(ctx context.Context, businessID uint64) error
}
