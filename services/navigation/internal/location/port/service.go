package port

import (
	"context"
	"navigation_service/internal/common/types"
	"navigation_service/internal/location/domain"
)

type Service interface {
	CreateLocation(ctx context.Context, location *domain.Location) error
	UpdateLocation(ctx context.Context, location *domain.Location) error
	DeleteLocation(ctx context.Context, id uint) error
	GetLocation(ctx context.Context, id uint) (*domain.Location, error)
	GetLocationsByType(ctx context.Context, locationType types.LocationType) ([]domain.Location, error)
	ListLocations(ctx context.Context, activeOnly bool) ([]domain.Location, error)
}
