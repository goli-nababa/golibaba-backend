package storage

import (
	"context"
	bookingDomain "hotels-service/internal/booking/domain"
	bookinPort "hotels-service/internal/booking/port"

	"gorm.io/gorm"
)

type bookingRepo struct {
	db *gorm.DB
}

func NewBookinRepo(db *gorm.DB) bookinPort.Repo {
	return &bookingRepo{
		db: db,
	}
}

func (br *bookingRepo) CancelBooking(ctx context.Context, UUID bookingDomain.BookingID) error {
	panic("")
}
func (br *bookingRepo) Create(ctx context.Context, hotel bookingDomain.Booking) (bookingDomain.BookingID, error) {
	panic("")
}
func (br *bookingRepo) Delete(ctx context.Context, UUID bookingDomain.BookingID) error {
	panic("")
}
func (br *bookingRepo) Get(ctx context.Context, filter bookingDomain.BookingFilterItem) ([]bookingDomain.Booking, error) {
	panic("")
}
func (br *bookingRepo) GetByID(ctx context.Context, UUID bookingDomain.BookingID) (*bookingDomain.Booking, error) {
	panic("")
}
func (br *bookingRepo) Update(ctx context.Context, UUID bookingDomain.BookingID, newData bookingDomain.Booking) (bookingDomain.BookingID, error) {
	panic("")
}
