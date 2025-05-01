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
	"sort"
)

type service struct {
	repo         port.Repo
	locationRepo locationPort.Repo
}

func (s *service) CalculateDistance(ctx context.Context, fromID, toID uint64, vehicleType types.VehicleType) (*domain.RouteDetails, error) {
	from, to, err := s.getLocationPair(ctx, uint(fromID), uint(toID))
	if err != nil {
		return nil, err
	}

	const earthRadius = 6371 // km

	lat1 := from.Latitude * math.Pi / 180
	lon1 := from.Longitude * math.Pi / 180
	lat2 := to.Latitude * math.Pi / 180
	lon2 := to.Longitude * math.Pi / 180

	dlat := lat2 - lat1
	dlon := lon2 - lon1

	a := math.Sin(dlat/2)*math.Sin(dlat/2) +
		math.Cos(lat1)*math.Cos(lat2)*
			math.Sin(dlon/2)*math.Sin(dlon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance := earthRadius * c

	// Estimated time based on vehicle type average speed
	var avgSpeed float64
	switch vehicleType {
	case types.VehicleTypeAirplane:
		avgSpeed = 800 // km/h
	case types.VehicleTypeTrain:
		avgSpeed = 100 // km/h
	case types.VehicleTypeBus:
		avgSpeed = 60 // km/h
	case types.VehicleTypeShip:
		avgSpeed = 40 // km/h
	default:
		avgSpeed = 60 // default speed
	}

	timeInHours := distance / avgSpeed
	timeInMinutes := timeInHours * 60

	// Basic cost calculation (can be enhanced based on business rules)
	baseCost := distance * 2 // Example: $2 per km

	return &domain.RouteDetails{
		Distance:      distance,
		EstimatedTime: int(timeInMinutes),
		EstimatedCost: baseCost,
	}, nil
}

func (s *service) FindOptimalRoutes(ctx context.Context, fromID, toID uint64, allowedVehicles []types.VehicleType) ([]domain.OptimalRoute, error) {
	filter := domain.RouteFilter{
		FromID:     uint(fromID),
		ToID:       uint(toID),
		ActiveOnly: true,
	}

	routes, err := s.repo.FindRoutes(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to find routes: %w", err)
	}

	optimalRoutes := make([]domain.OptimalRoute, 0)
	for _, route := range routes {
		for _, allowedType := range allowedVehicles {
			if route.SupportsVehicleType(allowedType) {
				details, err := s.CalculateDistance(ctx, uint64(route.FromID), uint64(route.ToID), allowedType)
				if err != nil {
					continue
				}

				score := calculateEfficiencyScore(route, details)

				optimalRoutes = append(optimalRoutes, domain.OptimalRoute{
					Route:           &route,
					EfficiencyScore: score,
					Criteria:        "distance,time,cost",
				})
				break
			}
		}
	}

	sortOptimalRoutesByScore(optimalRoutes)

	return optimalRoutes, nil
}

func (s *service) FindNearbyLocations(ctx context.Context, locationID uint64, radiusKm float64, locationType types.LocationType) ([]locationDomain.Location, map[uint64]float64, error) {
	center, err := s.locationRepo.GetByID(ctx, uint(locationID))
	if err != nil {
		return nil, nil, err
	}

	locations, err := s.locationRepo.GetByType(ctx, locationType)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get locations: %w", err)
	}

	nearbyLocations := make([]locationDomain.Location, 0)
	distances := make(map[uint64]float64)

	for _, loc := range locations {
		distance := calculateHaversineDistance(
			center.Latitude, center.Longitude,
			loc.Latitude, loc.Longitude,
		)

		if distance <= radiusKm {
			nearbyLocations = append(nearbyLocations, loc)
			distances[uint64(loc.ID)] = distance
		}
	}

	return nearbyLocations, distances, nil
}

func (s *service) GetRouteStatistics(ctx context.Context, filter domain.StatisticsFilter) (*domain.RouteStatistics, error) {
	return s.repo.GetStatistics(ctx, filter)
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

	if !route.SupportsVehicleType(vehicleType) {
		return fmt.Errorf("vehicle type %s is not supported for route %d", vehicleType, routeID)
	}

	return nil
}

func (s *service) calculateDistance(route *domain.Routing) (float64, error) {
	if route.From.ID == 0 || route.To.ID == 0 {
		return 0, fmt.Errorf("route locations must be loaded before calculating distance")
	}

	distance := calculateHaversineDistance(
		route.From.Latitude, route.From.Longitude,
		route.To.Latitude, route.To.Longitude,
	)

	return math.Round(distance*100) / 100, nil // Round to 2 decimal places
}

func calculateHaversineDistance(lat1, lon1, lat2, lon2 float64) float64 {
	const earthRadius = 6371

	lat1 = lat1 * math.Pi / 180
	lon1 = lon1 * math.Pi / 180
	lat2 = lat2 * math.Pi / 180
	lon2 = lon2 * math.Pi / 180

	// Haversine formula
	dlat := lat2 - lat1
	dlon := lon2 - lon1

	a := math.Sin(dlat/2)*math.Sin(dlat/2) +
		math.Cos(lat1)*math.Cos(lat2)*
			math.Sin(dlon/2)*math.Sin(dlon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadius * c
}

func calculateEfficiencyScore(route domain.Routing, details *domain.RouteDetails) float64 {
	const (
		distanceWeight = 0.4
		timeWeight     = 0.3
		costWeight     = 0.3
	)

	normalizedDistance := 1000 / details.Distance
	normalizedTime := 300 / float64(details.EstimatedTime)
	normalizedCost := 1000 / details.EstimatedCost

	score := (normalizedDistance * distanceWeight) +
		(normalizedTime * timeWeight) +
		(normalizedCost * costWeight)

	return math.Round(score*100) / 100
}

func sortOptimalRoutesByScore(routes []domain.OptimalRoute) {
	sort.Slice(routes, func(i, j int) bool {
		return routes[i].EfficiencyScore > routes[j].EfficiencyScore
	})
}
