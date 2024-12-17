package app

import (
	"context"
	"transportation/config"
	companyPort "transportation/internal/company/port"

	"gorm.io/gorm"
)

type App interface {
	CompanyService(ctx context.Context) companyPort.Service
	DB() *gorm.DB
	Config() config.Config
}
