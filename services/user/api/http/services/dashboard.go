package services

import (
	"context"
	"user_service/api/http/handlers/helpers"
	"user_service/app"
	"user_service/config"
	"user_service/internal/domain"
	"user_service/internal/user/port"

	"github.com/goli-nababa/golibaba-backend/common"
)

type DashboardService struct {
	UserService port.Service
}

func NewDashboardService(svc port.Service) *DashboardService {
	return &DashboardService{
		UserService: svc,
	}
}

func DashboardServiceGetter(appContainer app.App, cfg config.ServerConfig) helpers.ServiceGetter[*DashboardService] {
	return func(ctx context.Context) *DashboardService {
		return NewDashboardService(appContainer.UserService(ctx))
	}
}

func (ds *DashboardService) SaveLog(c context.Context, log *common.Log) error {
	return ds.UserService.SaveLog(c, log)
}

func (ds *DashboardService) GetHistory(c context.Context, userId uint, page int, pageSize int) ([]common.Log, error) {
	return ds.UserService.GetHistory(c, userId, page, pageSize)
}

func (ds *DashboardService) GetNotifications(c context.Context, userId uint) ([]domain.Notification, error) {
	return ds.UserService.GetNotifications(c, userId)
}

func (ds *DashboardService) CreateNotification(c context.Context, notif *domain.Notification) error {
	return ds.UserService.CreateNotification(c, notif)
}
