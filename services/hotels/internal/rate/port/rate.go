package port

import (
	"context"
	"hotels-service/internal/rate/domain"
)

type Repo interface {
	Create(ctx context.Context, rate domain.Rate) (domain.RateID, error)
	GetByID(ctx context.Context, UUID domain.RateID) (*domain.Rate, error)
	Get(ctx context.Context, pageIndex, pageSize uint, filter ...domain.RateFilterItem) ([]domain.Rate, error)
	Update(ctx context.Context, UUID domain.RateID, newData domain.Rate) error
	Delete(ctx context.Context, UUID domain.RateID) error
}
