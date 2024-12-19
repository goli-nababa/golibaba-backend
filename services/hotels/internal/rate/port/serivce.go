package port

import (
	"context"
	"hotels-service/internal/rate/domain"
)

type service interface {
	Create(ctx context.Context, hotel domain.Rate) (domain.RateID, error)
	GetByID(ctx context.Context, UUID domain.RateID) (*domain.Rate, error)
	Get(ctx context.Context, filter domain.RateFilterItem) ([]domain.Rate, error)
	Update(ctx context.Context, UUID domain.RateID, newData domain.Rate) (domain.RateID, error)
	Delete(ctx context.Context, UUID domain.RateID) error
}
