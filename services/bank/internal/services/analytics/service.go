package analytics

import (
	"bank_service/internal/common/types"
	"bank_service/internal/services/analytics/domain"
	"bank_service/internal/services/analytics/port"
	domain2 "bank_service/internal/services/business/domain"
	commissionDomain "bank_service/internal/services/commission/domain"
	commissionPort "bank_service/internal/services/commission/port"
	transactionDomain "bank_service/internal/services/transaction/domain"
	transactionPort "bank_service/internal/services/transaction/port"
	walletDomain "bank_service/internal/services/wallet/domain"
	"context"
	"fmt"
	"github.com/google/uuid"
	"sort"
	"strconv"
	"strings"
	"time"
)

type analyticsService struct {
	txRepo         transactionPort.TransactionRepository
	commissionRepo commissionPort.CommissionRepository
	analyticsRepo  port.AnalyticsRepository
}

func NewAnalyticsService(
	txRepo transactionPort.TransactionRepository,
	commissionRepo commissionPort.CommissionRepository,
	analyticsRepo port.AnalyticsRepository,
) port.AnalyticsService {
	return &analyticsService{
		txRepo:         txRepo,
		commissionRepo: commissionRepo,
		analyticsRepo:  analyticsRepo,
	}
}

func (s *analyticsService) TrackTransaction(ctx context.Context, tx *transactionDomain.Transaction) error {
	if tx == nil {
		return fmt.Errorf("cannot track nil transaction")
	}

	metrics, err := s.txRepo.FindByWallet(ctx, tx.ToWalletID, &transactionPort.TransactionFilter{
		From: time.Now().AddDate(0, -1, 0),
		To:   time.Now(),
	})
	if err != nil {
		return fmt.Errorf("failed to get wallet metrics: %w", err)
	}

	var (
		totalAmount     int64
		successfulCount int64
		failedCount     int64
	)

	for _, metric := range metrics {
		switch metric.Status {
		case transactionDomain.TransactionStatusSuccess:
			totalAmount += metric.Amount.Amount
			successfulCount++
		case transactionDomain.TransactionStatusFailed:
			failedCount++
		}
	}

	commissions, err := s.commissionRepo.FindByBusinessIDAndDateRange(
		ctx,
		extractBusinessID(tx.ToWalletID),
		time.Now().AddDate(0, -1, 0),
		time.Now(),
	)
	if err != nil {
		return fmt.Errorf("failed to get commission data: %w", err)
	}

	var totalCommission int64
	for _, commission := range commissions {
		if commission.Status == commissionDomain.PaymentStatusProcessed {
			totalCommission += commission.Amount.Amount
		}
	}

	analyticsData := &domain.AnalyticsReport{
		ID:          uuid.New().String(),
		BusinessID:  extractBusinessID(tx.ToWalletID),
		StartDate:   time.Now().AddDate(0, -1, 0),
		EndDate:     time.Now(),
		GeneratedAt: time.Now(),
		Metrics: map[string]float64{
			"total_amount":     float64(totalAmount),
			"successful_count": float64(successfulCount),
			"failed_count":     float64(failedCount),
			"total_commission": float64(totalCommission),
			"success_rate":     calculateSuccessRate(successfulCount, failedCount),
			"average_amount":   calculateAverageAmount(totalAmount, successfulCount),
		},
		Trends: map[string][]domain.DataPoint{
			"daily_transactions": generateDailyTrends(metrics),
		},
		Comparisons: map[string]domain.Comparison{
			"previous_month": calculateMonthlyComparison(metrics),
		},
	}

	if err := s.analyticsRepo.SaveReport(ctx, analyticsData); err != nil {
		return fmt.Errorf("failed to save analytics data: %w", err)
	}

	return nil
}

func extractBusinessID(walletID walletDomain.WalletID) uint64 {
	parts := strings.Split(string(walletID), "_")
	if len(parts) != 2 {
		return 0
	}
	id, _ := strconv.ParseUint(parts[1], 10, 64)
	return id
}

func calculateSuccessRate(successful, failed int64) float64 {
	total := successful + failed
	if total == 0 {
		return 0
	}
	return float64(successful) / float64(total) * 100
}

func calculateAverageAmount(total, count int64) float64 {
	if count == 0 {
		return 0
	}
	return float64(total) / float64(count)
}

