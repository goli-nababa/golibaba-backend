package dto

import "time"

type Trip struct {
	ID                   uint `json:"id"`
	PassengersCountLimit uint `json:"passengers_count_limit"`
}
type VehicleRequest struct {
	ID                  uint       `json:"id"`
	Trip                Trip       `json:"trip"`
	VehicleTypeId       uint       `json:"vehicle_type_id"`
	VehicleCost         uint       `json:"vehicle_cost"`
	VehicleCreationDate *time.Time `json:"vehicle_creation_date,omitempty"`
}
