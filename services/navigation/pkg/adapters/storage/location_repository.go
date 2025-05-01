package storage

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"navigation_service/internal/common/types"
	"navigation_service/internal/location/domain"
	"navigation_service/internal/location/port"
	"navigation_service/pkg/adapters/storage/mapper"
	storageTypes "navigation_service/pkg/adapters/storage/types"
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

	storageLocation := mapper.LocationToStorage(location)
	result := r.db.WithContext(ctx).Create(storageLocation)
	if result.Error != nil {
		return fmt.Errorf("failed to create location: %w", result.Error)
	}

	// Update domain model with generated ID and timestamps
	*location = *mapper.LocationFromStorage(storageLocation)
	return nil
}

func (r *locationRepository) Update(ctx context.Context, location *domain.Location) error {
	if err := location.Validate(); err != nil {
		return fmt.Errorf("invalid location: %w", err)
	}

	storageLocation := mapper.LocationToStorage(location)
	result := r.db.WithContext(ctx).Model(&storageTypes.Location{}).
		Where("id = ?", location.ID).
		Updates(map[string]interface{}{
			"name":      storageLocation.Name,
			"type":      storageLocation.Type,
			"address":   storageLocation.Address,
			"latitude":  storageLocation.Latitude,
			"longitude": storageLocation.Longitude,
			"active":    storageLocation.Active,
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
	result := r.db.WithContext(ctx).Delete(&storageTypes.Location{}, id)
	if result.Error != nil {
		return fmt.Errorf("failed to delete location: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("location not found: %d", id)
	}

	return nil
}

func (r *locationRepository) GetByID(ctx context.Context, id uint) (*domain.Location, error) {
	var location storageTypes.Location

	result := r.db.WithContext(ctx).First(&location, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get location: %w", result.Error)
	}

	return mapper.LocationFromStorage(&location), nil
}

func (r *locationRepository) GetByType(ctx context.Context, locationType types.LocationType) ([]domain.Location, error) {
	var locations []storageTypes.Location

	result := r.db.WithContext(ctx).
		Where("type = ? AND active = true", locationType).
		Find(&locations)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to get locations by type: %w", result.Error)
	}

	return mapper.LocationsFromStorage(locations), nil
}

func (r *locationRepository) List(ctx context.Context, activeOnly bool) ([]domain.Location, error) {
	var locations []storageTypes.Location
	query := r.db.WithContext(ctx)

	if activeOnly {
		query = query.Where("active = true")
	}

	result := query.Find(&locations)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to list locations: %w", result.Error)
	}

	return mapper.LocationsFromStorage(locations), nil
}
