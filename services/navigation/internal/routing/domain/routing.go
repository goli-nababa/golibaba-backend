package domain

import (
	"errors"
	"github.com/google/uuid"
	"navigation_service/internal/common/types"
	"navigation_service/internal/location/domain"
	"time"
)

type RoutingID uint

type Routing struct {
	ID           RoutingID
	UUID         string
	Code         string
	FromID       uint
	ToID         uint
	From         domain.Location
	To           domain.Location
	Distance     float64
	VehicleTypes types.VehicleTypes
	Active       bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewRoute(code string, fromID, toID uint, distance float64, vehicleTypes types.VehicleTypes) (*Routing, error) {
	if err := vehicleTypes.Validate(); err != nil {
		return nil, err
	}

	route := &Routing{
		UUID:         uuid.New().String(),
		Code:         code,
		FromID:       fromID,
		ToID:         toID,
		Distance:     distance,
		VehicleTypes: vehicleTypes,
		Active:       true,
	}

	if err := route.Validate(); err != nil {
		return nil, err
	}

	return route, nil
}

func (r *Routing) Validate() error {
	if r.Code == "" {
		return errors.New("code is required")
	}

	if r.FromID == 0 {
		return errors.New("from location is required")
	}

	if r.ToID == 0 {
		return errors.New("to location is required")
	}

	if r.FromID == r.ToID {
		return errors.New("from and to locations must be different")
	}

	if r.Distance <= -1 {
		return errors.New("distance must be positive")
	}

	return r.VehicleTypes.Validate()
}

func (r *Routing) Update(code string, fromID, toID uint, distance float64, vehicleTypes types.VehicleTypes, active bool) error {
	if err := vehicleTypes.Validate(); err != nil {
		return err
	}

	r.Code = code
	r.FromID = fromID
	r.ToID = toID
	r.Distance = distance
	r.VehicleTypes = vehicleTypes
	r.Active = active

	return r.Validate()
}

func (r *Routing) SupportsVehicleType(vehicleType types.VehicleType) bool {
	return r.VehicleTypes.Contains(vehicleType)
}

type RouteDetails struct {
	Distance      float64
	EstimatedTime int // in minutes
	EstimatedCost float64
}

type OptimalRoute struct {
	Route           *Routing
	EfficiencyScore float64
	Criteria        string
}

type RouteStatistics struct {
	TotalRoutes         int
	TotalDistance       float64
	RoutesByVehicleType map[string]int
	PopularRoutes       []PopularRoute
}

type PopularRoute struct {
	Route         *Routing
	UsageCount    int
	AverageRating float64
}
