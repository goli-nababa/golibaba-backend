package app

import (
	"bank_service/config"
	analyticsPort "bank_service/internal/services/analytics/port"
	businessPort "bank_service/internal/services/business/port"
	commissionPort "bank_service/internal/services/commission/port"
	financialReportPort "bank_service/internal/services/financial_report/port"
	notificationPort "bank_service/internal/services/notification/port"
	paymentPort "bank_service/internal/services/payment/port"
	transactionPort "bank_service/internal/services/transaction/port"
	walletPort "bank_service/internal/services/wallet/port"
	"context"
	"gorm.io/gorm"
)

type App interface {
	WalletService(ctx context.Context) walletPort.Service
	TransactionService(ctx context.Context) transactionPort.TransactionService
	PaymentService(ctx context.Context) paymentPort.Service
	BusinessService(ctx context.Context) businessPort.Service
	CommissionService(ctx context.Context) commissionPort.Service
	NotificationService(ctx context.Context) notificationPort.Service
	AnalyticsService(ctx context.Context) analyticsPort.AnalyticsService
	FinancialReportService(ctx context.Context) financialReportPort.Service
	DB() *gorm.DB
	Config() config.Config
}
