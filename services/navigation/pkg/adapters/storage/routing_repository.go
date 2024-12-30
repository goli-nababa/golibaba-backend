package storage

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"navigation_service/internal/common/types"
	"navigation_service/internal/routing/domain"
	"navigation_service/internal/routing/port"
	"navigation_service/pkg/adapters/storage/mapper"
	storageTypes "navigation_service/pkg/adapters/storage/types"
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

	storageRoute := mapper.RouteToStorage(route)
	result := r.db.WithContext(ctx).Create(storageRoute)
	if result.Error != nil {
		return fmt.Errorf("failed to create route: %w", result.Error)
	}

	// Update domain model with generated ID and timestamps
	*route = *mapper.RouteFromStorage(storageRoute)
	return nil
}

func (r *routingRepository) Update(ctx context.Context, route *domain.Routing) error {
	if err := route.Validate(); err != nil {
		return fmt.Errorf("invalid route: %w", err)
	}

	storageRoute := mapper.RouteToStorage(route)
	result := r.db.WithContext(ctx).Model(&storageTypes.Route{}).
		Where("id = ?", route.ID).
		Updates(map[string]interface{}{
			"code":          storageRoute.Code,
			"from_id":       storageRoute.FromID,
			"to_id":         storageRoute.ToID,
			"distance":      storageRoute.Distance,
			"vehicle_types": storageRoute.VehicleTypes,
			"active":        storageRoute.Active,
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
	result := r.db.WithContext(ctx).Delete(&storageTypes.Route{}, id)
	if result.Error != nil {
		return fmt.Errorf("failed to delete route: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("route not found: %d", id)
	}

	return nil
}

func (r *routingRepository) GetByID(ctx context.Context, id uint) (*domain.Routing, error) {
	var route storageTypes.Route

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

	return mapper.RouteFromStorage(&route), nil
}

func (r *routingRepository) GetByUUID(ctx context.Context, uuid string) (*domain.Routing, error) {
	var route storageTypes.Route

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

	return mapper.RouteFromStorage(&route), nil
}

func (r *routingRepository) GetByCode(ctx context.Context, code string) (*domain.Routing, error) {
	var route storageTypes.Route

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

	return mapper.RouteFromStorage(&route), nil
}

func (r *routingRepository) FindRoutes(ctx context.Context, filter domain.RouteFilter) ([]domain.Routing, error) {
	if err := filter.Validate(); err != nil {
		return nil, fmt.Errorf("invalid filter: %w", err)
	}

	var routes []storageTypes.Route
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

	result := query.Find(&routes)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to find routes: %w", result.Error)
	}

	return mapper.RoutesFromStorage(routes), nil
}

func (r *routingRepository) GetStatistics(ctx context.Context, filter domain.StatisticsFilter) (*domain.RouteStatistics, error) {
	stats := &domain.RouteStatistics{
		RoutesByVehicleType: make(map[string]int),
	}

	baseQuery := r.db.WithContext(ctx).Model(&storageTypes.Route{}).Where("deleted_at IS NULL")
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

	var routes []storageTypes.Route
	if err := baseQuery.Find(&routes).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch routes for vehicle type stats: %w", err)
	}

	vehicleTypeCounts := make(map[string]int)
	for _, route := range routes {
		for _, vType := range route.VehicleTypes {
			vehicleTypeCounts[vType]++
		}
	}
	stats.RoutesByVehicleType = vehicleTypeCounts

	return stats, nil
}
