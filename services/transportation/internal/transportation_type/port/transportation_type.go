package port

import (
	"context"
	commonDomain "transportation/internal/common/domain"

	"transportation/internal/transportation_type/domain"
)

type Repo interface {
	Create(ctx context.Context, TransportationTypeDomain domain.TransportationType) (*domain.TransportationType, error)
	Update(ctx context.Context, id domain.TransportationTypeId, TransportationTypeDomain domain.TransportationType) (*domain.TransportationType, error)
	GetByID(ctx context.Context, id domain.TransportationTypeId) (*domain.TransportationType, error)
	Delete(ctx context.Context, id domain.TransportationTypeId) error
	Get(ctx context.Context, request *commonDomain.RepositoryRequest) ([]domain.TransportationType, error)
}
