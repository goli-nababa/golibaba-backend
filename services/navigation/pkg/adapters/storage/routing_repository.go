package storage

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"navigation_service/internal/common/types"
	"navigation_service/internal/routing/domain"
	"navigation_service/internal/routing/port"
)

type routingRepository struct {
	db *gorm.DB
}

func NewRoutingRepository(db *gorm.DB) port.Repo {
	return &routingRepository{db: db}
}

func (r *routingRepository) Create(ctx context.Context, route *domain.Routing) error {
	if err := route.Validate(); err != nil {
		return fmt.Errorf("invalid route: %w", err)
	}

	result := r.db.WithContext(ctx).Create(route)
	if result.Error != nil {
		return fmt.Errorf("failed to create route: %w", result.Error)
	}

	return nil
}

func (r *routingRepository) Update(ctx context.Context, route *domain.Routing) error {
	if err := route.Validate(); err != nil {
		return fmt.Errorf("invalid route: %w", err)
	}

	result := r.db.WithContext(ctx).Model(&domain.Routing{}).
		Where("id = ?", route.ID).
		Updates(map[string]interface{}{
			"code":          route.Code,
			"from_id":       route.FromID,
			"to_id":         route.ToID,
			"distance":      route.Distance,
			"vehicle_types": route.VehicleTypes,
			"active":        route.Active,
		})

	if result.Error != nil {
		return fmt.Errorf("failed to update route: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("route not found: %d", route.ID)
	}

	return nil
}

func (r *routingRepository) Delete(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).Delete(&domain.Routing{}, id)
	if result.Error != nil {
		return fmt.Errorf("failed to delete route: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("route not found: %d", id)
	}

	return nil
}

func (r *routingRepository) GetByID(ctx context.Context, id uint) (*domain.Routing, error) {
	var route domain.Routing

	result := r.db.WithContext(ctx).
		Preload("From").
		Preload("To").
		First(&route, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get route: %w", result.Error)
	}

	return &route, nil
}

func (r *routingRepository) GetByUUID(ctx context.Context, uuid string) (*domain.Routing, error) {
	var route domain.Routing

	result := r.db.WithContext(ctx).
		Preload("From").
		Preload("To").
		Where("uuid = ?", uuid).
		First(&route)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get route by UUID: %w", result.Error)
	}

	return &route, nil
}

func (r *routingRepository) GetByCode(ctx context.Context, code string) (*domain.Routing, error) {
	var route domain.Routing

	result := r.db.WithContext(ctx).
		Preload("From").
		Preload("To").
		Where("code = ?", code).
		First(&route)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get route by code: %w", result.Error)
	}

	return &route, nil
}

func (r *routingRepository) FindRoutes(ctx context.Context, filter domain.RouteFilter) ([]domain.Routing, error) {
	if err := filter.Validate(); err != nil {
		return nil, fmt.Errorf("invalid filter: %w", err)
	}

	query := r.db.WithContext(ctx).
		Preload("From").
		Preload("To").
		Offset(filter.GetOffset()).
		Limit(filter.GetLimit())

	if filter.FromID != 0 {
		query = query.Where("from_id = ?", filter.FromID)
	}

	if filter.ToID != 0 {
		query = query.Where("to_id = ?", filter.ToID)
	}

	if filter.VehicleType != "" {
		query = query.Where("vehicle_types @> ?", types.VehicleTypes{filter.VehicleType})
	}

	if filter.ActiveOnly {
		query = query.Where("active = true")
	}

	var routes []domain.Routing
	result := query.Find(&routes)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to find routes: %w", result.Error)
	}

	return routes, nil
}

