package storage

import (
	"context"
	"hotels-service/internal/hotel/domain"
	hotelPort "hotels-service/internal/hotel/port"
	"hotels-service/pkg/adapters/storage/types"
	"time"

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

func (hr *hotelRepo) Create(ctx context.Context, newRecord domain.Hotel) (domain.HotelID, error) {
	hotel := new(types.Hotel)
	copier.Copy(hotel, &newRecord)
	hotel.CreatedAt = time.Now()
	hotel.UpdateAt = time.Now()
	result := hr.db.WithContext(ctx).Create(hotel)
	return newRecord.ID, result.Error
}

func (hr *hotelRepo) Delete(ctx context.Context, UUID domain.HotelID) error {
	hotel := new(types.Hotel)
	hotel.DeletedAt = time.Now()
	result := hr.db.WithContext(ctx).Model(hotel).Where("id=?", UUID.String()).Updates(hotel)
	return result.Error
}
func (hr *hotelRepo) Get(ctx context.Context, pageIndex uint, pageSize uint, filters ...domain.HotelFilterItem) ([]domain.Hotel, error) {
	var result *gorm.DB
	domainHotels := new([]domain.Hotel)
	hotel := new([]types.Hotel)
	offset := (pageIndex - 1) * pageSize
	if len(filters) > 0 {
		result = hr.db.WithContext(ctx).Limit(int(pageSize)).Offset(int(offset)).Where(&filters[0]).Find(hotel)
	} else {
		result = hr.db.WithContext(ctx).Limit(int(pageSize)).Offset(int(offset)).Find(hotel)
	}
	copier.Copy(domainHotels, hotel)
	return *domainHotels, result.Error
}
func (hr *hotelRepo) GetByID(ctx context.Context, UUID domain.HotelID) (*domain.Hotel, error) {
	hotel := new(types.Hotel)
	domainHotel := new(domain.Hotel)
	result := hr.db.WithContext(ctx).First(&hotel, UUID.String())
	copier.Copy(domainHotel, hotel)
	return domainHotel, result.Error
}
func (hr *hotelRepo) Update(ctx context.Context, UUID domain.HotelID, newRecord domain.Hotel) error {
	hotel := new(types.Hotel)
	copier.Copy(hotel, &newRecord)
	hotel.UpdateAt = time.Now()
	result := hr.db.WithContext(ctx).Model(&hotel).Where("id = ?", UUID.String()).Updates(hotel)
	return result.Error
}
