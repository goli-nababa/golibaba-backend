package domain

import (
	"time"
	transportationDomain "transportation/internal/transportation_type/domain"
)

type (
	CompanyId             uint
	TechnicalTeamId       uint
	TechnicalTeamMemberId uint
)

type Company struct {
	ID                   CompanyId                                 `json:"id"`
	CreatedAt            time.Time                                 `json:"created_at"`
	UpdatedAt            time.Time                                 `json:"updated_at"`
	DeletedAt            *time.Time                                `json:"deleted_at,omitempty"`
	Name                 string                                    `json:"name"`
	OwnerId              uint                                      `json:"owner_id"`
	TransportationTypeId transportationDomain.TransportationTypeId `json:"transportation_type_id"`
	TransportationType   transportationDomain.TransportationType   `json:"transportation_type"`
	TechnicalTeams       []TechnicalTeam                           `json:"technical_teams"`
}

type TechnicalTeam struct {
	ID                   TechnicalTeamId       `json:"id"`
	CreatedAt            time.Time             `json:"created_at"`
	UpdatedAt            time.Time             `json:"updated_at"`
	DeletedAt            *time.Time            `json:"deleted_at,omitempty"`
	Name                 string                `json:"name"`
	CompanyId            CompanyId             `json:"company_id"`
	Company              Company               `json:"company"`
	TechnicalTeamMembers []TechnicalTeamMember `json:"technical_team_members"`
}

type TechnicalTeamMember struct {
	ID              TechnicalTeamMemberId `json:"id"`
	CreatedAt       time.Time             `json:"created_at"`
	UpdatedAt       time.Time             `json:"updated_at"`
	DeletedAt       *time.Time            `json:"deleted_at,omitempty"`
	TechnicalTeamId TechnicalTeamId       `json:"technical_team_id"`
	MemberId        uint                  `json:"member_id"`
	TechnicalTeam   TechnicalTeam         `json:"technical_team"`
}
