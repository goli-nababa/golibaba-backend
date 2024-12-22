package booking

import (
	"context"
	"hotels-service/internal/booking/domain"
	bookingPort "hotels-service/internal/booking/port"

	"github.com/google/uuid"
)

type service struct {
	repo bookingPort.Repo
}

func NewService(repo bookingPort.Repo) bookingPort.Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CancelBooking(ctx context.Context, UUID domain.BookingID) error {
	oldBooking, err := s.repo.GetByID(ctx, UUID)
	if err != nil {
		return err
	}
	newBooking := oldBooking
	newBooking.Status = domain.StatusTypeCancelled
	s.repo.Update(ctx, UUID, *newBooking)
	return nil
}

func (s *service) CreateNewBooking(ctx context.Context, booking domain.Booking) (domain.BookingID, error) {
	err := booking.Validate()
	if err != nil {
		return uuid.Nil, err
	}
	bookingID, err := s.repo.Create(ctx, booking)
	if err != nil {
		return uuid.Nil, err
	}
	return bookingID, err
}
func (s *service) FindBooking(ctx context.Context, filters domain.BookingFilterItem) ([]domain.Booking, error) {
	return s.repo.Find(ctx, filters)
}
func (s *service) EditeBooking(ctx context.Context, bookingID domain.BookingID, newBook domain.Booking) error {
	err := newBook.Validate()
	if err != nil {
		return err
	}
	return s.repo.Update(ctx, bookingID, newBook)
}
func (s *service) DeleteBooking(ctx context.Context, UUID domain.BookingID) error {
	return s.repo.Delete(ctx, UUID)
}
