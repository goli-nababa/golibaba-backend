package app

import (
	"context"
	"fmt"
	"github.com/goli-nababa/golibaba-backend/modules/user_service_client"
	"hotels-service/config"
	"hotels-service/internal/booking"
	bookingPort "hotels-service/internal/booking/port"
	"hotels-service/internal/hotel"
	hotelPort "hotels-service/internal/hotel/port"
	"hotels-service/internal/rate"
	ratePort "hotels-service/internal/rate/port"
	"hotels-service/internal/room"
	roomPort "hotels-service/internal/room/port"
	"hotels-service/pkg/adapters/migrations"
	"hotels-service/pkg/adapters/storage"
	appCtx "hotels-service/pkg/context"
	"hotels-service/pkg/postgress"

	"gorm.io/gorm"
)

type app struct {
	db                *gorm.DB
	cfg               config.Config
	bookingService    bookingPort.Service
	hotelService      hotelPort.Service
	rateService       ratePort.Service
	roomService       roomPort.Service
	userServiceClient user_service_client.UserServiceClient
}

func (a *app) UserServiceClient() user_service_client.UserServiceClient {
	return a.userServiceClient
}

func NewApp(cfg config.Config) (App, error) {
	userClient, err := user_service_client.NewUserServiceClient(
		cfg.UserService.URL,
		cfg.UserService.Version,
		uint64(cfg.UserService.Port),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize user service client: %w", err)
	}

	a := &app{
		cfg:               cfg,
		userServiceClient: userClient,
	}

	if err := a.setDB(); err != nil {
		return nil, err
	}

	return a, nil
}

func (a *app) setDB() error {
	db, err := postgress.NewPsqlGormConnection(postgress.DBConnOptions{
		Host:   a.cfg.Database.Host,
		Port:   a.cfg.Database.Port,
		User:   a.cfg.Database.User,
		Pass:   a.cfg.Database.Password,
		Name:   a.cfg.Database.Database,
		Schema: a.cfg.Database.Schema,
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