func (r *routingRepository) GetStatistics(ctx context.Context, filter domain.StatisticsFilter) (*domain.RouteStatistics, error) {
	db := r.db.WithContext(ctx)
	stats := &domain.RouteStatistics{
		RoutesByVehicleType: make(map[string]int),
	}

	baseQuery := db.Model(&domain.Routing{}).Where("deleted_at IS NULL")
	if !filter.StartTime.IsZero() {
		baseQuery = baseQuery.Where("created_at >= ?", filter.StartTime)
	}
	if !filter.EndTime.IsZero() {
		baseQuery = baseQuery.Where("created_at <= ?", filter.EndTime)
	}

	var totalCount int64
	if err := baseQuery.Count(&totalCount).Error; err != nil {
		return nil, fmt.Errorf("failed to count routes: %w", err)
	}
	stats.TotalRoutes = int(totalCount)

	var totalDistance struct {
		Total float64
	}
	if err := baseQuery.Select("COALESCE(SUM(distance), 0) as total").Scan(&totalDistance).Error; err != nil {
		return nil, fmt.Errorf("failed to calculate total distance: %w", err)
	}
	stats.TotalDistance = totalDistance.Total

	var routes []struct {
		ID           uint
		VehicleTypes types.VehicleTypes `gorm:"type:vehicle_type[]"`
	}

	if err := baseQuery.Find(&routes).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch routes for vehicle type stats: %w", err)
	}

	vehicleTypeCounts := make(map[string]int)
	for _, route := range routes {
		for _, vType := range route.VehicleTypes {
			vehicleTypeCounts[string(vType)]++
		}
	}
	stats.RoutesByVehicleType = vehicleTypeCounts

	type PopularRouteResult struct {
		RouteID       uint
		UsageCount    int
		AverageRating float64
	}
	var popularRoutes []PopularRouteResult

	popularRoutesQuery := db.Table("routings").
		Select(`
            routings.id as route_id,
            COUNT(route_usages.id) as usage_count,
            COALESCE(AVG(NULLIF(route_usages.rating, 0)), 0) as average_rating
        `).
		Joins("LEFT JOIN route_usages ON routings.id = route_usages.route_id").
		Where("routings.deleted_at IS NULL")

	if !filter.StartTime.IsZero() {
		popularRoutesQuery = popularRoutesQuery.Where("route_usages.used_at >= ?", filter.StartTime)
	}
	if !filter.EndTime.IsZero() {
		popularRoutesQuery = popularRoutesQuery.Where("route_usages.used_at <= ?", filter.EndTime)
	}

	if err := popularRoutesQuery.
		Group("routings.id").
		Order("usage_count DESC, average_rating DESC").
		Limit(10).
		Scan(&popularRoutes).Error; err != nil {
		return nil, fmt.Errorf("failed to get popular routes: %w", err)
	}

	stats.PopularRoutes = make([]domain.PopularRoute, len(popularRoutes))
	for i, pr := range popularRoutes {
		var route domain.Routing
		if err := db.Preload("From").
			Preload("To").
			First(&route, pr.RouteID).Error; err != nil {
			return nil, fmt.Errorf("failed to get route details for ID %d: %w", pr.RouteID, err)
		}

		stats.PopularRoutes[i] = domain.PopularRoute{
			Route:         &route,
			UsageCount:    pr.UsageCount,
			AverageRating: pr.AverageRating,
		}
	}

	if filter.VehicleType != "" {
		filteredStats := &domain.RouteStatistics{
			RoutesByVehicleType: map[string]int{
				string(filter.VehicleType): stats.RoutesByVehicleType[string(filter.VehicleType)],
			},
		}

		var vehicleTypeStats struct {
			Count    int64
			Distance float64
		}

		vehicleTypeQuery := baseQuery.
			Where("? = ANY(vehicle_types)", filter.VehicleType).
			Select("COUNT(*) as count, COALESCE(SUM(distance), 0) as distance")

		if err := vehicleTypeQuery.Scan(&vehicleTypeStats).Error; err != nil {
			return nil, fmt.Errorf("failed to get vehicle type specific stats: %w", err)
		}

		filteredStats.TotalRoutes = int(vehicleTypeStats.Count)
		filteredStats.TotalDistance = vehicleTypeStats.Distance

		var filteredPopularRoutes []domain.PopularRoute
		for _, pr := range stats.PopularRoutes {
			if pr.Route.SupportsVehicleType(filter.VehicleType) {
				filteredPopularRoutes = append(filteredPopularRoutes, pr)
			}
		}
		filteredStats.PopularRoutes = filteredPopularRoutes

		return filteredStats, nil
	}

	return stats, nil
}
