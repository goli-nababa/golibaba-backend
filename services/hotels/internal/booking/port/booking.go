package port

import (
	"context"
	"hotels-service/internal/booking/domain"
)

type Repo interface {
	Create(ctx context.Context, booking domain.Booking) (domain.BookingID, error)
	Delete(ctx context.Context, UUID domain.BookingID) error
	Get(ctx context.Context, pageIndex uint, pageSize uint, filters ...domain.BookingFilterItem) ([]domain.Booking, error)
	GetByID(ctx context.Context, UUID domain.BookingID) (*domain.Booking, error)
	Update(ctx context.Context, UUID domain.BookingID, newData domain.Booking) error
}
