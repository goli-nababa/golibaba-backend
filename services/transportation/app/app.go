package app

import (
	"context"
	"fmt"
	"transportation/config"
	"transportation/internal/company"
	companyPort "transportation/internal/company/port"
	"transportation/internal/trip"
	tripPort "transportation/internal/trip/port"
	vehicleRequestPoller "transportation/internal/vehicle_request_poller"
	vehicleRequestPollerPort "transportation/internal/vehicle_request_poller/port"

	"transportation/pkg/adapters/queue"
	"transportation/pkg/adapters/storage"
	"transportation/pkg/adapters/storage/migrations"
	"transportation/pkg/logging"
	"transportation/pkg/postgres"

	appCtx "transportation/pkg/context"

	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

type app struct {
	db                          *gorm.DB
	mqConn                      *amqp.Connection
	cfg                         config.Config
	logger                      logging.Logger
	companyService              companyPort.Service
	vehicleRequestPollerService vehicleRequestPollerPort.Service
	tripService                 tripPort.Service
}

func (a *app) setDB() error {
	db, err := postgres.NewPsqlGormConnection(postgres.DBConnOptions{
		User:   a.cfg.DB.User,
		Pass:   a.cfg.DB.Pass,
		Host:   a.cfg.DB.Host,
		Port:   a.cfg.DB.Port,
		Name:   a.cfg.DB.Name,
		Schema: a.cfg.DB.Schema,
	})

	if err != nil {
		return err
	}

	a.db = db
	return nil
}
func (a *app) companyServiceWithDB(db *gorm.DB) companyPort.Service {
	return company.NewService(storage.NewCompanyRepo(db), a.logger)
}

func (a *app) vehicleRequestPollerServiceWithRequirements() vehicleRequestPollerPort.Service {
	return vehicleRequestPoller.NewService(
		queue.NewVehicleRequestQueueRepo(a.mqConn, a.cfg.MessageQueue.VehicleRequestQueueName),
		storage.NewTripRepo(a.db))
}

func (a *app) tripServiceWithDB(db *gorm.DB) tripPort.Service {
	return trip.NewTripService(storage.NewTripRepo(db))
}

func NewApp(cfg config.Config) (App, error) {

	logger := logging.NewLogger(&cfg)
	a := &app{
		cfg:    cfg,
		logger: logger,
	}

	if err := a.setDB(); err != nil {
		return nil, err
	}

	err := migrations.AutoMigrate(a.db)
	if err != nil {
		return nil, err
	}

	mqConn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", cfg.MessageQueue.RabbitMqUsername, cfg.MessageQueue.RabbitMqPassword, cfg.MessageQueue.RabbitMqHost, cfg.MessageQueue.RabbitMqPort))
	if err != nil {
		return nil, err
	}
	a.mqConn = mqConn

	return a, err
}
func (a *app) DB() *gorm.DB {
	return a.db
}

func (a *app) CompanyService(ctx context.Context) companyPort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.companyService == nil {
			a.companyService = a.companyServiceWithDB(a.db)
		}
		return a.companyService
	}

	return a.companyServiceWithDB(db)
}

func (a *app) VehicleRequestPollerService(ctx context.Context) vehicleRequestPollerPort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.vehicleRequestPollerService == nil {
			a.vehicleRequestPollerService = a.vehicleRequestPollerServiceWithRequirements()
		}
		return a.vehicleRequestPollerService
	}

	return a.vehicleRequestPollerServiceWithRequirements()
}

func (a *app) TripService(ctx context.Context) tripPort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.tripService == nil {
			a.tripService = a.tripServiceWithDB(a.db)
		}
		return a.tripService
	}

	return a.tripServiceWithDB(db)
}

func (a *app) Config() config.Config {
	return a.cfg
}

func NewMustApp(cfg config.Config) App {
	app, err := NewApp(cfg)
	if err != nil {
		panic(err)
	}
	return app
}
