package storage

import (
	"context"
	"hotels-service/internal/hotel/domain"
	hotelPort "hotels-service/internal/hotel/port"
	"hotels-service/pkg/adapters/storage/types"

	"github.com/jinzhu/copier"

	"gorm.io/gorm"
)

type hotelRepo struct {
	db *gorm.DB
}

func NewHotelRepo(db *gorm.DB) hotelPort.Repo {
	return &hotelRepo{
		db: db,
	}
}

func (hr *hotelRepo) Create(ctx context.Context, booking domain.Hotel) (domain.HotelID, error) {
	bookingType := new(types.Hotel)
	copier.Copy(bookingType, &booking)
	result := hr.db.Create(bookingType)
	return booking.ID, result.Error
}

func (hr *hotelRepo) Delete(ctx context.Context, UUID domain.HotelID) error {
	booking := new(types.Hotel)
	result := hr.db.Delete(&booking, UUID.String())
	return result.Error
}
func (hr *hotelRepo) Get(ctx context.Context, pageIndex uint, pageSize uint, filters ...domain.HotelFilterItem) ([]domain.Hotel, error) {
	var result *gorm.DB
	domainHotels := new([]domain.Hotel)
	booking := new([]types.Hotel)
	offset := (pageIndex - 1) * pageSize
	if len(filters) > 0 {
		result = hr.db.Limit(int(pageSize)).Offset(int(offset)).Where(&filters[0]).Find(booking)
	} else {
		result = hr.db.Limit(int(pageSize)).Offset(int(offset)).Find(booking)
	}
	copier.Copy(domainHotels, booking)
	return *domainHotels, result.Error
}
func (hr *hotelRepo) GetByID(ctx context.Context, UUID domain.HotelID) (*domain.Hotel, error) {
	booking := new(types.Hotel)
	domainHotel := new(domain.Hotel)
	result := hr.db.First(&booking, UUID.String())
	copier.Copy(domainHotel, booking)
	return domainHotel, result.Error
}
func (hr *hotelRepo) Update(ctx context.Context, UUID domain.HotelID, newData domain.Hotel) error {
	booking := new(types.Hotel)
	copier.Copy(booking, &newData)
	result := hr.db.Model(&booking).Where("id = ?", UUID.String()).Updates(booking)
	return result.Error
}
