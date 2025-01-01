package services

import (
	"context"
	"navigation_service/internal/common/types"
	"navigation_service/internal/routing/domain"
	routingPort "navigation_service/internal/routing/port"
)

type RoutingService struct {
	svc routingPort.Service
}

func NewRoutingService(svc routingPort.Service) *RoutingService {
	return &RoutingService{
		svc: svc,
	}
}

func (s *RoutingService) CreateRoute(ctx context.Context, route *domain.Routing) error {
	return s.svc.CreateRouting(ctx, route)
}

func (s *RoutingService) UpdateRoute(ctx context.Context, route *domain.Routing) error {
	return s.svc.UpdateRouting(ctx, route)
}

func (s *RoutingService) DeleteRoute(ctx context.Context, id uint) error {
	return s.svc.DeleteRouting(ctx, id)
}

func (s *RoutingService) GetRoute(ctx context.Context, id uint) (*domain.Routing, error) {
	return s.svc.GetRouting(ctx, id)
}

func (s *RoutingService) GetRouteByUUID(ctx context.Context, uuid string) (*domain.Routing, error) {
	return s.svc.GetRoutingByUUID(ctx, uuid)
}

func (s *RoutingService) FindRoutes(ctx context.Context, filter domain.RouteFilter) ([]domain.Routing, error) {
	return s.svc.FindRouting(ctx, filter)
}

func (s *RoutingService) ValidateRouteForVehicleType(ctx context.Context, routeID uint, vehicleType types.VehicleType) error {
	return s.svc.ValidateRoutingForVehicleType(ctx, routeID, vehicleType)
}
