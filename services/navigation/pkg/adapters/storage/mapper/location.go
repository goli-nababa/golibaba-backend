package mapper

import (
	"navigation_service/internal/common/types"
	"navigation_service/internal/location/domain"
	storageTypes "navigation_service/pkg/adapters/storage/types"
)

func LocationToStorage(location *domain.Location) *storageTypes.Location {
	if location == nil {
		return nil
	}

	return &storageTypes.Location{
		ID:        uint(location.ID),
		Name:      location.Name,
		Type:      string(location.Type),
		Address:   location.Address,
		Latitude:  location.Latitude,
		Longitude: location.Longitude,
		Active:    location.Active,
		CreatedAt: location.CreatedAt,
		UpdatedAt: location.UpdatedAt,
	}
}

func LocationFromStorage(location *storageTypes.Location) *domain.Location {
	if location == nil {
		return nil
	}

	return &domain.Location{
		ID:        domain.LocationID(location.ID),
		Name:      location.Name,
		Type:      types.LocationType(location.Type),
		Address:   location.Address,
		Latitude:  location.Latitude,
		Longitude: location.Longitude,
		Active:    location.Active,
		CreatedAt: location.CreatedAt,
		UpdatedAt: location.UpdatedAt,
	}
}

func LocationsFromStorage(locations []storageTypes.Location) []domain.Location {
	result := make([]domain.Location, len(locations))
	for i, loc := range locations {
		result[i] = *LocationFromStorage(&loc)
	}
	return result
}
