package storage

import (
	"context"
	"hotels-service/internal/rate/domain"
	"hotels-service/internal/rate/port"
	"hotels-service/pkg/adapters/storage/types"
	"time"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type rateRepo struct {
	db *gorm.DB
}

func NewRateRepo(db *gorm.DB) port.Repo {
	return &rateRepo{
		db: db,
	}
}

func (r *rateRepo) Create(ctx context.Context, rate domain.Rate) (domain.RateID, error) {
	rateType := new(types.Rate)
	copier.Copy(rateType, &rate)
	rateType.CreateAt = time.Now()
	rateType.UpdateAt = time.Now()
	result := r.db.Create(rateType)
	return rate.ID, result.Error
}
func (r *rateRepo) GetByID(ctx context.Context, UUID domain.RateID) (*domain.Rate, error) {
	rateDomain := new(domain.Rate)
	rate := new(types.Rate)
	result := r.db.First(rate, UUID)
	copier.Copy(rateDomain, rate)
	return rateDomain, result.Error
}
func (r *rateRepo) Get(ctx context.Context, pageIndex, pageSize uint, filter ...domain.RateFilterItem) ([]domain.Rate, error) {
	var result *gorm.DB
	rates := new([]types.Rate)
	rateDomain := new([]domain.Rate)
	offset := (pageIndex - 1) * pageSize
	if len(filter) > 0 {
		result = r.db.Limit(int(pageSize)).Offset(int(offset)).Where(&filter[0]).Find(rates)
	} else {
		result = r.db.Limit(int(pageSize)).Offset(int(offset)).Where(&filter[0]).Find(rates)

	}
	copier.Copy(rateDomain, rates)
	return *rateDomain, result.Error
}
func (r *rateRepo) Update(ctx context.Context, UUID domain.RateID, newData domain.Rate) error {
	rate := new(types.Rate)
	copier.Copy(rate, &newData)
	rate.UpdateAt = time.Now()
	result := r.db.Model(&rate).Where("id = ?", UUID.String()).Updates(rate)
	return result.Error
}
func (r *rateRepo) Delete(ctx context.Context, UUID domain.RateID) error {
	rate := new(types.Rate)
	rate.DeletedAt = time.Now()
	result := r.db.Model(rate).Where("id=?", UUID.String()).Updates(rate)
	return result.Error
}
