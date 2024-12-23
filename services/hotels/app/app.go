package app

import (
	"hotels-service/config"
	"hotels-service/internal/booking"
	bookingPort "hotels-service/internal/booking/port"
	"hotels-service/pkg/adapters/storage"
	"hotels-service/pkg/postgress"

	"gorm.io/gorm"
)

type app struct {
	db             *gorm.DB
	cfg            config.Config
	bookingService bookingPort.Service
}

func AppNew(cfg config.Config) (App, error) {
	a := &app{
		cfg: cfg,
	}
	return a, nil
}

func (a *app) setDB() error {
	db, err := postgress.NewConnection(postgress.DBOpt{
		Host:     a.cfg.Database.Host,
		Port:     a.cfg.Database.Port,
		User:     a.cfg.Database.User,
		Password: a.cfg.Database.Password,
		DBName:   a.cfg.Database.DBName,
		SSLMode:  a.cfg.Database.SSLMode,
		Schema:   a.cfg.Database.Schema,
	})

	if err != nil {
		return nil
	}

	a.db = db

	return nil
}

func (a *app) DB() *gorm.DB {
	return a.db
}

func (a *app) bookingSerivce() bookingPort.Service {

	return nil
}

func (a *app) bookingServiceWithDB(db *gorm.DB) bookingPort.Service {
	return booking.NewService(storage.NewBookinRepo(db))
}

func (a *app) Config() config.Config {
	return a.cfg
}

func (a *app) BookingService() bookingPort.Service {
	return a.bookingService
}
