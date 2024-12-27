package app

import (
	"context"
	"transportation/config"
	companyPort "transportation/internal/company/port"
	tripPort "transportation/internal/trip/port"
	vehicleRequestPollerPort "transportation/internal/vehicle_request_poller/port"

	"gorm.io/gorm"
)

type App interface {
	CompanyService(ctx context.Context) companyPort.Service
	TripService(ctx context.Context) tripPort.Service
	VehicleRequestPollerService(ctx context.Context) vehicleRequestPollerPort.Service

	DB() *gorm.DB
	Config() config.Config
}
