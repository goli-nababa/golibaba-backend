package port

import (
	"context"
	"hotels-service/internal/rate/domain"
)

type Service interface {
	CreateNewRate(ctx context.Context, booking domain.Rate) (domain.RateID, error)
	GetRateByID(ctx context.Context, UUID domain.RateID) (*domain.Rate, error)
	GetAllRate(ctx context.Context, pageIndex, pageSize uint) ([]domain.Rate, error)
	FindRate(ctx context.Context, filters domain.RateFilterItem, pageIndex, pageSize uint) ([]domain.Rate, error)
	EditeRate(ctx context.Context, UUID domain.RateID, newBook domain.Rate) error
	DeleteRate(ctx context.Context, UUID domain.RateID) error
}