func generateDailyTrends(transactions []*transactionDomain.Transaction) []domain.DataPoint {
	dailyGroups := make(map[string]int64)
	for _, tx := range transactions {
		day := tx.CreatedAt.Format("2006-01-02")
		if tx.Status == transactionDomain.TransactionStatusSuccess {
			dailyGroups[day] += tx.Amount.Amount
		}
	}

	points := make([]domain.DataPoint, 0, len(dailyGroups))
	for day, amount := range dailyGroups {
		timestamp, _ := time.Parse("2006-01-02", day)
		points = append(points, domain.DataPoint{
			Timestamp: timestamp,
			Value:     float64(amount),
			Label:     day,
		})
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i].Timestamp.Before(points[j].Timestamp)
	})

	return points
}

func calculateMonthlyComparison(transactions []*transactionDomain.Transaction) domain.Comparison {
	now := time.Now()
	currentMonthStart := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
	lastMonthStart := currentMonthStart.AddDate(0, -1, 0)

	var currentMonthTotal, lastMonthTotal int64

	for _, tx := range transactions {
		if tx.Status != transactionDomain.TransactionStatusSuccess {
			continue
		}

		if tx.CreatedAt.After(currentMonthStart) {
			currentMonthTotal += tx.Amount.Amount
		} else if tx.CreatedAt.After(lastMonthStart) {
			lastMonthTotal += tx.Amount.Amount
		}
	}

	return domain.Comparison{
		CurrentValue:  float64(currentMonthTotal),
		PreviousValue: float64(lastMonthTotal),
		ChangePercent: calculateChangePercent(currentMonthTotal, lastMonthTotal),
	}
}

func calculateChangePercent(current, previous int64) float64 {
	if previous == 0 {
		if current == 0 {
			return 0
		}
		return 100
	}
	return (float64(current-previous) / float64(previous)) * 100
}

func (s *analyticsService) GenerateBusinessStats(ctx context.Context, businessID uint64, startDate, endDate time.Time) (*domain.BusinessStats, error) {
	transactions, err := s.txRepo.FindByBusinessIDAndDateRange(ctx, businessID, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch transactions: %w", err)
	}

	commissions, err := s.commissionRepo.FindByBusinessIDAndDateRange(ctx, businessID, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch commissions: %w", err)
	}

	totalRevenue, err := calculateTotalRevenue(transactions)
	if err != nil {
		return nil, err
	}

	totalCommission, err := calculateTotalCommission(commissions)
	if err != nil {
		return nil, err
	}

	totalPayouts, err := calculateTotalPayouts(transactions)
	if err != nil {
		return nil, err
	}

	successfulCount, failedCount := countTransactionsByStatus(transactions)

	avgOrderValue, err := calculateAverageOrderValue(transactions)
	if err != nil {
		return nil, err
	}

	avgCommissionRate := calculateAverageCommissionRate(commissions)

	return &domain.BusinessStats{
		TotalRevenue:      totalRevenue,
		TotalCommission:   totalCommission,
		TotalPayouts:      totalPayouts,
		TransactionCount:  int64(len(transactions)),
		SuccessfulTxCount: successfulCount,
		FailedTxCount:     failedCount,
		AverageOrderValue: avgOrderValue,
		CommissionRate:    avgCommissionRate,
		Period: struct {
			Start time.Time
			End   time.Time
		}{
			Start: startDate,
			End:   endDate,
		},
	}, nil
}

func (s *analyticsService) GenerateBusinessReport(ctx context.Context, businessID uint64, startDate, endDate time.Time) (*domain.AnalyticsReport, error) {
	stats, err := s.GenerateBusinessStats(ctx, businessID, startDate, endDate)
	if err != nil {
		return nil, err
	}

	report := &domain.AnalyticsReport{
		ID:          uuid.New().String(),
		BusinessID:  businessID,
		StartDate:   startDate,
		EndDate:     endDate,
		GeneratedAt: time.Now(),
		Metrics:     make(map[string]float64),
	}

	report.AddMetric("total_revenue", float64(stats.TotalRevenue.Amount))
	report.AddMetric("total_commission", float64(stats.TotalCommission.Amount))
	report.AddMetric("total_payouts", float64(stats.TotalPayouts.Amount))
	report.AddMetric("transaction_count", float64(stats.TransactionCount))
	report.AddMetric("successful_tx_count", float64(stats.SuccessfulTxCount))
	report.AddMetric("failed_tx_count", float64(stats.FailedTxCount))
	report.AddMetric("average_order_value", float64(stats.AverageOrderValue.Amount))
	report.AddMetric("commission_rate", stats.CommissionRate)

	err = s.analyticsRepo.SaveReport(ctx, report)
	if err != nil {
		return nil, err
	}

	return report, nil
}

