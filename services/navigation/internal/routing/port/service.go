package port

import (
	"context"
	"navigation_service/internal/common/types"
	locationDomain "navigation_service/internal/location/domain"
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
	CalculateDistance(ctx context.Context, fromID, toID uint64, vehicleType types.VehicleType) (*domain.RouteDetails, error)
	FindOptimalRoutes(ctx context.Context, fromID, toID uint64, allowedVehicles []types.VehicleType) ([]domain.OptimalRoute, error)
	FindNearbyLocations(ctx context.Context, locationID uint64, radiusKm float64, locationType types.LocationType) ([]locationDomain.Location, map[uint64]float64, error)
	GetRouteStatistics(ctx context.Context, filter domain.StatisticsFilter) (*domain.RouteStatistics, error)
}
