package storage

import (
	"context"
	"hotels-service/internal/booking/domain"
	bookingPort "hotels-service/internal/booking/port"
	"hotels-service/pkg/adapters/storage/types"

	"github.com/jinzhu/copier"

	"gorm.io/gorm"
)

type bookingRepo struct {
	db *gorm.DB
}

func NewBookinRepo(db *gorm.DB) bookingPort.Repo {
	return &bookingRepo{
		db: db,
	}
}

func (br *bookingRepo) Create(ctx context.Context, booking domain.Booking) (domain.BookingID, error) {
	bookingType := new(types.Booking)
	copier.Copy(bookingType, &booking)
	result := br.db.Create(bookingType)
	return booking.ID, result.Error
}

func (br *bookingRepo) Delete(ctx context.Context, UUID domain.BookingID) error {
	booking := new(types.Booking)
	result := br.db.Delete(&booking, UUID.String())
	return result.Error
}
func (br *bookingRepo) Get(ctx context.Context, pageIndex uint, pageSize uint, filters ...domain.BookingFilterItem) ([]domain.Booking, error) {
	var result *gorm.DB
	domainBookings := new([]domain.Booking)
	booking := new([]types.Booking)
	offset := (pageIndex - 1) * pageSize
	if len(filters) > 0 {
		result = br.db.Limit(int(pageSize)).Offset(int(offset)).Where(&filters).Find(booking)
	} else {
		result = br.db.Limit(int(pageSize)).Offset(int(offset)).Find(booking)
	}
	copier.Copy(domainBookings, booking)
	return *domainBookings, result.Error
}
func (br *bookingRepo) GetByID(ctx context.Context, UUID domain.BookingID) (*domain.Booking, error) {
	booking := new(types.Booking)
	domainBooking := new(domain.Booking)
	result := br.db.First(&booking, UUID.String())
	copier.Copy(domainBooking, booking)
	return domainBooking, result.Error
}
func (br *bookingRepo) Update(ctx context.Context, UUID domain.BookingID, newData domain.Booking) error {
	booking := new(types.Booking)
	copier.Copy(booking, &newData)
	result := br.db.Model(&booking).Where("id = ?", UUID.String()).Updates(booking)
	return result.Error
}
