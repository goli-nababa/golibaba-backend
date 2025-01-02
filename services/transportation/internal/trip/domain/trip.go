package domain

import (
	"strings"
	"time"
	companyDomain "transportation/internal/company/domain"
)

type (
	TripId           uint
	VehicleId        uint
	VehicleRequestId uint
)

type Trip struct {
	ID                       TripId                        `json:"id"`
	CreatedAt                time.Time                     `json:"created_at"`
	UpdatedAt                time.Time                     `json:"updated_at"`
	DeletedAt                *time.Time                    `json:"deleted_at,omitempty"`
	Title                    string                        `json:"title"`
	OriginStationId          uint                          `json:"origin_station_id"`
	DestinationStationId     uint                          `json:"destination_station_id"`
	CompanyId                companyDomain.CompanyId       `json:"company_id"`
	Company                  companyDomain.Company         `json:"company"`
	VehicleId                VehicleId                     `json:"vehicle_id"`
	StartTime                time.Time                     `json:"start_time"`
	EndTime                  *time.Time                    `json:"end_time"`
	ExpectedEndTime          *time.Time                    `json:"expected_end_time"`
	PassengersCountLimit     uint                          `json:"passengers_count_limit"`
	NormalCost               uint                          `json:"normal_cost"`
	AgencyCost               uint                          `json:"agency_cost"`
	AgencyReleaseTime        time.Time                     `json:"agency_release_time"`
	ReleaseTime              time.Time                     `json:"release_time"`
	TechTeamId               companyDomain.TechnicalTeamId `json:"tech_team_id"`
	TechnicalTeam            companyDomain.TechnicalTeam   `json:"technical_team"`
	TechTeamConfirmationTime *time.Time                    `json:"tech_team_confirmation_time"`
	EndConfirmationTime      *time.Time                    `json:"end_confirmation_time"`
}

type CreateTripRequest struct {
	Title                string                        `json:"title"`
	OriginStationId      uint                          `json:"origin_station_id"`
	DestinationStationId uint                          `json:"destination_station_id"`
	CompanyId            companyDomain.CompanyId       `json:"company_id"`
	StartTime            time.Time                     `json:"start_time"`
	PassengersCountLimit uint                          `json:"passengers_count_limit"`
	NormalCost           uint                          `json:"normal_cost"`
	AgencyCost           uint                          `json:"agency_cost"`
	AgencyReleaseTime    time.Time                     `json:"agency_release_time"`
	ReleaseTime          time.Time                     `json:"release_time"`
	TechTeamId           companyDomain.TechnicalTeamId `json:"tech_team_id"`
}

type UpdateTripRequest struct {
	Title                string                        `json:"title"`
	OriginStationId      uint                          `json:"origin_station_id"`
	DestinationStationId uint                          `json:"destination_station_id"`
	CompanyId            companyDomain.CompanyId       `json:"company_id"`
	StartTime            time.Time                     `json:"start_time"`
	PassengersCountLimit uint                          `json:"passengers_count_limit"`
	NormalCost           uint                          `json:"normal_cost"`
	AgencyCost           uint                          `json:"agency_cost"`
	AgencyReleaseTime    time.Time                     `json:"agency_release_time"`
	ReleaseTime          time.Time                     `json:"release_time"`
	TechTeamId           companyDomain.TechnicalTeamId `json:"tech_team_id"`
}

type GetTripsRequest struct {
	CompanyId            companyDomain.CompanyId       `json:"company_id"`
	TechTeamId           companyDomain.TechnicalTeamId `json:"tech_team_id"`
	OriginStationId      uint                          `json:"origin_station_id"`
	DestinationStationId uint                          `json:"destination_station_id"`
	FromStartTime        *time.Time                    `json:"from_start_time"`
	ToStartTime          *time.Time                    `json:"to_start_time"`
	VehicleId            VehicleId                     `json:"vehicle_id"`
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

type CustomDate struct {
	time.Time
}

func (cd *CustomDate) UnmarshalJSON(b []byte) error {
	str := strings.Trim(string(b), `"`)
	t, err := time.Parse("2006-01-02", str)
	if err != nil {
		return err
	}
	cd.Time = t
	return nil
}

type CreateVehicleRequest struct {
	TripId              TripId      `json:"trip_id"`
	VehicleTypeId       uint        `json:"vehicle_type_id"`
	VehicleCost         uint        `json:"vehicle_cost"`
	VehicleCreationDate *CustomDate `json:"vehicle_creation_date,omitempty"`
}
type GetVehicleRequests struct {
	TripId TripId `json:"trip_id"`
}
