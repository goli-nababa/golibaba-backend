package app

import (
	"context"
	"fmt"
	"github.com/goli-nababa/golibaba-backend/modules/cache"
	"user_service/config"
	"user_service/internal/user"
	userPort "user_service/internal/user/port"
	"user_service/pkg/adapters/storage"
	appCtx "user_service/pkg/context"
	"user_service/pkg/email"
	"user_service/pkg/postgres"

	"gorm.io/gorm"
)

type app struct {
	db          *gorm.DB
	cfg         config.Config
	redis       cache.Provider
	userService userPort.Service
	mailService email.Adapter
}

func (a *app) Config() config.Config {
	return a.cfg
}

func (a *app) DB() *gorm.DB {
	return a.db
}

func (a *app) userServiceWithDB(db *gorm.DB) userPort.Service {
	return user.NewService(storage.NewUserRepo(db, a.cfg.Server.Secret))
}

func (a *app) Cache() cache.Provider {
	return a.redis
}

func (a *app) MailService() email.Adapter {
	return a.mailService
}

func (a *app) UserService(ctx context.Context) userPort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.userService == nil {
			a.userService = a.userServiceWithDB(a.db)
		}
		return a.userService
	}

	return a.userServiceWithDB(db)
}

func (a *app) setDB() error {
	db, err := postgres.NewPsqlGormConnection(postgres.DBConnOptions{
		Host:   a.cfg.DB.Host,
		Port:   a.cfg.DB.Port,
		User:   a.cfg.DB.User,
		Pass:   a.cfg.DB.Pass,
		Name:   a.cfg.DB.Name,
		Schema: a.cfg.DB.Schema,
	})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	a.db = db
	return nil
}

func NewApp(cfg config.Config) (App, error) {
	a := &app{cfg: cfg}

	if err := a.setDB(); err != nil {
		return nil, err
	}

	err := a.userService.RunMigrations()

	if err != nil {
		return nil, err
	}

	a.setRedis()
	a.setEmailService()

	return a, nil
}

func (a *app) setRedis() {
	a.redis = cache.NewRedisProvider(
		fmt.Sprintf("%s:%d", a.cfg.Redis.Host, a.cfg.Redis.Port),
		"", "", 0,
	)
}

func (a *app) setEmailService() {
	a.mailService = email.NewEmailAdapter(a.cfg.SMTP)
}

func MustNewApp(cfg config.Config) App {
	a, err := NewApp(cfg)
	if err != nil {
		panic(err)
	}
	return a
}
