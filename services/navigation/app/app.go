package app

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"navigation_service/config"
	"navigation_service/internal/location"
	locationPort "navigation_service/internal/location/port"
	routeSystem "navigation_service/internal/routing"
	routePort "navigation_service/internal/routing/port"
	redisAdapter "navigation_service/pkg/adapters/cache"
	"navigation_service/pkg/adapters/storage"
	"navigation_service/pkg/adapters/storage/migrations"
	"navigation_service/pkg/cache"
	appCtx "navigation_service/pkg/context"
	"navigation_service/pkg/postgres"
)

type app struct {
	db              *gorm.DB
	cfg             config.Config
	redisProvider   cache.Provider
	locationService locationPort.Service
	routeService    routePort.Service
}

func (a *app) Config() config.Config {
	return a.cfg
}

func (a *app) DB() *gorm.DB {
	return a.db
}

func (a *app) LocationService(ctx context.Context) locationPort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.locationService == nil {
			a.locationService = a.locationServiceWithDB(a.db)
		}
		return a.locationService
	}

	return a.locationServiceWithDB(db)
}

func (a *app) locationServiceWithDB(db *gorm.DB) locationPort.Service {
	return location.NewService(storage.NewLocationRepository(db))
}

func (a *app) routeServiceWithDB(db *gorm.DB) routePort.Service {
	locationRepo := storage.NewLocationRepository(db)
	routeRepo := storage.NewRoutingRepository(db)
	return routeSystem.NewService(routeRepo, locationRepo)
}

func (a *app) RoteService(ctx context.Context) routePort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.routeService == nil {
			a.routeService = a.routeServiceWithDB(a.db)
		}
		return a.routeService
	}

	return a.routeServiceWithDB(db)
}

func (a *app) setDB() error {
	db, err := postgres.NewPsqlGormConnection(postgres.DBConnOptions{
		Host:   a.cfg.DB.Host,
		Port:   a.cfg.DB.Port,
		User:   a.cfg.DB.User,
		Pass:   a.cfg.DB.Password,
		Name:   a.cfg.DB.Database,
		Schema: a.cfg.DB.Schema,
	})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// Run migrations
	migrationManager := migrations.NewManager(db)
	if err := migrationManager.RunMigrations(); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	a.db = db
	return nil
}

func (a *app) setRedis() {
	a.redisProvider = redisAdapter.NewRedisProvider(fmt.Sprintf("%s:%d", a.cfg.Redis.Host, a.cfg.Redis.Port))
}

func NewApp(cfg config.Config) (App, error) {
	a := &app{
		cfg: cfg,
	}
	if err := a.setDB(); err != nil {
		return nil, err
	}
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
