package app

import (
	"context"
	"user_service/config"
	userPort "user_service/internal/user/port"

	"gorm.io/gorm"
)

type App interface {
	Config() config.Config
	DB() *gorm.DB
	UserService(ctx context.Context) userPort.Service
}
