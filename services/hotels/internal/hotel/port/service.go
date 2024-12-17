package port

import (
	"context"
	"hotels-service/internal/hotel/domain"
)

type Service interface {
	Create(ctx context.Context, hotel domain.Hotel) (domain.HotelId, error)
	GetById(ctx context.Context, UUID domain.HotelId) (*domain.Hotel, error)
	Get(ctx context.Context, filter domain.HotelFilterItem) ([]domain.Hotel, error)
	Update(ctx context.Context, UUID domain.HotelId, newData domain.Hotel) (domain.HotelId, error)
	Delete(ctx context.Context, UUID domain.HotelId) error
}
