package port

import (
	"context"
	"hotels-service/internal/booking/domain"
)

type Service interface {
	CancelBooking(ctx context.Context, UUID domain.BookingID) error
	CreateNewBooking(ctx context.Context, booking domain.Booking) (domain.BookingID, error)
	GetBookingByID(ctx context.Context, UUID domain.BookingID) (*domain.Booking, error)
	GetAllBooking(ctx context.Context, pageIndex, pageSize uint) ([]domain.Booking, error)
	FindBooking(ctx context.Context, filters domain.BookingFilterItem, pageIndex, pageSize uint) ([]domain.Booking, error)
	EditeBooking(ctx context.Context, UUID domain.BookingID, newBook domain.Booking) error
	DeleteBooking(ctx context.Context, UUID domain.BookingID) error
}
