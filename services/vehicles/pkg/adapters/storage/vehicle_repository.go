package storage

import (
	"context"
	"vehicles/internal/common/domain"
	"vehicles/pkg/adapters/storage/types"

	"gorm.io/gorm"
)

type VehicleRepo struct {
	db *gorm.DB
}

func NewVehicleRepo(db *gorm.DB) *VehicleRepo {
	repo := &VehicleRepo{db}
	return repo
}

func (r *VehicleRepo) Get(ctx context.Context, request *domain.RepositoryRequest) ([]types.Vehicle, error) {
	companies, err := GetRecords[types.Vehicle](r.db, request)
	if err != nil {
		return []types.Vehicle{}, err
	}

	return companies, nil
}
