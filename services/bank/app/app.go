package app

import (
	"bank_service/pkg/adapters/events"
	"bank_service/pkg/adapters/payment/zarinpal"
	"bank_service/pkg/adapters/storage/migrations"
	"context"
	"fmt"

	"bank_service/config"
	"bank_service/internal/services/analytics"
	"bank_service/internal/services/business"
	service3 "bank_service/internal/services/commission"
	service4 "bank_service/internal/services/financial_report"
	"bank_service/internal/services/notification"
	"bank_service/internal/services/payment"
	service2 "bank_service/internal/services/transaction"
	service "bank_service/internal/services/wallet"

	analyticsPort "bank_service/internal/services/analytics/port"
	businessPort "bank_service/internal/services/business/port"
	commissionPort "bank_service/internal/services/commission/port"
	financialReportPort "bank_service/internal/services/financial_report/port"
	notificationPort "bank_service/internal/services/notification/port"
	paymentPort "bank_service/internal/services/payment/port"
	transactionPort "bank_service/internal/services/transaction/port"
	walletPort "bank_service/internal/services/wallet/port"

	cacheAdapter "bank_service/pkg/adapters/cache"
	"bank_service/pkg/adapters/storage/postgres/repository"
	"bank_service/pkg/cache"
	appCtx "bank_service/pkg/context"
	"bank_service/pkg/logging"
	"bank_service/pkg/postgres"

	"gorm.io/gorm"
)

type app struct {
	db                     *gorm.DB
	cfg                    config.Config
	logger                 logging.Logger
	walletService          walletPort.Service
	transactionService     transactionPort.TransactionService
	paymentService         paymentPort.Service
	businessService        businessPort.Service
	commissionService      commissionPort.Service
	notificationService    notificationPort.Service
	analyticsService       analyticsPort.AnalyticsService
	financialReportService financialReportPort.Service
	redisProvider          cache.Provider
}

func (a *app) DB() *gorm.DB {
	return a.db
}

func (a *app) Config() config.Config {
	return a.cfg
}

func (a *app) WalletService(ctx context.Context) walletPort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.walletService == nil {
			a.walletService = a.walletServiceWithDB(a.db)
		}
		return a.walletService
	}
	return a.walletServiceWithDB(db)
}

func (a *app) TransactionService(ctx context.Context) transactionPort.TransactionService {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.transactionService == nil {
			a.transactionService = a.transactionServiceWithDB(a.db)
		}
		return a.transactionService
	}
	return a.transactionServiceWithDB(db)
}

func (a *app) PaymentService(ctx context.Context) paymentPort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.paymentService == nil {
			a.paymentService = a.paymentServiceWithDB(a.db)
		}
		return a.paymentService
	}
	return a.paymentServiceWithDB(db)
}

func (a *app) BusinessService(ctx context.Context) businessPort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.businessService == nil {
			a.businessService = a.businessServiceWithDB(a.db)
		}
		return a.businessService
	}
	return a.businessServiceWithDB(db)
}

func (a *app) CommissionService(ctx context.Context) commissionPort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.commissionService == nil {
			a.commissionService = a.commissionServiceWithDB(a.db)
		}
		return a.commissionService
	}
	return a.commissionServiceWithDB(db)
}

func (a *app) NotificationService(ctx context.Context) notificationPort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.notificationService == nil {
			a.notificationService = a.notificationServiceWithDB(a.db)
		}
		return a.notificationService
	}
	return a.notificationServiceWithDB(db)
}

func (a *app) AnalyticsService(ctx context.Context) analyticsPort.AnalyticsService {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.analyticsService == nil {
			a.analyticsService = a.analyticsServiceWithDB(a.db)
		}
		return a.analyticsService
	}
	return a.analyticsServiceWithDB(db)
}

func (a *app) FinancialReportService(ctx context.Context) financialReportPort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.financialReportService == nil {
			a.financialReportService = a.financialReportServiceWithDB(a.db)
		}
		return a.financialReportService
	}
	return a.financialReportServiceWithDB(db)
}

