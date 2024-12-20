package port

import (
	"context"
	"navigation_service/internal/common/types"
	"navigation_service/internal/routing/domain"
)

type Service interface {
	CreateRouting(ctx context.Context, route *domain.Routing) error
	UpdateRouting(ctx context.Context, route *domain.Routing) error
	DeleteRouting(ctx context.Context, id uint) error
	GetRouting(ctx context.Context, id uint) (*domain.Routing, error)
	GetRoutingByUUID(ctx context.Context, uuid string) (*domain.Routing, error)
	FindRouting(ctx context.Context, filter domain.RouteFilter) ([]domain.Routing, error)
	ValidateRoutingForVehicleType(ctx context.Context, routeID uint, vehicleType types.VehicleType) error
}
