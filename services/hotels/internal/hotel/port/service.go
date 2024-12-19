package port

import (
	"context"
	"hotels-service/internal/hotel/domain"
)

type Service interface {
	Create(ctx context.Context, hotel domain.Hotel) (domain.HotelID, error)
	Delete(ctx context.Context, UUID domain.HotelID) error
	Get(ctx context.Context, filter domain.HotelFilterItem) ([]domain.Hotel, error)
	GetByID(ctx context.Context, UUID domain.HotelID) (*domain.Hotel, error)
	Update(ctx context.Context, UUID domain.HotelID, newData domain.Hotel) (domain.HotelID, error)
}
