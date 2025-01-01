package services

import (
	"context"
	"navigation_service/internal/common/types"
	"navigation_service/internal/location/domain"
	locationPort "navigation_service/internal/location/port"
)

type LocationService struct {
	svc locationPort.Service
}

func NewLocationService(svc locationPort.Service) *LocationService {
	return &LocationService{
		svc: svc,
	}
}

func (s *LocationService) CreateLocation(ctx context.Context, location *domain.Location) error {
	return s.svc.CreateLocation(ctx, location)
}

func (s *LocationService) UpdateLocation(ctx context.Context, location *domain.Location) error {
	return s.svc.UpdateLocation(ctx, location)
}

func (s *LocationService) DeleteLocation(ctx context.Context, id uint) error {
	return s.svc.DeleteLocation(ctx, id)
}

func (s *LocationService) GetLocation(ctx context.Context, id uint) (*domain.Location, error) {
	return s.svc.GetLocation(ctx, id)
}

func (s *LocationService) GetLocationsByType(ctx context.Context, locationType types.LocationType) ([]domain.Location, error) {
	return s.svc.GetLocationsByType(ctx, locationType)
}

func (s *LocationService) ListLocations(ctx context.Context, activeOnly bool) ([]domain.Location, error) {
	return s.svc.ListLocations(ctx, activeOnly)
}
