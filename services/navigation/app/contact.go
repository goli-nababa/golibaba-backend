package app

import (
	"context"
	"gorm.io/gorm"
	"navigation_service/config"
	LocationPort "navigation_service/internal/location/port"
	routePort "navigation_service/internal/routing/port"
)

type App interface {
	LocationService(ctx context.Context) LocationPort.Service
	RoteService(ctx context.Context) routePort.Service
	DB() *gorm.DB
	Config() config.Config
}
