package trip

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"
	commonDomain "transportation/internal/common/domain"
	"transportation/internal/trip/domain"

	"transportation/internal/trip/port"
)

type TripService struct {
	repo port.Repo
}

func NewTripService(repo port.Repo) port.Service {
	return &TripService{repo: repo}
}

func (s *TripService) CreateTrip(ctx context.Context, req domain.CreateTripRequest) (domain.Trip, error) {
	trip := domain.Trip{
		Title:                req.Title,
		OriginStationId:      req.OriginStationId,
		DestinationStationId: req.DestinationStationId,
		CompanyId:            req.CompanyId,
		StartTime:            req.StartTime,
		PassengersCountLimit: req.PassengersCountLimit,
		NormalCost:           req.NormalCost,
		AgencyCost:           req.AgencyCost,
		AgencyReleaseTime:    req.AgencyReleaseTime,
		ReleaseTime:          req.ReleaseTime,
		TechTeamId:           req.TechTeamId,
		CreatedAt:            time.Now(),
		UpdatedAt:            time.Now(),
	}
	createdTrip, err := s.repo.Create(ctx, trip)
	if err != nil {
		return domain.Trip{}, err
	}
	return *createdTrip, nil
}

func (s *TripService) UpdateTrip(ctx context.Context, id domain.TripId, req domain.UpdateTripRequest) (domain.Trip, error) {
	trip, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return domain.Trip{}, err
	}

	// Update fields
	trip.Title = req.Title
	trip.OriginStationId = req.OriginStationId
	trip.DestinationStationId = req.DestinationStationId
	trip.CompanyId = req.CompanyId
	trip.StartTime = req.StartTime
	trip.PassengersCountLimit = req.PassengersCountLimit
	trip.NormalCost = req.NormalCost
	trip.AgencyCost = req.AgencyCost
	trip.AgencyReleaseTime = req.AgencyReleaseTime
	trip.ReleaseTime = req.ReleaseTime
	trip.TechTeamId = req.TechTeamId
	trip.UpdatedAt = time.Now()

	updatedTrip, err := s.repo.Update(ctx, id, *trip)
	if err != nil {
		return domain.Trip{}, err
	}
	return *updatedTrip, nil
}

func (s *TripService) GetTrip(ctx context.Context, id domain.TripId) (domain.Trip, error) {

	trip, err := s.repo.GetByID(ctx, id, "Company")

	if err != nil {
		return domain.Trip{}, err
	}

	return *trip, nil

}

func (s *TripService) GetTrips(ctx context.Context, req domain.GetTripsRequest) ([]domain.Trip, error) {

	role := "admin"

	repoRequest := commonDomain.RepositoryRequest{Sorts: []*commonDomain.RepositorySort{&commonDomain.RepositorySort{Field: "created_at", SortType: "desc"}},
		Filters: []*commonDomain.RepositoryFilter{}}

	if req.CompanyId > 0 {
		repoRequest.Filters = append(repoRequest.Filters, &commonDomain.RepositoryFilter{Field: "company_id", Operator: "=", Value: strconv.Itoa(int(req.CompanyId))})
	}
	if req.TechTeamId > 0 {
		repoRequest.Filters = append(repoRequest.Filters, &commonDomain.RepositoryFilter{Field: "technical_team_id", Operator: "=", Value: strconv.Itoa(int(req.TechTeamId))})
	}
	if req.DestinationStationId > 0 {
		repoRequest.Filters = append(repoRequest.Filters, &commonDomain.RepositoryFilter{Field: "destination_station_id", Operator: "=", Value: strconv.Itoa(int(req.DestinationStationId))})
	}
	if req.OriginStationId > 0 {
		repoRequest.Filters = append(repoRequest.Filters, &commonDomain.RepositoryFilter{Field: "origin_station_id", Operator: "=", Value: strconv.Itoa(int(req.OriginStationId))})
	}
	if req.VehicleId > 0 {
		repoRequest.Filters = append(repoRequest.Filters, &commonDomain.RepositoryFilter{Field: "vehicle_id", Operator: "=", Value: strconv.Itoa(int(req.VehicleId))})
	}
	if req.FromStartTime != nil {
		repoRequest.Filters = append(repoRequest.Filters, &commonDomain.RepositoryFilter{Field: "from_start_time", Operator: ">=", Value: req.FromStartTime.Format(time.RFC3339)})
	}
	if req.ToStartTime != nil {
		repoRequest.Filters = append(repoRequest.Filters, &commonDomain.RepositoryFilter{Field: "to_start_time", Operator: "<=", Value: req.ToStartTime.Format(time.RFC3339)})
	}

	if role == "technical" {
		repoRequest.Filters = append(repoRequest.Filters, &commonDomain.RepositoryFilter{Field: "ent_time", Operator: "IS", Value: "NULL"})
		repoRequest.Filters = append(repoRequest.Filters, &commonDomain.RepositoryFilter{Field: "vehicle_id", Operator: "IS", Value: "NOT NULL"})
	} else if role != "admin" {
		return nil, errors.New("role is not defined")
	}

	trips, err := s.repo.Get(ctx, &repoRequest)
	fmt.Println(err)

	if err != nil {
		return nil, err
	}
	return trips, nil
}

