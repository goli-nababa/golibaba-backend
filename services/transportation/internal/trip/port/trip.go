package port

import (
	"context"
	commonDomain "transportation/internal/common/domain"
	"transportation/internal/trip/domain"
)

type Repo interface {
	Create(ctx context.Context, TripDomain domain.Trip) (*domain.Trip, error)
	Update(ctx context.Context, id domain.TripId, TripDomain domain.Trip) (*domain.Trip, error)
	GetByID(ctx context.Context, id domain.TripId) (*domain.Trip, error)
	Delete(ctx context.Context, id domain.TripId) error
	Get(ctx context.Context, request *commonDomain.RepositoryRequest) ([]domain.Trip, error)

	CreateVehicleRequest(ctx context.Context, vehicleRequestDomain domain.VehicleRequest) (*domain.VehicleRequest, error)
	UpdateVehicleRequest(ctx context.Context, id domain.VehicleRequestId, vehicleRequestDomain domain.VehicleRequest) (*domain.VehicleRequest, error)
	GetVehicleRequestByID(ctx context.Context, id domain.VehicleRequestId) (*domain.VehicleRequest, error)
	DeleteVehicleRequest(ctx context.Context, id domain.VehicleRequestId) error
	GetVehicleRequests(ctx context.Context, request *commonDomain.RepositoryRequest) ([]domain.VehicleRequest, error)
}
