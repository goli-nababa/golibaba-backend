package port

import (
	"context"
	"time"
	"transportation/internal/trip/domain"
)

type Service interface {
	CreateTrip(ctx context.Context, t domain.CreateTripRequest) (domain.Trip, error)
	UpdateTrip(ctx context.Context, id domain.TripId, t domain.UpdateTripRequest) (domain.Trip, error)
	GetTrips(ctx context.Context, req domain.GetTripsRequest) ([]domain.Trip, error)
	SearchTrips(ctx context.Context, req domain.GetTripsRequest) ([]domain.Trip, error)
	GetTrip(ctx context.Context, id domain.TripId) (domain.Trip, error)

	SetVehicle(ctx context.Context, id domain.TripId, vehicleId domain.VehicleId) (domain.Trip, error)
	SetExpectedEndTime(ctx context.Context, id domain.TripId, expectedTime time.Time) (domain.Trip, error)
	ConfirmTechnicalTeam(ctx context.Context, id domain.TripId) (domain.Trip, error)
	EndTrip(ctx context.Context, id domain.TripId) (domain.Trip, error)
	ConfirmEndTrip(ctx context.Context, id domain.TripId) (domain.Trip, error)

	CreateVehicleRequest(ctx context.Context, req domain.CreateVehicleRequest) (domain.VehicleRequest, error)
	UpdateVehicleRequest(ctx context.Context, id domain.VehicleRequestId, req domain.CreateVehicleRequest) (domain.VehicleRequest, error)
	GetVehicleRequests(ctx context.Context, req domain.GetVehicleRequests) ([]domain.VehicleRequest, error)
	DeleteVehicleRequest(ctx context.Context, id domain.VehicleRequestId) error
}