func (s *TripService) SearchTrips(ctx context.Context, req domain.GetTripsRequest) ([]domain.Trip, error) {

	role := "user"

	repoRequest := commonDomain.RepositoryRequest{Sorts: []*commonDomain.RepositorySort{&commonDomain.RepositorySort{Field: "created_at", SortType: "desc"}},
		Filters: []*commonDomain.RepositoryFilter{}}

	if req.CompanyId > 0 {
		repoRequest.Filters = append(repoRequest.Filters, &commonDomain.RepositoryFilter{Field: "company_id", Operator: "=", Value: strconv.Itoa(int(req.CompanyId))})
	}
	if req.TechTeamId > 0 {
		repoRequest.Filters = append(repoRequest.Filters, &commonDomain.RepositoryFilter{Field: "technical_team_id", Operator: "=", Value: strconv.Itoa(int(req.TechTeamId))})
	}
	if req.DestinationStationId > 0 {
		repoRequest.Filters = append(repoRequest.Filters, &commonDomain.RepositoryFilter{Field: "destination_station_id", Operator: "=", Value: strconv.Itoa(int(req.DestinationStationId))})
	}
	if req.OriginStationId > 0 {
		repoRequest.Filters = append(repoRequest.Filters, &commonDomain.RepositoryFilter{Field: "origin_station_id", Operator: "=", Value: strconv.Itoa(int(req.OriginStationId))})
	}
	if req.VehicleId > 0 {
		repoRequest.Filters = append(repoRequest.Filters, &commonDomain.RepositoryFilter{Field: "vehicle_id", Operator: "=", Value: strconv.Itoa(int(req.VehicleId))})
	}
	if req.FromStartTime != nil {
		repoRequest.Filters = append(repoRequest.Filters, &commonDomain.RepositoryFilter{Field: "from_start_time", Operator: ">=", Value: req.FromStartTime.Format(time.RFC3339)})
	}
	if req.ToStartTime != nil {
		repoRequest.Filters = append(repoRequest.Filters, &commonDomain.RepositoryFilter{Field: "to_start_time", Operator: "<=", Value: req.ToStartTime.Format(time.RFC3339)})
	}

	repoRequest.Filters = append(repoRequest.Filters, &commonDomain.RepositoryFilter{Field: "end_time", Operator: "IS", Value: "NULL"})
	repoRequest.Filters = append(repoRequest.Filters, &commonDomain.RepositoryFilter{Field: "tech_team_confirmation_time", Operator: "IS", Value: "NOT NULL"})
	repoRequest.Filters = append(repoRequest.Filters, &commonDomain.RepositoryFilter{Field: "vehicle_id", Operator: "!=", Value: "0"})

	if role == "agency" {
		repoRequest.Filters = append(repoRequest.Filters, &commonDomain.RepositoryFilter{Field: "agency_release_time", Operator: "<", Value: time.Now().Format(time.RFC3339)})

	} else {
		repoRequest.Filters = append(repoRequest.Filters, &commonDomain.RepositoryFilter{Field: "release_time", Operator: "<", Value: time.Now().Format(time.RFC3339)})
	}

	oneDayAfterNow := time.Now().Add(24 * time.Hour)
	repoRequest.Filters = append(repoRequest.Filters, &commonDomain.RepositoryFilter{Field: "start_time", Operator: ">", Value: oneDayAfterNow.Format(time.RFC3339)})

	trips, err := s.repo.Get(ctx, &repoRequest)

	if err != nil {
		return nil, err
	}
	return trips, nil
}

func (s *TripService) SetVehicle(ctx context.Context, id domain.TripId, vehicleId domain.VehicleId) (domain.Trip, error) {
	trip, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return domain.Trip{}, err
	}
	trip.VehicleId = vehicleId
	trip.UpdatedAt = time.Now()

	trip, err = s.repo.Update(ctx, id, *trip)
	if err != nil {
		return domain.Trip{}, err
	}
	return *trip, err
}

