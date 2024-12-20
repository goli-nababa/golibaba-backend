package booking

import (
	"context"
	"hotels-service/internal/booking/domain"
	bookingPort "hotels-service/internal/booking/port"
	roomPort "hotels-service/internal/room/port"
	userPort "hotels-service/internal/user/port"

	"github.com/google/uuid"
)

type service struct {
	userService userPort.Service
	roomService roomPort.Service
}

func NewService(userService userPort.Service, roomService roomPort.Service) bookingPort.Service {
	return &service{
		userService: userService,
		roomService: roomService,
	}
}

func (s *service) CancelBooking(ctx context.Context, UUID domain.BookingID) error {
	err := domain.ValidateID(UUID)
	if err != nil {
		return err
	}
	return nil
}
func (s *service) Create(ctx context.Context, hotel domain.Booking) (domain.BookingID, error) {
	return uuid.Nil, nil
}
func (s *service) Delete(ctx context.Context, UUID domain.BookingID) error {
	return nil
}
func (s *service) Get(ctx context.Context, filter domain.BookingFilterItem) ([]domain.Booking, error) {
	return []domain.Booking{}, nil
}
func (s *service) GetByID(ctx context.Context, UUID domain.BookingID) (*domain.Booking, error) {
	return nil, nil
}
func (s *service) Update(ctx context.Context, UUID domain.BookingID, newData domain.Booking) (domain.BookingID, error) {
	return uuid.Nil, nil
}
