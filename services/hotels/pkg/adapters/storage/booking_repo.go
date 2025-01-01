package storage

import (
	"context"
	"hotels-service/internal/booking/domain"
	bookingPort "hotels-service/internal/booking/port"
	"hotels-service/pkg/adapters/storage/types"
	"time"

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

func (br *bookingRepo) Create(ctx context.Context, newRecord domain.Booking) (domain.BookingID, error) {
	book := new(types.Booking)
	copier.Copy(book, &newRecord)
	book.CreateAt = time.Now()
	book.UpdatedAt = time.Now()
	result := br.db.Create(book)
	return newRecord.ID, result.Error
}

func (br *bookingRepo) Delete(ctx context.Context, UUID domain.BookingID) error {
	book := new(types.Booking)
	book.DeletedAt = time.Now()
	result := br.db.Model(book).Where("id = ?", UUID.String()).Updates(book)
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
	book := new(types.Booking)
	domainBooking := new(domain.Booking)
	result := br.db.First(book, UUID.String())
	copier.Copy(domainBooking, book)
	return domainBooking, result.Error
}
func (br *bookingRepo) Update(ctx context.Context, UUID domain.BookingID, newRecord domain.Booking) error {
	book := new(types.Booking)
	copier.Copy(book, &newRecord)
	book.UpdatedAt = time.Now()
	result := br.db.Model(book).Where("id = ?", UUID.String()).Updates(book)
	return result.Error
}
