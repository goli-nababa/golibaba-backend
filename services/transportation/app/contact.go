package app

import (
	"context"
	"transportation/config"
	companyPort "transportation/internal/company/port"
	tripPort "transportation/internal/trip/port"

	"gorm.io/gorm"
)

type App interface {
	CompanyService(ctx context.Context) companyPort.Service
	TripService(ctx context.Context) tripPort.Service

	DB() *gorm.DB
	Config() config.Config
}