func (s *TripService) SetExpectedEndTime(ctx context.Context, id domain.TripId, expectedTime time.Time) (domain.Trip, error) {
	trip, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return domain.Trip{}, err
	}
	trip.ExpectedEndTime = &expectedTime
	trip.UpdatedAt = time.Now()

	trip, err = s.repo.Update(ctx, id, *trip)
	if err != nil {
		return domain.Trip{}, err
	}
	return *trip, err
}

func (s *TripService) ConfirmTechnicalTeam(ctx context.Context, id domain.TripId) (domain.Trip, error) {
	print("asd")
	trip, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return domain.Trip{}, err
	}

	now := time.Now()
	trip.TechTeamConfirmationTime = &now
	trip.UpdatedAt = now

	updatedTrip, err := s.repo.Update(ctx, id, *trip)
	if err != nil {
		return domain.Trip{}, err
	}
	return *updatedTrip, nil
}

func (s *TripService) EndTrip(ctx context.Context, id domain.TripId) (domain.Trip, error) {
	trip, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return domain.Trip{}, err
	}

	now := time.Now()

	if trip.VehicleId == 0 {
		return domain.Trip{}, errors.New("no vehicle set on trip")
	}

	if trip.ExpectedEndTime == nil {
		return domain.Trip{}, errors.New("no expected end time set on trip")

	}

	if now.Before(*trip.ExpectedEndTime) {
		return domain.Trip{}, errors.New("ent time must be after expected end time")
	}

	trip.EndTime = &now
	trip.UpdatedAt = now

	updatedTrip, err := s.repo.Update(ctx, id, *trip)
	if err != nil {
		return domain.Trip{}, err
	}
	return *updatedTrip, nil
}

func (s *TripService) ConfirmEndTrip(ctx context.Context, id domain.TripId) (domain.Trip, error) {
	trip, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return domain.Trip{}, err
	}

	now := time.Now()
	trip.EndConfirmationTime = &now
	trip.UpdatedAt = now

	updatedTrip, err := s.repo.Update(ctx, id, *trip)
	if err != nil {
		return domain.Trip{}, err
	}
	return *updatedTrip, nil
}

func (s *TripService) CreateVehicleRequest(ctx context.Context, req domain.CreateVehicleRequest) (domain.VehicleRequest, error) {
	vehicleRequest := domain.VehicleRequest{
		TripId:              req.TripId,
		VehicleTypeId:       req.VehicleTypeId,
		VehicleCost:         req.VehicleCost,
		VehicleCreationDate: &req.VehicleCreationDate.Time,
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
	}

	createdVehicleRequest, err := s.repo.CreateVehicleRequest(ctx, vehicleRequest)
	if err != nil {
		return domain.VehicleRequest{}, err
	}
	return *createdVehicleRequest, nil
}

func (s *TripService) UpdateVehicleRequest(ctx context.Context, id domain.VehicleRequestId, req domain.CreateVehicleRequest) (domain.VehicleRequest, error) {
	vehicleRequest, err := s.repo.GetVehicleRequestByID(ctx, id)
	if err != nil {
		return domain.VehicleRequest{}, err
	}

	vehicleRequest.VehicleTypeId = req.VehicleTypeId
	vehicleRequest.VehicleCost = req.VehicleCost
	vehicleRequest.VehicleCreationDate = &req.VehicleCreationDate.Time
	vehicleRequest.UpdatedAt = time.Now()

	updatedVehicleRequest, err := s.repo.UpdateVehicleRequest(ctx, id, *vehicleRequest)
	if err != nil {
		return domain.VehicleRequest{}, err
	}
	return *updatedVehicleRequest, nil
}

func (s *TripService) GetVehicleRequests(ctx context.Context, req domain.GetVehicleRequests) ([]domain.VehicleRequest, error) {
	filters := []*commonDomain.RepositoryFilter{}
	if req.TripId > 0 {
		filters = append(filters, &commonDomain.RepositoryFilter{Field: "trip_id", Operator: "=", Value: strconv.Itoa(int(req.TripId))})

	}
	repoRequest := commonDomain.RepositoryRequest{Sorts: []*commonDomain.RepositorySort{&commonDomain.RepositorySort{Field: "created_at", SortType: "desc"}}, Filters: filters}
	vehicleRequests, err := s.repo.GetVehicleRequests(ctx, &repoRequest)
	if err != nil {
		return nil, err
	}
	return vehicleRequests, nil
}

func (s *TripService) DeleteVehicleRequest(ctx context.Context, id domain.VehicleRequestId) error {
	return s.repo.DeleteVehicleRequest(ctx, id)
}
