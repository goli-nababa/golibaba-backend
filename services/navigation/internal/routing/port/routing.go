package port

import (
	"context"
	"navigation_service/internal/routing/domain"
)

type Repo interface {
	Create(ctx context.Context, route *domain.Routing) error
	Update(ctx context.Context, route *domain.Routing) error
	Delete(ctx context.Context, id uint) error
	GetByID(ctx context.Context, id uint) (*domain.Routing, error)
	GetByUUID(ctx context.Context, uuid string) (*domain.Routing, error)
	GetByCode(ctx context.Context, code string) (*domain.Routing, error)
	FindRoutes(ctx context.Context, filter domain.RouteFilter) ([]domain.Routing, error)
	GetStatistics(ctx context.Context, filter domain.StatisticsFilter) (*domain.RouteStatistics, error)
}
