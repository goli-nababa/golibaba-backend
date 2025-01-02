package types

import (
	"time"

	"gorm.io/gorm"
)

type Company struct {
	ID                   uint               `gorm:"primarykey" json:"id"`
	CreatedAt            time.Time          `json:"created_at"`
	UpdatedAt            time.Time          `json:"updated_at"`
	DeletedAt            gorm.DeletedAt     `gorm:"index" json:"deleted_at,omitempty"`
	Name                 string             `json:"name"`
	OwnerId              uint               `json:"owner_id"`
	TransportationTypeId uint               `json:"transportation_type_id"`
	TransportationType   TransportationType `gorm:"foreignKey:TransportationTypeId" json:"transportation_type"`
	TechnicalTeams       []TechnicalTeam    `gorm:"foreignKey:CompanyId" json:"technical_teams"`
}
