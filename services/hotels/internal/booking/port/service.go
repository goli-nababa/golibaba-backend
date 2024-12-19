package port

import (
	"context"
	"hotels-service/internal/booking/domain"
)

type Service interface {
	CancelBooking(ctx context.Context, UUID domain.BookingID) error
	Create(ctx context.Context, hotel domain.Booking) (domain.BookingID, error)
	Delete(ctx context.Context, UUID domain.BookingID) error
	Get(ctx context.Context, filter domain.BookingFilterItem) ([]domain.Booking, error)
	GetByID(ctx context.Context, UUID domain.BookingID) (*domain.Booking, error)
	Update(ctx context.Context, UUID domain.BookingID, newData domain.Booking) (domain.BookingID, error)
}
