package location

import (
	"context"
	"fmt"
	"navigation_service/internal/common/types"
	"navigation_service/internal/location/domain"
	"navigation_service/internal/location/port"
)

type service struct {
	locationRepo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{
		locationRepo: repo,
	}
}

func (s *service) CreateLocation(ctx context.Context, location *domain.Location) error {
	if err := location.Validate(); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	if err := s.locationRepo.Create(ctx, location); err != nil {
		return fmt.Errorf("failed to create location: %w", err)
	}
	return nil
}

func (s *service) UpdateLocation(ctx context.Context, location *domain.Location) error {
	if err := location.Validate(); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	existing, err := s.locationRepo.GetByID(ctx, uint(location.ID))
	if err != nil {
		return fmt.Errorf("failed to get location: %w", err)
	}
	if existing == nil {
		return fmt.Errorf("location not found: %d", location.ID)
	}

	if err := s.locationRepo.Update(ctx, location); err != nil {
		return fmt.Errorf("failed to update location: %w", err)
	}
	return nil
}

func (s *service) DeleteLocation(ctx context.Context, id uint) error {
	existing, err := s.locationRepo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to get location: %w", err)
	}
	if existing == nil {
		return fmt.Errorf("location not found: %d", id)
	}

	if err := s.locationRepo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete location: %w", err)
	}
	return nil
}

func (s *service) GetLocation(ctx context.Context, id uint) (*domain.Location, error) {
	location, err := s.locationRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get location: %w", err)
	}
	return location, nil
}

func (s *service) GetLocationsByType(ctx context.Context, locationType types.LocationType) ([]domain.Location, error) {
	if err := locationType.Validate(); err != nil {
		return nil, fmt.Errorf("invalid location type: %w", err)
	}

	locations, err := s.locationRepo.GetByType(ctx, locationType)
	if err != nil {
		return nil, fmt.Errorf("failed to get locations by type: %w", err)
	}
	return locations, nil
}

func (s *service) ListLocations(ctx context.Context, activeOnly bool) ([]domain.Location, error) {
	locations, err := s.locationRepo.List(ctx, activeOnly)
	if err != nil {
		return nil, fmt.Errorf("failed to list locations: %w", err)
	}
	return locations, nil
}