func (s *analyticsService) GetCommissionHistory(ctx context.Context, businessID uint64, startDate, endDate time.Time) ([]*domain.CommissionEntry, error) {
	commissions, err := s.commissionRepo.FindByBusinessIDAndDateRange(ctx, businessID, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch commissions: %w", err)
	}

	entries := make([]*domain.CommissionEntry, len(commissions))
	for i, commission := range commissions {
		entries[i] = &domain.CommissionEntry{
			TransactionID: commission.TransactionID,
			Amount:        commission.Amount,
			Rate:          commission.Rate,
			Status:        string(commission.Status),
			ProcessedAt:   commission.CreatedAt,
		}
	}

	return entries, nil
}

func (s *analyticsService) GetPayoutHistory(ctx context.Context, walletID walletDomain.WalletID, startDate, endDate time.Time) ([]*domain2.PayoutEntry, error) {
	transactions, err := s.txRepo.FindByWallet(ctx, walletID, &transactionPort.TransactionFilter{
		Status: []transactionDomain.TransactionStatus{transactionDomain.TransactionStatusSuccess},
		Types:  []transactionDomain.TransactionType{transactionDomain.TransactionTypeWithdrawal},
		From:   startDate,
		To:     endDate,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch payouts: %w", err)
	}

	entries := make([]*domain2.PayoutEntry, len(transactions))
	for i, tx := range transactions {
		entries[i] = &domain2.PayoutEntry{
			ID:            string(tx.ID),
			Amount:        tx.Amount,
			Status:        string(tx.Status),
			RequestedAt:   tx.CreatedAt,
			ProcessedAt:   tx.CompletedAt,
			ReferenceID:   tx.ReferenceID,
			FailureReason: tx.FailureReason,
		}
	}

	return entries, nil
}

func calculateTotalRevenue(transactions []*transactionDomain.Transaction) (*types.Money, error) {
	if len(transactions) == 0 {
		return types.NewMoney(0, "IRR")
	}

	total, err := types.NewMoney(0, transactions[0].Amount.Currency)
	if err != nil {
		return nil, err
	}

	for _, tx := range transactions {
		if tx.Status == transactionDomain.TransactionStatusSuccess {
			total, err = total.Add(tx.Amount)
			if err != nil {
				return nil, err
			}
		}
	}

	return total, nil
}

func calculateTotalCommission(commissions []*commissionDomain.Commission) (*types.Money, error) {
	if len(commissions) == 0 {
		return types.NewMoney(0, "IRR")
	}

	total, err := types.NewMoney(0, commissions[0].Amount.Currency)
	if err != nil {
		return nil, err
	}

	for _, commission := range commissions {
		if commission.Status == commissionDomain.PaymentStatusProcessed {
			total, err = total.Add(commission.Amount)
			if err != nil {
				return nil, err
			}
		}
	}

	return total, nil
}

func calculateTotalPayouts(transactions []*transactionDomain.Transaction) (*types.Money, error) {
	if len(transactions) == 0 {
		return types.NewMoney(0, "IRR")
	}

	total, err := types.NewMoney(0, transactions[0].Amount.Currency)
	if err != nil {
		return nil, err
	}

	for _, tx := range transactions {
		if tx.Type == transactionDomain.TransactionTypeWithdrawal &&
			tx.Status == transactionDomain.TransactionStatusSuccess {
			total, err = total.Add(tx.Amount)
			if err != nil {
				return nil, err
			}
		}
	}

	return total, nil
}

func calculateAverageOrderValue(transactions []*transactionDomain.Transaction) (*types.Money, error) {
	if len(transactions) == 0 {
		return types.NewMoney(0, "IRR")
	}

	total, err := calculateTotalRevenue(transactions)
	if err != nil {
		return nil, err
	}

	successfulTxCount := int64(0)
	for _, tx := range transactions {
		if tx.Status == transactionDomain.TransactionStatusSuccess {
			successfulTxCount++
		}
	}

	if successfulTxCount == 0 {
		return types.NewMoney(0, transactions[0].Amount.Currency)
	}

	return total.Divide(float64(successfulTxCount))
}

func calculateAverageCommissionRate(commissions []*commissionDomain.Commission) float64 {
	if len(commissions) == 0 {
		return 0
	}

	totalRate := 0.0
	count := 0

	for _, commission := range commissions {
		if commission.Status == commissionDomain.PaymentStatusProcessed {
			totalRate += commission.Rate
			count++
		}
	}

	if count == 0 {
		return 0
	}

	return totalRate / float64(count)
}

func countTransactionsByStatus(transactions []*transactionDomain.Transaction) (successful, failed int64) {
	for _, tx := range transactions {
		switch tx.Status {
		case transactionDomain.TransactionStatusSuccess:
			successful++
		case transactionDomain.TransactionStatusFailed:
			failed++
		}
	}
	return
}
