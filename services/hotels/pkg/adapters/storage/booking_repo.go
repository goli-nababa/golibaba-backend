package storage

import (
	"context"
	"hotels-service/internal/booking/domain"
	bookinPort "hotels-service/internal/booking/port"
	"hotels-service/pkg/adapters/storage/mapper"
	"hotels-service/pkg/adapters/storage/types"

	"github.com/google/uuid"
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

func (br *bookingRepo) Create(ctx context.Context, booking domain.Booking) (domain.BookingID, error) {
	bookingType := new(types.Booking)

	if err := mapper.ConvertTypes(booking, bookingType); err != nil {
		return uuid.Nil, err
	}

	result := br.db.Create(bookingType)
	return booking.ID, result.Error
}

func (br *bookingRepo) Delete(ctx context.Context, UUID domain.BookingID) error {
	bookingType := new(types.Booking)
	result := br.db.Delete(&bookingType, UUID)
	return result.Error
}
func (br *bookingRepo) Find(ctx context.Context, filter domain.BookingFilterItem) ([]domain.Booking, error) {
	// var bookingType []types.Booking
	// result := br.db.Where(&filter).Find(&bookingType)
	// return bookinType , result.Error
	panic("v any")
}
func (br *bookingRepo) GetByID(ctx context.Context, UUID domain.BookingID) (*domain.Booking, error) {
	panic("v any")
}
func (br *bookingRepo) Update(ctx context.Context, UUID domain.BookingID, newData domain.Booking) error {
	panic("v any")
}
