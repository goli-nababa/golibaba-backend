package port

import (
	"bank_service/internal/services/analytics/domain"
	domain2 "bank_service/internal/services/business/domain"
	transactionDomain "bank_service/internal/services/transaction/domain"
	walletDomain "bank_service/internal/services/wallet/domain"
	"context"
	"time"
)

type AnalyticsService interface {
	TrackTransaction(ctx context.Context, tx *transactionDomain.Transaction) error

	GenerateBusinessReport(ctx context.Context, businessID uint64, startDate, endDate time.Time) (*domain.AnalyticsReport, error)

	GenerateBusinessStats(ctx context.Context, businessID uint64, startDate, endDate time.Time) (*domain.BusinessStats, error)
	GetCommissionHistory(ctx context.Context, businessID uint64, startDate, endDate time.Time) ([]*domain.CommissionEntry, error)
	GetPayoutHistory(ctx context.Context, walletID walletDomain.WalletID, startDate, endDate time.Time) ([]*domain2.PayoutEntry, error)
}
