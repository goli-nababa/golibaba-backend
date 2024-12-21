package types

import (
	"time"

	"gorm.io/gorm"
)

type VehicleRequest struct {
	ID                  uint           `gorm:"primarykey" json:"id"`
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           time.Time      `json:"updated_at"`
	DeletedAt           gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	TripId              uint           `json:"trip_id"`
	Trip                Trip           `gorm:"foreignKey:TripId" json:"trip"`
	VehicleTypeId       uint           `json:"vehicle_type_id"`
	VehicleCost         uint           `json:"vehicle_cost"`
	VehicleCreationDate *time.Time     `json:"vehicle_creation_date,omitempty"`
}
