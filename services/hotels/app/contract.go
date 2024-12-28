package app

import (
	"hotels-service/config"
	bookingPort "hotels-service/internal/booking/port"
)

type App interface {
	Config() config.Config
	BookingService() bookingPort.Service
}
