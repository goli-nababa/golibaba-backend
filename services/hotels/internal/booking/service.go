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

func (s *service) GetBookingByID(ctx context.Context, UUID domain.BookingID) (*domain.Booking, error) {

	return s.repo.GetByID(ctx, UUID)
}

func (s *service) GetAllBooking(ctx context.Context, pageIndex, pageSize uint) ([]domain.Booking, error) {
	return s.repo.Get(ctx, domain.BookingFilterItem{}, pageIndex, pageSize)
}

func (s *service) FindBooking(ctx context.Context, filters domain.BookingFilterItem, pageIndex, pageSize uint) ([]domain.Booking, error) {
	return s.repo.Get(ctx, filters, pageIndex, pageSize)
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
