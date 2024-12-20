package hotel

import (
	"context"
	"hotels-service/internal/hotel/domain"
	hotelPort "hotels-service/internal/hotel/port"
	userPort "hotels-service/internal/user/port"

	"github.com/google/uuid"
)

type service struct {
	userService userPort.Service
}

func NewService(userService userPort.Service) hotelPort.Service {
	return &service{
		userService: userService,
	}
}

func (s *service) Create(ctx context.Context, hotel domain.Hotel) (domain.HotelID, error) {
	return uuid.Nil, nil
}
func (s *service) Delete(ctx context.Context, UUID domain.HotelID) error {
	return nil
}
func (s *service) Get(ctx context.Context, filter domain.HotelFilterItem) ([]domain.Hotel, error) {
	return []domain.Hotel{}, nil
}
func (s *service) GetByID(ctx context.Context, UUID domain.HotelID) (*domain.Hotel, error) {
	return &domain.Hotel{}, nil
}
func (s *service) Update(ctx context.Context, UUID domain.HotelID, newData domain.Hotel) (domain.HotelID, error) {
	return uuid.Nil, nil
}
