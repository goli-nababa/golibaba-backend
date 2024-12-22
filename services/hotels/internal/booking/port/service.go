package port

import (
	"context"
	"hotels-service/internal/booking/domain"
)

type Service interface {
	CancelBooking(ctx context.Context, UUID domain.BookingID) error
	CreateNewBooking(ctx context.Context, booking domain.Booking) (domain.BookingID, error)
	FindBooking(ctx context.Context, filters domain.BookingFilterItem) ([]domain.Booking, error)
	EditeBooking(ctx context.Context, UUID domain.BookingID, newBook domain.Booking) error
	DeleteBooking(ctx context.Context, UUID domain.BookingID) error
}