func (a *app) walletServiceWithDB(db *gorm.DB) walletPort.Service {
	eventPub, err := events.NewEventPublisher("amqp://guest:guest@localhost:5672/", "wallet_events")
	if err != nil {
		panic(err)
	}
	return service.NewWalletService(
		repository.NewPostgresWalletRepository(db),
		cacheAdapter.NewWalletCache(a.redisProvider, a.logger),
		cacheAdapter.NewRedisWalletLocker(a.redisProvider),
		eventPub,
	)
}

func (a *app) transactionServiceWithDB(db *gorm.DB) transactionPort.TransactionService {
	eventPub, err := events.NewEventPublisher("amqp://guest:guest@localhost:5672/", "transaction_events")
	if err != nil {
		panic(err)
	}
	return service2.NewTransactionService(
		repository.NewTransactionRepository(db),
		cacheAdapter.NewTransactionLocker(a.redisProvider),
		eventPub,
	)
}

func (a *app) paymentServiceWithDB(db *gorm.DB) paymentPort.Service {

	gateway := zarinpal.NewZarinpalGateway("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx", fmt.Sprintf("http://localhost:%v/api/v1/payments/callback", a.cfg.Server.HTTPPort))

	return payment.NewPaymentService(
		repository.NewTransactionRepository(db),
		a.WalletService(context.Background()),
		gateway,
		a.CommissionService(context.Background()),
		a.redisProvider,
	)
}

func (a *app) businessServiceWithDB(db *gorm.DB) businessPort.Service {
	return business.NewBusinessService(
		repository.NewBusinessRepository(db),
		a.CommissionService(context.Background()),
		a.PaymentService(context.Background()),
		a.AnalyticsService(context.Background()),
		a.redisProvider,
	)
}

func (a *app) commissionServiceWithDB(db *gorm.DB) commissionPort.Service {
	rate := service3.NewCommissionRateProvider()
	return service3.NewCommissionService(
		repository.NewCommissionRepository(db),
		rate,
		a.WalletService(context.Background()),
		a.redisProvider,
	)
}

func (a *app) notificationServiceWithDB(db *gorm.DB) notificationPort.Service {
	return notification.NewService()
}

func (a *app) analyticsServiceWithDB(db *gorm.DB) analyticsPort.AnalyticsService {
	return analytics.NewAnalyticsService(
		repository.NewTransactionRepository(db),
		repository.NewCommissionRepository(db),
		repository.NewAnalyticsRepository(db),
	)
}

func (a *app) financialReportServiceWithDB(db *gorm.DB) financialReportPort.Service {
	return service4.NewFinancialReportService(
		repository.NewFinancialReportRepository(db),
		repository.NewTransactionRepository(db),
		a.redisProvider,
	)
}

func (a *app) setDB() error {
	db, err := postgres.NewPsqlGormConnection(postgres.DBConnOptions{
		User:   a.cfg.DB.User,
		Pass:   a.cfg.DB.Password,
		Host:   a.cfg.DB.Host,
		Port:   a.cfg.DB.Port,
		Name:   a.cfg.DB.Database,
		Schema: a.cfg.DB.Schema,
	})

	if err != nil {
		return err
	}

	a.db = db
	return nil
}

func (a *app) setRedis() {
	a.redisProvider = cacheAdapter.NewRedisProvider(fmt.Sprintf("%s:%d", a.cfg.Redis.Host, a.cfg.Redis.Port))
}

func (a *app) migration() {
	err := migrations.Migrate(a.db)
	if err != nil {
		return
	}
}

func NewApp(cfg config.Config) (App, error) {
	a := &app{
		cfg:    cfg,
		logger: logging.NewLogger(&cfg),
	}

	a.logger.Init()

	if err := a.setDB(); err != nil {
		return nil, err
	}
	a.migration()

	a.setRedis()

	return a, nil
}

func NewMustApp(cfg config.Config) App {
	app, err := NewApp(cfg)
	if err != nil {
		panic(err)
	}
	return app
}
