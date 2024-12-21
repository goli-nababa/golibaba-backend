package app

import (
	"hotels-service/config"
	"hotels-service/internal/booking"
	bookingPort "hotels-service/internal/booking/port"

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
	a.bookingService = booking.NewService()
	return a, nil
}

func (a *app) setDB() error {
	//TODO
	return nil
}

func (a *app) Config() config.Config {
	return a.cfg
}

func (a *app) BookingService() bookingPort.Service {
	return a.bookingService
}
