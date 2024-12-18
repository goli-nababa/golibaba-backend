package port

import (
	"context"
	"hotels-service/internal/booking/domain"
)

type service interface {
	Create(ctx context.Context, hotel domain.Booking) (domain.BookingID, error)
	GetByID(ctx context.Context, UUID domain.BookingID) (*domain.Booking, error)
	Get(ctx context.Context, filter domain.BookingFilterItem) ([]domain.Booking, error)
	Update(ctx context.Context, UUID domain.BookingID, newData domain.Booking) (domain.BookingID, error)
	Delete(ctx context.Context, UUID domain.BookingID) error
}
