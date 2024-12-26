package app

import (
	"context"
	"transportation/config"
	"transportation/internal/company"
	companyPort "transportation/internal/company/port"
	"transportation/internal/trip"
	tripPort "transportation/internal/trip/port"

	"transportation/pkg/adapters/storage"
	"transportation/pkg/adapters/storage/migrations"
	"transportation/pkg/logging"
	"transportation/pkg/postgres"

	appCtx "transportation/pkg/context"

	"gorm.io/gorm"
)

type app struct {
	db             *gorm.DB
	cfg            config.Config
	logger         logging.Logger
	companyService companyPort.Service
	tripService    tripPort.Service
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
