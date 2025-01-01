package storage

import (
	"context"
	"hotels-service/internal/rate/domain"
	ratePort "hotels-service/internal/rate/port"
	"hotels-service/pkg/adapters/storage/types"
	"time"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type rateRepo struct {
	db *gorm.DB
}

func NewRateRepo(db *gorm.DB) ratePort.Repo {
	return &rateRepo{
		db: db,
	}
}

func (r *rateRepo) Create(ctx context.Context, newRecord domain.Rate) (domain.RateID, error) {
	rate := new(types.Rate)
	copier.Copy(rate, &newRecord)
	rate.CreateAt = time.Now()
	rate.UpdateAt = time.Now()
	result := r.db.WithContext(ctx).Create(rate)
	return newRecord.ID, result.Error
}
func (r *rateRepo) GetByID(ctx context.Context, UUID domain.RateID) (*domain.Rate, error) {
	rateDomain := new(domain.Rate)
	rate := new(types.Rate)
	result := r.db.WithContext(ctx).First(rate, UUID)
	copier.Copy(rateDomain, rate)
	return rateDomain, result.Error
}
func (r *rateRepo) Get(ctx context.Context, pageIndex, pageSize uint, filter ...domain.RateFilterItem) ([]domain.Rate, error) {
	var result *gorm.DB
	rates := new([]types.Rate)
	rateDomain := new([]domain.Rate)
	offset := (pageIndex - 1) * pageSize
	if len(filter) > 0 {
		result = r.db.WithContext(ctx).Limit(int(pageSize)).Offset(int(offset)).Where(&filter[0]).Find(rates)
	} else {
		result = r.db.WithContext(ctx).Limit(int(pageSize)).Offset(int(offset)).Where(&filter[0]).Find(rates)

	}
	copier.Copy(rateDomain, rates)
	return *rateDomain, result.Error
}
func (r *rateRepo) Update(ctx context.Context, UUID domain.RateID, newRecord domain.Rate) error {
	rate := new(types.Rate)
	copier.Copy(rate, &newRecord)
	rate.UpdateAt = time.Now()
	result := r.db.WithContext(ctx).Model(&rate).Where("id = ?", UUID.String()).Updates(rate)
	return result.Error
}
func (r *rateRepo) Delete(ctx context.Context, UUID domain.RateID) error {
	rate := new(types.Rate)
	rate.DeletedAt = time.Now()
	result := r.db.WithContext(ctx).Model(rate).Where("id=?", UUID.String()).Updates(rate)
	return result.Error
}
