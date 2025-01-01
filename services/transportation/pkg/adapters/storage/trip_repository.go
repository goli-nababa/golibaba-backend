package storage

import (
	"context"
	"time"
	commonDomain "transportation/internal/common/domain"

	"transportation/internal/trip/domain"
	"transportation/internal/trip/port"
	"transportation/pkg/adapters/storage/mapper"
	"transportation/pkg/adapters/storage/types"

	"gorm.io/gorm"
)

type TripRepo struct {
	db *gorm.DB
}

func NewTripRepo(db *gorm.DB) port.Repo {
	repo := &TripRepo{db}
	return repo
}

func (r *TripRepo) Create(ctx context.Context, TripDomain domain.Trip) (*domain.Trip, error) {
	Trip := types.Trip{}
	if err := mapper.ConvertTypes(TripDomain, &Trip); err != nil {
		return nil, err
	}

	if err := CreateRecord(r.db, &Trip); err != nil {
		return nil, err
	}

	if err := mapper.ConvertTypes(Trip, &TripDomain); err != nil {
		return nil, err
	}
	return &TripDomain, nil
}

func (r *TripRepo) Update(ctx context.Context, id domain.TripId, TripDomain domain.Trip) (*domain.Trip, error) {
	Trip := types.Trip{}
	if err := mapper.ConvertTypes(TripDomain, &Trip); err != nil {
		return nil, err
	}
	if err := UpdateRecord(r.db, id, Trip); err != nil {
		return nil, err
	}

	domain, err := r.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return domain, nil
}

func (r *TripRepo) GetByID(ctx context.Context, id domain.TripId, preloads ...string) (*domain.Trip, error) {
	Trip, err := GetRecordByID[types.Trip](r.db, id, preloads)
	if err != nil {
		return nil, err
	}
	if Trip == nil {
		return nil, gorm.ErrRecordNotFound
	}

	TripDomain := domain.Trip{}
	if err := mapper.ConvertTypes(*Trip, &TripDomain); err != nil {
		return nil, err
	}

	return &TripDomain, nil
}

func (r *TripRepo) Delete(ctx context.Context, id domain.TripId) error {
	return DeleteRecordByID[types.Trip](r.db, id)
}

func (r *TripRepo) Get(ctx context.Context, request *commonDomain.RepositoryRequest) ([]domain.Trip, error) {
	companies, err := GetRecords[types.Trip](r.db, request)
	if err != nil {
		return []domain.Trip{}, err
	}

	TripDomains := []domain.Trip{}

	mapper.ConvertTypes(companies, &TripDomains)
	return TripDomains, nil
}

func (r *TripRepo) CreateVehicleRequest(ctx context.Context, vehicleRequestDomain domain.VehicleRequest) (*domain.VehicleRequest, error) {
	vehicleRequestTeam := types.VehicleRequest{}
	if err := mapper.ConvertTypes(vehicleRequestDomain, &vehicleRequestTeam); err != nil {
		return nil, err
	}

	if err := CreateRecord(r.db, &vehicleRequestTeam); err != nil {
		return nil, err
	}

	if err := mapper.ConvertTypes(vehicleRequestTeam, &vehicleRequestDomain); err != nil {
		return nil, err
	}
	return &vehicleRequestDomain, nil
}

func (r *TripRepo) UpdateVehicleRequest(ctx context.Context, id domain.VehicleRequestId, vehicleRequestDomain domain.VehicleRequest) (*domain.VehicleRequest, error) {
	VehicleRequest := types.VehicleRequest{}
	if err := mapper.ConvertTypes(vehicleRequestDomain, &VehicleRequest); err != nil {
		return nil, err
	}
	if err := UpdateRecord(r.db, id, VehicleRequest); err != nil {
		return nil, err
	}

	domain, err := r.GetVehicleRequestByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return domain, nil
}

func (r *TripRepo) GetVehicleRequestByID(ctx context.Context, id domain.VehicleRequestId) (*domain.VehicleRequest, error) {
	team, err := GetRecordByID[types.VehicleRequest](r.db, id, nil)
	if err != nil {
		return nil, err
	}
	if team == nil {
		return nil, gorm.ErrRecordNotFound
	}

	domain := domain.VehicleRequest{}
	if err := mapper.ConvertTypes(*team, &domain); err != nil {
		return nil, err
	}

	return &domain, nil
}

func (r *TripRepo) DeleteVehicleRequest(ctx context.Context, id domain.VehicleRequestId) error {
	return DeleteRecordByID[types.VehicleRequest](r.db, id)
}

func (r *TripRepo) GetVehicleRequests(ctx context.Context, request *commonDomain.RepositoryRequest) ([]domain.VehicleRequest, error) {
	companies, err := GetRecords[types.VehicleRequest](r.db, request)
	if err != nil {
		return []domain.VehicleRequest{}, err
	}

	vehicleRequestDomains := []domain.VehicleRequest{}

	mapper.ConvertTypes(companies, &vehicleRequestDomains)
	return vehicleRequestDomains, nil
}

func (r *TripRepo) GetShouldCheckVehicleRequests(ctx context.Context, limit int, offset int) (requests []domain.VehicleRequest, err error) {

	var vehicleRequests []types.VehicleRequest

	db := r.db.WithContext(ctx).
		Joins("inner join trips t on t.id = vehicle_requests.trip_id").
		Preload("Trip").
		Where("vehicle_requests.deleted_at IS NULL AND t.deleted_at IS NULL AND t.start_time > ?", time.Now()).
		Order("vehicle_requests.last_check_time ASC").
		Limit(limit).
		Offset(offset)

	if err := db.Find(&vehicleRequests).Error; err != nil {
		return nil, err
	}

	return requests, mapper.ConvertTypes(vehicleRequests, &requests)
}
func (r *TripRepo) UpdateVehicleRequestsLastCheckTime(ctx context.Context, ids []domain.VehicleRequestId) error {
	if len(ids) == 0 {
		return nil
	}

	err := r.db.WithContext(ctx).
		Model(&types.VehicleRequest{}).
		Where("id IN ?", ids).
		Update("last_check_time", time.Now()).
		Error

	return err
}
