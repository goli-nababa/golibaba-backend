package domain

import (
	"time"
	companyDomain "transportation/internal/company/domain"
)

type (
	TripId           uint
	VehicleRequestId uint
)

type Trip struct {
	ID                       TripId                      `json:"id"`
	CreatedAt                time.Time                   `json:"created_at"`
	UpdatedAt                time.Time                   `json:"updated_at"`
	DeletedAt                *time.Time                  `json:"deleted_at,omitempty"`
	Title                    string                      `json:"title"`
	OriginStationId          uint                        `json:"origin_station_id"`
	DestinationStationId     uint                        `json:"destination_station_id"`
	CompanyId                uint                        `json:"company_id"`
	Company                  companyDomain.Company       `json:"company"`
	VehicleId                uint                        `json:"vehicle_id"`
	StartTime                time.Time                   `json:"start_time"`
	EndTime                  *time.Time                  `json:"end_time"`
	PassengersCountLimit     uint                        `json:"passengers_count_limit"`
	NormalCost               uint                        `json:"normal_cost"`
	AgencyCost               uint                        `json:"agency_cost"`
	AgencyReleaseTime        time.Time                   `json:"agency_release_time"`
	ReleaseTime              time.Time                   `json:"release_time"`
	TechTeamId               uint                        `json:"tech_team_id"`
	TechnicalTeam            companyDomain.TechnicalTeam `json:"technical_team"`
	TechTeamConfirmationTime *time.Time                  `json:"tech_team_confirmation_time"`
	EndConfirmationTime      *time.Time                  `json:"end_confirmation_time"`
}

type VehicleRequest struct {
	ID                  VehicleRequestId `json:"id"`
	CreatedAt           time.Time        `json:"created_at"`
	UpdatedAt           time.Time        `json:"updated_at"`
	DeletedAt           *time.Time       `json:"deleted_at,omitempty"`
	TripId              TripId           `json:"trip_id"`
	Trip                Trip             `json:"trip"`
	VehicleTypeId       uint             `json:"vehicle_type_id"`
	VehicleCost         uint             `json:"vehicle_cost"`
	VehicleCreationDate *time.Time       `json:"vehicle_creation_date,omitempty"`
}
