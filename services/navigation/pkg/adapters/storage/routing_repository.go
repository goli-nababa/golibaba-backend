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
