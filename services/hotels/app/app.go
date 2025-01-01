package app

import (
	"context"
	"hotels-service/config"
	"hotels-service/internal/booking"
	bookingPort "hotels-service/internal/booking/port"
	"hotels-service/internal/hotel"
	hotelPort "hotels-service/internal/hotel/port"
	"hotels-service/internal/rate"
	ratePort "hotels-service/internal/rate/port"
	"hotels-service/internal/room"
	roomPort "hotels-service/internal/room/port"
	"hotels-service/pkg/adapters/storage"
	appCtx "hotels-service/pkg/context"
	"hotels-service/pkg/postgress"

	"gorm.io/gorm"
)

type app struct {
	db             *gorm.DB
	cfg            config.Config
	bookingService bookingPort.Service
	hotelService   hotelPort.Service
	rateService    ratePort.Service
	roomService    roomPort.Service
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

func (a *app) Config() config.Config {
	return a.cfg
}

func (a *app) bookingServiceWithDB(db *gorm.DB) bookingPort.Service {
	return booking.NewService(storage.NewBookingRepo(db))
}

func (a *app) BookingService(ctx context.Context) bookingPort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.bookingService == nil {
			a.bookingService = a.bookingServiceWithDB(a.db)
		}
		return a.bookingService
	}
	return nil
}

func (a *app) hotelServiceWithDB(db *gorm.DB) hotelPort.Service {
	return hotel.NewService(storage.NewHotelRepo(db))
}

func (a *app) HotelService(ctx context.Context) hotelPort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.hotelService == nil {
			a.hotelService = a.hotelServiceWithDB(a.db)
		}
		return a.hotelService
	}
	return nil
}

func (a *app) rateServiceWithDB(db *gorm.DB) ratePort.Service {
	return rate.NewService(storage.NewRateRepo(db))
}

func (a *app) RateService(ctx context.Context) ratePort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.rateService == nil {
			a.rateService = a.rateServiceWithDB(a.db)
		}
		return a.rateService
	}
	return nil
}

func (a *app) roomServiceWithDB(db *gorm.DB) roomPort.Service {
	return room.NewService(storage.NewRoomRepo(db))
}

func (a *app) RoomService(ctx context.Context) roomPort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.roomService == nil {
			a.roomService = a.roomServiceWithDB(a.db)
		}
		return a.roomService
	}
	return nil
}
