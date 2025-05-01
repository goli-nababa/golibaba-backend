package hotel

import (
	"context"
	"hotels-service/internal/hotel/domain"
	hotelDomain "hotels-service/internal/hotel/domain"
	hotelPort "hotels-service/internal/hotel/port"

	"github.com/google/uuid"
)

type service struct {
	repo hotelPort.Repo
}

func NewService(repo hotelPort.Repo) hotelPort.Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateHotel(ctx context.Context, hotel hotelDomain.Hotel) (hotelDomain.HotelID, error) {
	if err := hotel.Validate(); err != nil {
		return uuid.Nil, err
	}
	return hotel.ID, nil
}

func (s *service) UpdateHotel(ctx context.Context, UUID hotelDomain.HotelID, hotel hotelDomain.Hotel) error {
	var err error
	if err = hotelDomain.ValidateID(UUID); err != nil {
		return err
	}
	if err = hotel.Validate(); err != nil {
		return err
	}
	err = s.repo.Update(ctx, UUID, hotel)
	if err != nil {
		return err
	}
	return err
}
func (s *service) DeleteHotel(ctx context.Context, UUID hotelDomain.HotelID) error {
	var err error
	if err = hotelDomain.ValidateID(UUID); err != nil {
		return err
	}
	return s.repo.Delete(ctx, UUID)
}
func (s *service) GetHotelByID(ctx context.Context, UUID hotelDomain.HotelID) (*hotelDomain.Hotel, error) {
	err := domain.ValidateID(UUID)
	if err != nil {
		return nil, err
	}
	hotel, err := s.repo.GetByID(ctx, UUID)
	if err != nil {
		return &hotelDomain.Hotel{}, err
	}
	return hotel, err
}
func (s *service) ListHotels(ctx context.Context, pageIndex uint, pageSize uint) ([]hotelDomain.Hotel, error) {
	hotels, err := s.repo.Get(ctx, pageIndex, pageSize)
	if err != nil {
		return nil, err
	}
	return hotels, nil
}
func (s *service) FindHotels(ctx context.Context, filters hotelDomain.HotelFilterItem, pageIndex uint, pageSize uint) ([]hotelDomain.Hotel, error) {
	hotels, err := s.repo.Get(ctx, pageIndex, pageSize, filters)
	if err != nil {
		return nil, err
	}
	return hotels, nil
}
