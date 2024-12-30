package repository

import (
	analyticsModel "bank_service/pkg/adapters/storage/postgres/model"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"bank_service/internal/services/analytics/domain"
	"gorm.io/gorm"
)

type AnalyticsRepository struct {
	BaseRepository
}

func NewAnalyticsRepository(db *gorm.DB) *AnalyticsRepository {
	return &AnalyticsRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

func (r *AnalyticsRepository) SaveReport(ctx context.Context, report *domain.AnalyticsReport) error {
	metrics, err := json.Marshal(report.Metrics)
	if err != nil {
		return fmt.Errorf("failed to marshal metrics: %w", err)
	}

	trends, err := json.Marshal(report.Trends)
	if err != nil {
		return fmt.Errorf("failed to marshal trends: %w", err)
	}

	comparisons, err := json.Marshal(report.Comparisons)
	if err != nil {
		return fmt.Errorf("failed to marshal comparisons: %w", err)
	}

	model := &analyticsModel.AnalyticsModel{
		ID:          report.ID,
		BusinessID:  report.BusinessID,
		StartDate:   report.StartDate,
		EndDate:     report.EndDate,
		Metrics:     metrics,
		Trends:      trends,
		Comparisons: comparisons,
		GeneratedAt: report.GeneratedAt,
	}

	return r.DB(ctx).Create(model).Error
}

func (r *AnalyticsRepository) GetReport(ctx context.Context, reportID string) (*domain.AnalyticsReport, error) {
	var model analyticsModel.AnalyticsModel
	if err := r.DB(ctx).First(&model, "id = ?", reportID).Error; err != nil {
		return nil, fmt.Errorf("failed to get report: %w", err)
	}

	var metrics map[string]float64
	if err := json.Unmarshal(model.Metrics, &metrics); err != nil {
		return nil, fmt.Errorf("failed to unmarshal metrics: %w", err)
	}

	var trends map[string][]domain.DataPoint
	if err := json.Unmarshal(model.Trends, &trends); err != nil {
		return nil, fmt.Errorf("failed to unmarshal trends: %w", err)
	}

	var comparisons map[string]domain.Comparison
	if err := json.Unmarshal(model.Comparisons, &comparisons); err != nil {
		return nil, fmt.Errorf("failed to unmarshal comparisons: %w", err)
	}

	return &domain.AnalyticsReport{
		ID:          model.ID,
		BusinessID:  model.BusinessID,
		StartDate:   model.StartDate,
		EndDate:     model.EndDate,
		Metrics:     metrics,
		Trends:      trends,
		Comparisons: comparisons,
		GeneratedAt: model.GeneratedAt,
	}, nil
}

func (r *AnalyticsRepository) GetBusinessReports(ctx context.Context, businessID uint64, startDate, endDate time.Time) ([]*domain.AnalyticsReport, error) {
	var models []analyticsModel.AnalyticsModel
	if err := r.DB(ctx).
		Where("business_id = ? AND start_date >= ? AND end_date <= ?", businessID, startDate, endDate).
		Order("generated_at DESC").
		Find(&models).Error; err != nil {
		return nil, fmt.Errorf("failed to get business reports: %w", err)
	}

	reports := make([]*domain.AnalyticsReport, len(models))
	for i, model := range models {
		var metrics map[string]float64
		if err := json.Unmarshal(model.Metrics, &metrics); err != nil {
			return nil, fmt.Errorf("failed to unmarshal metrics: %w", err)
		}

		var trends map[string][]domain.DataPoint
		if err := json.Unmarshal(model.Trends, &trends); err != nil {
			return nil, fmt.Errorf("failed to unmarshal trends: %w", err)
		}

		var comparisons map[string]domain.Comparison
		if err := json.Unmarshal(model.Comparisons, &comparisons); err != nil {
			return nil, fmt.Errorf("failed to unmarshal comparisons: %w", err)
		}

		reports[i] = &domain.AnalyticsReport{
			ID:          model.ID,
			BusinessID:  model.BusinessID,
			StartDate:   model.StartDate,
			EndDate:     model.EndDate,
			Metrics:     metrics,
			Trends:      trends,
			Comparisons: comparisons,
			GeneratedAt: model.GeneratedAt,
		}
	}

	return reports, nil
}

func (r *AnalyticsRepository) DeleteReport(ctx context.Context, reportID string) error {
	return r.DB(ctx).Delete(&analyticsModel.AnalyticsModel{}, "id = ?", reportID).Error
}

func (r *AnalyticsRepository) DeleteBusinessReports(ctx context.Context, businessID uint64) error {
	return r.DB(ctx).Delete(&analyticsModel.AnalyticsModel{}, "business_id = ?", businessID).Error
}
