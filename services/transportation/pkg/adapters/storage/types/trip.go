package types

import (
	"time"

	"gorm.io/gorm"
)

type Trip struct {
	ID                       uint           `gorm:"primarykey" json:"id"`
	CreatedAt                time.Time      `json:"created_at"`
	UpdatedAt                time.Time      `json:"updated_at"`
	DeletedAt                gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	Title                    string         `json:"title"`
	OriginStationId          uint           `json:"origin_station_id"`
	DestinationStationId     uint           `json:"destination_station_id"`
	CompanyId                uint           `json:"company_id"`
	Company                  Company        `gorm:"foreignKey:CompanyId" json:"company"`
	VehicleId                uint           `json:"vehicle_id"`
	StartTime                time.Time      `json:"start_time"`
	EndTime                  *time.Time     `json:"end_time"`
	PassengersCountLimit     uint           `json:"passengers_count_limit"`
	NormalCost               uint           `json:"normal_cost"`
	AgencyCost               uint           `json:"agency_cost"`
	AgencyReleaseTime        time.Time      `json:"agency_release_time"`
	ReleaseTime              time.Time      `json:"release_time"`
	TechTeamId               uint           `json:"tech_team_id"`
	TechnicalTeam            TechnicalTeam  `gorm:"foreignKey:TechTeamId" json:"technical_team"`
	TechTeamConfirmationTime *time.Time     `json:"tech_team_confirmation_time"`
	EndConfirmationTime      *time.Time     `json:"end_confirmation_time"`
}
