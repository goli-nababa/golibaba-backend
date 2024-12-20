package routing

import (
	"context"
	"fmt"
	"math"
	"navigation_service/internal/common/types"
	locationDomain "navigation_service/internal/location/domain"
	locationPort "navigation_service/internal/location/port"
	"navigation_service/internal/routing/domain"
	"navigation_service/internal/routing/port"
)

type service struct {
	repo         port.Repo
	locationRepo locationPort.Repo
}

func NewService(repo port.Repo, locationRepo locationPort.Repo) port.Service {
	return &service{
		repo:         repo,
		locationRepo: locationRepo,
	}
}
func (s *service) CreateRouting(ctx context.Context, route *domain.Routing) error {
	from, to, err := s.getLocationPair(ctx, route.FromID, route.ToID)
	if err != nil {
		return fmt.Errorf("failed to get locations: %w", err)
	}

	route.From = *from
	route.To = *to

	distance, err := s.calculateDistance(route)
	if err != nil {
		return fmt.Errorf("failed to calculate distance: %w", err)
	}
	route.Distance = distance

	if err := route.Validate(); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	if err := s.repo.Create(ctx, route); err != nil {
		return fmt.Errorf("failed to create route: %w", err)
	}
	return nil
}

func (s *service) getLocationPair(ctx context.Context, fromID, toID uint) (*locationDomain.Location, *locationDomain.Location, error) {
	from, err := s.locationRepo.GetByID(ctx, fromID)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get source location: %w", err)
	}
	if from == nil {
		return nil, nil, fmt.Errorf("source location not found: %d", fromID)
	}

	to, err := s.locationRepo.GetByID(ctx, toID)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get destination location: %w", err)
	}
	if to == nil {
		return nil, nil, fmt.Errorf("destination location not found: %d", toID)
	}

	return from, to, nil
}

func (s *service) UpdateRouting(ctx context.Context, route *domain.Routing) error {
	if err := route.Validate(); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	existing, err := s.repo.GetByID(ctx, uint(route.ID))
	if err != nil {
		return fmt.Errorf("failed to get route: %w", err)
	}
	if existing == nil {
		return fmt.Errorf("route not found: %d", route.ID)
	}

	if existing.FromID != route.FromID || existing.ToID != route.ToID {
		from, to, err := s.getLocationPair(ctx, route.FromID, route.ToID)
		if err != nil {
			return fmt.Errorf("failed to get locations: %w", err)
		}

		route.From = *from
		route.To = *to

		distance, err := s.calculateDistance(route)
		if err != nil {
			return fmt.Errorf("failed to calculate distance: %w", err)
		}
		route.Distance = distance
	} else {
		route.Distance = existing.Distance
	}

	if err := s.repo.Update(ctx, route); err != nil {
		return fmt.Errorf("failed to update route: %w", err)
	}
	return nil
}

func (s *service) DeleteRouting(ctx context.Context, id uint) error {
	existing, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to get route: %w", err)
	}
	if existing == nil {
		return fmt.Errorf("route not found: %d", id)
	}

	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete route: %w", err)
	}
	return nil
}

func (s *service) GetRouting(ctx context.Context, id uint) (*domain.Routing, error) {
	route, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get route: %w", err)
	}
	return route, nil
}

func (s *service) GetRoutingByUUID(ctx context.Context, uuid string) (*domain.Routing, error) {
	route, err := s.repo.GetByUUID(ctx, uuid)
	if err != nil {
		return nil, fmt.Errorf("failed to get route by UUID: %w", err)
	}
	return route, nil
}

func (s *service) FindRouting(ctx context.Context, filter domain.RouteFilter) ([]domain.Routing, error) {
	routes, err := s.repo.FindRoutes(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to find routes: %w", err)
	}
	return routes, nil
}

func (s *service) ValidateRoutingForVehicleType(ctx context.Context, routeID uint, vehicleType types.VehicleType) error {
	route, err := s.repo.GetByID(ctx, routeID)
	if err != nil {
		return fmt.Errorf("failed to get route: %w", err)
	}
	if route == nil {
		return fmt.Errorf("route not found: %d", routeID)
	}

	supported := false
	for _, vt := range route.VehicleTypes {
		if vt == vehicleType {
			supported = true
			break
		}
	}

	if !supported {
		return fmt.Errorf("vehicle type %s is not supported for route %d", vehicleType, routeID)
	}

	return nil
}

func (s *service) calculateDistance(route *domain.Routing) (float64, error) {
	if route.From.ID == 0 || route.To.ID == 0 {
		return 0, fmt.Errorf("route locations must be loaded before calculating distance")
	}

	const (
		earthRadius = 6371 // Earth's radius in kilometers
		toRadians   = math.Pi / 180
	)

	// Convert latitude and longitude to radians
	lat1 := route.From.Latitude * toRadians
	lon1 := route.From.Longitude * toRadians
	lat2 := route.To.Latitude * toRadians
	lon2 := route.To.Longitude * toRadians

	// Haversine formula
	dlat := lat2 - lat1
	dlon := lon2 - lon1

	a := math.Pow(math.Sin(dlat/2), 2) +
		math.Cos(lat1)*math.Cos(lat2)*
			math.Pow(math.Sin(dlon/2), 2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance := earthRadius * c

	// Round to 2 decimal places
	return math.Round(distance*100) / 100, nil
}
