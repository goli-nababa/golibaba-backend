package port

import (
	"context"
	"hotels-service/internal/hotel/domain"
)

type Repo interface {
	Create(ctx context.Context, hotel domain.Hotel) (domain.HotelID, error)
	Delete(ctx context.Context, UUID domain.HotelID) error
	Get(ctx context.Context, pageIndex uint, pageSize uint, filter ...domain.HotelFilterItem) ([]domain.Hotel, error)
	GetByID(ctx context.Context, UUID domain.HotelID) (*domain.Hotel, error)
	Update(ctx context.Context, UUID domain.HotelID, newData domain.Hotel) error
}
