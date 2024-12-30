package service

import (
	moneyDomain "bank_service/internal/common/types"
	"bank_service/internal/services/financial_report/domain"
	"bank_service/internal/services/financial_report/port"
	transactionDomain "bank_service/internal/services/transaction/domain"
	transactionPort "bank_service/internal/services/transaction/port"
	walletDomain "bank_service/internal/services/wallet/domain"
	"bank_service/pkg/cache"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type financialReportService struct {
	reportRepo   port.ReportRepository
	txRepo       transactionPort.TransactionRepository
	cache        cache.Provider
	exportHelper exportHelper
}

func NewFinancialReportService(
	reportRepo port.ReportRepository,
	txRepo transactionPort.TransactionRepository,
	cache cache.Provider,
) port.Service {
	return &financialReportService{
		reportRepo:   reportRepo,
		txRepo:       txRepo,
		cache:        cache,
		exportHelper: newExportHelper(),
	}
}

func (s *financialReportService) GenerateDailyReport(ctx context.Context, businessID uint64, date time.Time) (*domain.FinancialReport, error) {
	startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	endOfDay := startOfDay.Add(24 * time.Hour).Add(-time.Second)

	return s.generateReport(ctx, businessID, startOfDay, endOfDay, domain.GranularityHourly)
}

func (s *financialReportService) GenerateMonthlyReport(ctx context.Context, businessID uint64, year int, month time.Month) (*domain.FinancialReport, error) {
	startOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	endOfMonth := startOfMonth.AddDate(0, 1, 0).Add(-time.Second)

	return s.generateReport(ctx, businessID, startOfMonth, endOfMonth, domain.GranularityDaily)
}

func (s *financialReportService) GenerateCustomReport(ctx context.Context, businessID uint64, startDate, endDate time.Time) (*domain.FinancialReport, error) {
	cacheKey := fmt.Sprintf("report:%d:%s:%s", businessID, startDate.Format("2006-01-02"), endDate.Format("2006-01-02"))

	if cached, err := s.cache.Get(ctx, cacheKey); err == nil {
		var report domain.FinancialReport
		if err := json.Unmarshal(cached, &report); err == nil {
			return &report, nil
		}
	}

	report, err := s.generateReport(ctx, businessID, startDate, endDate, domain.GranularityDaily)
	if err != nil {
		return nil, err
	}

	if data, err := json.Marshal(report); err == nil {
		s.cache.Set(ctx, cacheKey, time.Hour*24, data)
	}

	return report, nil
}

func (s *financialReportService) GetReportByID(ctx context.Context, reportID string) (*domain.FinancialReport, error) {
	cacheKey := fmt.Sprintf("report:%s", reportID)

	if cached, err := s.cache.Get(ctx, cacheKey); err == nil {
		var report domain.FinancialReport
		if err := json.Unmarshal(cached, &report); err == nil {
			return &report, nil
		}
	}

	report, err := s.reportRepo.GetByID(ctx, reportID)
	if err != nil {
		return nil, fmt.Errorf("failed to get report: %w", err)
	}

	if data, err := json.Marshal(report); err == nil {
		s.cache.Set(ctx, cacheKey, time.Hour*24, data)
	}

	return report, nil
}

func (s *financialReportService) ExportReport(ctx context.Context, reportID string, format domain.ReportFormat) ([]byte, error) {
	report, err := s.GetReportByID(ctx, reportID)
	if err != nil {
		return nil, err
	}

	return s.exportHelper.exportToFormat(report, format)
}

func (s *financialReportService) generateReport(
	ctx context.Context,
	businessID uint64,
	startDate, endDate time.Time,
	granularity domain.TimeGranularity,
) (*domain.FinancialReport, error) {
	report, err := domain.NewFinancialReport(
		domain.ReportTypeRevenue,
		businessID,
		startDate,
		endDate,
		granularity,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create report: %w", err)
	}

	filter := &transactionPort.TransactionFilter{
		WalletID: walletDomain.WalletID(fmt.Sprintf("%d", businessID)), // Convert businessID to WalletID
		From:     startDate,
		To:       endDate,
	}

	transactions, err := s.txRepo.FindByWallet(ctx, walletDomain.WalletID(fmt.Sprintf("%d", businessID)), filter)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch transactions: %w", err)
	}

	if err := s.calculateMetrics(report, transactions); err != nil {
		return nil, fmt.Errorf("failed to calculate metrics: %w", err)
	}

	report.SetStatus(domain.ReportStatusGenerated)

	if err := s.reportRepo.Save(ctx, report); err != nil {
		return nil, fmt.Errorf("failed to save report: %w", err)
	}

	return report, nil
}

func (s *financialReportService) calculateMetrics(report *domain.FinancialReport, transactions []*transactionDomain.Transaction) error {
	totalRevenue := int64(0)
	successfulCount := 0
	failedCount := 0

	for _, tx := range transactions {
		switch tx.Status {
		case transactionDomain.TransactionStatusSuccess:
			totalRevenue += tx.Amount.Amount
			successfulCount++
		case transactionDomain.TransactionStatusFailed:
			failedCount++
		}
	}

	report.AddMetric("total_revenue", &moneyDomain.Money{Amount: totalRevenue, Currency: "IRR"})
	report.AddMetric("successful_transactions", &moneyDomain.Money{Amount: int64(successfulCount), Currency: "COUNT"})
	report.AddMetric("failed_transactions", &moneyDomain.Money{Amount: int64(failedCount), Currency: "COUNT"})

	if successfulCount > 0 {
		avgRevenue := totalRevenue / int64(successfulCount)
		report.AddMetric("average_revenue", &moneyDomain.Money{Amount: avgRevenue, Currency: "IRR"})
	}

	return nil
}
