package app

import (
	"context"
	"github.com/goli-nababa/golibaba-backend/modules/user_service_client"
	"hotels-service/config"
	bookingPort "hotels-service/internal/booking/port"
	hotelPort "hotels-service/internal/hotel/port"
	RateService "hotels-service/internal/rate/port"
	RoomService "hotels-service/internal/room/port"

	"gorm.io/gorm"
)

type App interface {
	Config() config.Config
	DB() *gorm.DB
	BookingService(ctx context.Context) bookingPort.Service
	HotelService(ctx context.Context) hotelPort.Service
	RateService(ctx context.Context) RateService.Service
	RoomService(ctx context.Context) RoomService.Service
	UserServiceClient() user_service_client.UserServiceClient
}
