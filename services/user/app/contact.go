package app

import (
	"context"
	"github.com/goli-nababa/golibaba-backend/modules/cache"
	"user_service/config"
	userPort "user_service/internal/user/port"
	"user_service/pkg/email"

	"gorm.io/gorm"
)

type App interface {
	Config() config.Config
	DB() *gorm.DB

	Cache() cache.Provider
	MailService() email.Adapter
	UserService(ctx context.Context) userPort.Service
}
