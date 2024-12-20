package domain

import (
	"navigation_service/internal/common/types"
)

type RouteFilter struct {
	FromID      uint              `json:"from_id"`
	ToID        uint              `json:"to_id"`
	VehicleType types.VehicleType `json:"vehicle_type"`
	ActiveOnly  bool              `json:"active_only"`
	PageSize    int               `json:"page_size"`
	PageNumber  int               `json:"page_number"`
}

func NewRouteFilter() RouteFilter {
	return RouteFilter{
		PageSize:   10,
		PageNumber: 1,
	}
}

func (f *RouteFilter) Validate() error {
	if f.PageSize < 1 {
		f.PageSize = 10
	}
	if f.PageNumber < 1 {
		f.PageNumber = 1
	}

	if f.VehicleType != "" {
		if err := f.VehicleType.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// GetOffset calculates the offset for pagination
func (f *RouteFilter) GetOffset() int {
	return (f.PageNumber - 1) * f.PageSize
}

// GetLimit returns the limit for pagination
func (f *RouteFilter) GetLimit() int {
	return f.PageSize
}
