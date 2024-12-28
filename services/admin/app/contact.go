package app

import (
	"admin/config"
	adminPort "admin/internal/port"
	"context"

	"gorm.io/gorm"
)

type App interface {
	AdminService(ctx context.Context) adminPort.Service
	DB() *gorm.DB
	Config() config.Config
}
