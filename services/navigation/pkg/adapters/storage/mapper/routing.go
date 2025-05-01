package mapper

import (
	"navigation_service/internal/common/types"
	"navigation_service/internal/routing/domain"
	storageTypes "navigation_service/pkg/adapters/storage/types"
)

func RouteToStorage(route *domain.Routing) *storageTypes.Route {
	if route == nil {
		return nil
	}

	vehicleTypes := make([]string, len(route.VehicleTypes))
	for i, vt := range route.VehicleTypes {
		vehicleTypes[i] = string(vt)
	}

	return &storageTypes.Route{
		ID:           uint(route.ID),
		UUID:         route.UUID,
		Code:         route.Code,
		FromID:       route.FromID,
		ToID:         route.ToID,
		Distance:     route.Distance,
		VehicleTypes: vehicleTypes,
		Active:       route.Active,
		CreatedAt:    route.CreatedAt,
		UpdatedAt:    route.UpdatedAt,
	}
}

func RouteFromStorage(route *storageTypes.Route) *domain.Routing {
	if route == nil {
		return nil
	}

	vehicleTypes := make([]types.VehicleType, len(route.VehicleTypes))
	for i, vt := range route.VehicleTypes {
		vehicleTypes[i] = types.VehicleType(vt)
	}

	return &domain.Routing{
		ID:           domain.RoutingID(route.ID),
		UUID:         route.UUID,
		Code:         route.Code,
		FromID:       route.FromID,
		ToID:         route.ToID,
		Distance:     route.Distance,
		VehicleTypes: vehicleTypes,
		Active:       route.Active,
		CreatedAt:    route.CreatedAt,
		UpdatedAt:    route.UpdatedAt,
	}
}

func RoutesFromStorage(routes []storageTypes.Route) []domain.Routing {
	result := make([]domain.Routing, len(routes))
	for i, route := range routes {
		result[i] = *RouteFromStorage(&route)
	}
	return result
}
