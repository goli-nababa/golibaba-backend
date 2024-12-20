package port

import (
	"context"
	"navigation_service/internal/common/types"
	"navigation_service/internal/location/domain"
)

type Repo interface {
	Create(ctx context.Context, location *domain.Location) error
	Update(ctx context.Context, location *domain.Location) error
	Delete(ctx context.Context, id uint) error
	GetByID(ctx context.Context, id uint) (*domain.Location, error)
	GetByType(ctx context.Context, locationType types.LocationType) ([]domain.Location, error)
	List(ctx context.Context, active bool) ([]domain.Location, error)
}
