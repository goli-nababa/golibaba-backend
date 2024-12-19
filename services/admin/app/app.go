package app

import (
	"admin/config"
	admin "admin/internal/admin"
	adminPort "admin/internal/admin/port"
	"admin/pkg/adapters/storage"
	"admin/pkg/postgres"
	"context"

	appCtx "admin/pkg/context"

	"gorm.io/gorm"
)

type app struct {
	db           *gorm.DB
	cfg          config.Config
	adminService adminPort.Service
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
func (a *app) adminServiceWithDB(db *gorm.DB) adminPort.Service {
	return admin.NewService(storage.NewAdminRepo(db))
}

func NewApp(cfg config.Config) (App, error) {
	a := &app{
		cfg: cfg,
	}

	if err := a.setDB(); err != nil {
		return nil, err
	}

	return a, nil
}
func (a *app) DB() *gorm.DB {
	return a.db
}

func (a *app) AdminService(ctx context.Context) adminPort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.adminService == nil {
			a.adminService = a.adminServiceWithDB(a.db)
		}
		return a.adminService
	}

	return a.adminServiceWithDB(db)
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
