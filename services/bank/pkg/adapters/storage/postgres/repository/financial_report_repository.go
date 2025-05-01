package repository

import (
	moneyDomain "bank_service/internal/common/types"
	"bank_service/internal/services/financial_report/domain"
	financialModel "bank_service/pkg/adapters/storage/postgres/model"
	"context"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type FinancialReportRepository struct {
	BaseRepository
}

func NewFinancialReportRepository(db *gorm.DB) *FinancialReportRepository {
	return &FinancialReportRepository{

		BaseRepository: NewBaseRepository(db),
	}
}

func (r *FinancialReportRepository) Save(ctx context.Context, report *domain.FinancialReport) error {
	metrics, err := json.Marshal(report.Metrics)
	if err != nil {
		return fmt.Errorf("failed to marshal metrics: %w", err)
	}

	model := &financialModel.FinancialReportModel{
		ID:          report.ID,
		ReportType:  string(report.ReportType),
		BusinessID:  report.BusinessID,
		StartDate:   report.StartDate,
		EndDate:     report.EndDate,
		Granularity: string(report.Granularity),
		Status:      string(report.Status),
		Metrics:     metrics,
		GeneratedAt: report.GeneratedAt,
		GeneratedBy: report.GeneratedBy,
		Format:      string(report.Format),
		DataURL:     report.DataURL,
	}

	return r.DB(ctx).Create(model).Error
}

func (r *FinancialReportRepository) Update(ctx context.Context, report *domain.FinancialReport) error {
	metrics, err := json.Marshal(report.Metrics)
	if err != nil {
		return fmt.Errorf("failed to marshal metrics: %w", err)
	}

	model := &financialModel.FinancialReportModel{
		ID:          report.ID,
		ReportType:  string(report.ReportType),
		BusinessID:  report.BusinessID,
		StartDate:   report.StartDate,
		EndDate:     report.EndDate,
		Granularity: string(report.Granularity),
		Status:      string(report.Status),
		Metrics:     metrics,
		GeneratedAt: report.GeneratedAt,
		GeneratedBy: report.GeneratedBy,
		Format:      string(report.Format),
		DataURL:     report.DataURL,
	}

	return r.DB(ctx).Where("id = ?", report.ID).Updates(model).Error
}

func (r *FinancialReportRepository) GetByID(ctx context.Context, id string) (*domain.FinancialReport, error) {
	var model financialModel.FinancialReportModel
	if err := r.DB(ctx).First(&model, "id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("failed to get report: %w", err)
	}

	var metrics map[string]*moneyDomain.Money
	if err := json.Unmarshal(model.Metrics, &metrics); err != nil {
		return nil, fmt.Errorf("failed to unmarshal metrics: %w", err)
	}

	return &domain.FinancialReport{
		ID:          model.ID,
		ReportType:  domain.ReportType(model.ReportType),
		BusinessID:  model.BusinessID,
		StartDate:   model.StartDate,
		EndDate:     model.EndDate,
		Granularity: domain.TimeGranularity(model.Granularity),
		Status:      domain.ReportStatus(model.Status),
		Metrics:     metrics,
		GeneratedAt: model.GeneratedAt,
		GeneratedBy: model.GeneratedBy,
		Format:      domain.ReportFormat(model.Format),
		DataURL:     model.DataURL,
	}, nil
}

func (r *FinancialReportRepository) GetByBusinessID(ctx context.Context, businessID uint64, startDate, endDate time.Time) ([]*domain.FinancialReport, error) {
	var models []financialModel.FinancialReportModel
	if err := r.DB(ctx).
		Where("business_id = ? AND start_date >= ? AND end_date <= ?", businessID, startDate, endDate).
		Order("generated_at DESC").
		Find(&models).Error; err != nil {
		return nil, fmt.Errorf("failed to get business reports: %w", err)
	}

	reports := make([]*domain.FinancialReport, len(models))
	for i, model := range models {
		var metrics map[string]*moneyDomain.Money
		if err := json.Unmarshal(model.Metrics, &metrics); err != nil {
			return nil, fmt.Errorf("failed to unmarshal metrics: %w", err)
		}

		reports[i] = &domain.FinancialReport{
			ID:          model.ID,
			ReportType:  domain.ReportType(model.ReportType),
			BusinessID:  model.BusinessID,
			StartDate:   model.StartDate,
			EndDate:     model.EndDate,
			Granularity: domain.TimeGranularity(model.Granularity),
			Status:      domain.ReportStatus(model.Status),
			Metrics:     metrics,
			GeneratedAt: model.GeneratedAt,
			GeneratedBy: model.GeneratedBy,
			Format:      domain.ReportFormat(model.Format),
			DataURL:     model.DataURL,
		}
	}

	return reports, nil
}

func (r *FinancialReportRepository) DeleteReport(ctx context.Context, id string) error {
	return r.DB(ctx).Delete(&financialModel.FinancialReportModel{}, "id = ?", id).Error
}
