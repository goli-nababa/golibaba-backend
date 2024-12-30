package port

import (
	"context"
	hotelDomain "hotels-service/internal/hotel/domain"
)

type Service interface {
	CreateHotel(ctx context.Context, hotel hotelDomain.Hotel) (hotelDomain.HotelID, error)
	UpdateHotel(ctx context.Context, UUID hotelDomain.HotelID, hotel hotelDomain.Hotel) error
	DeleteHotel(ctx context.Context, UUID hotelDomain.HotelID) error
	GetHotelByID(ctx context.Context, UUID hotelDomain.HotelID) (*hotelDomain.Hotel, error)
	ListHotels(ctx context.Context, pageIndex uint, pageSize uint) ([]hotelDomain.Hotel, error)
	FindHotels(ctx context.Context, filters hotelDomain.HotelFilterItem, pageIndex uint, pageSize uint) ([]hotelDomain.Hotel, error)
}
