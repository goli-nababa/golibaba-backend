package types

import (
	"time"

	"gorm.io/gorm"
)

type Vehicle struct {
	ID                  uint           `gorm:"primarykey" json:"id"`
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           time.Time      `json:"updated_at"`
	DeletedAt           gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	Name                string         `json:"name"`
	OwnerId             uint           `json:"owner_id"`
	TechnicalTeamNumber uint           `json:"technical_team_number"`
	Speed               uint           `json:"speed"`
	VehicleTypeId       uint           `json:"vehicle_type_id"`
	VehicleCreationDate time.Time      `json:"vehicle_creation_date"`
	Cost                uint           `json:"cost"`
	PassengerCapacity   uint           `json:"passenger_capacity"`
}
