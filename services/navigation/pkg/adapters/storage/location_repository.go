package storage

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"navigation_service/internal/common/types"
	"navigation_service/internal/location/domain"
	"navigation_service/internal/location/port"
)

type locationRepository struct {
	db *gorm.DB
}

func NewLocationRepository(db *gorm.DB) port.Repo {
	return &locationRepository{db: db}
}

func (r *locationRepository) Create(ctx context.Context, location *domain.Location) error {
	if err := location.Validate(); err != nil {
		return fmt.Errorf("invalid location: %w", err)
	}

	result := r.db.WithContext(ctx).Create(location)
	if result.Error != nil {
		return fmt.Errorf("failed to create location: %w", result.Error)
	}

	return nil
}

func (r *locationRepository) Update(ctx context.Context, location *domain.Location) error {
	if err := location.Validate(); err != nil {
		return fmt.Errorf("invalid location: %w", err)
	}

	result := r.db.WithContext(ctx).Model(&domain.Location{}).
		Where("id = ?", location.ID).
		Updates(map[string]interface{}{
			"name":      location.Name,
			"type":      location.Type,
			"address":   location.Address,
			"latitude":  location.Latitude,
			"longitude": location.Longitude,
			"active":    location.Active,
		})

	if result.Error != nil {
		return fmt.Errorf("failed to update location: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("location not found: %d", location.ID)
	}

	return nil
}

func (r *locationRepository) Delete(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).Delete(&domain.Location{}, id)
	if result.Error != nil {
		return fmt.Errorf("failed to delete location: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("location not found: %d", id)
	}

	return nil
}

func (r *locationRepository) GetByID(ctx context.Context, id uint) (*domain.Location, error) {
	var location domain.Location

	result := r.db.WithContext(ctx).First(&location, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get location: %w", result.Error)
	}

	return &location, nil
}

func (r *locationRepository) GetByType(ctx context.Context, locationType types.LocationType) ([]domain.Location, error) {
	var locations []domain.Location

	result := r.db.WithContext(ctx).
		Where("type = ? AND active = true", locationType).
		Find(&locations)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to get locations by type: %w", result.Error)
	}

	return locations, nil
}

func (r *locationRepository) List(ctx context.Context, activeOnly bool) ([]domain.Location, error) {
	var locations []domain.Location
	query := r.db.WithContext(ctx)

	if activeOnly {
		query = query.Where("active = true")
	}

	result := query.Find(&locations)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to list locations: %w", result.Error)
	}

	return locations, nil
}
