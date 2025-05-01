package types

import (
	"time"

	"gorm.io/gorm"
)

type TechnicalTeam struct {
	ID                   uint                  `gorm:"primarykey" json:"id"`
	CreatedAt            time.Time             `json:"created_at"`
	UpdatedAt            time.Time             `json:"updated_at"`
	DeletedAt            gorm.DeletedAt        `gorm:"index" json:"deleted_at,omitempty"`
	Name                 string                `json:"name"`
	CompanyId            uint                  `json:"company_id"`
	Company              Company               `gorm:"foreignKey:CompanyId" json:"company"`
	TechnicalTeamMembers []TechnicalTeamMember `gorm:"foreignKey:TechnicalTeamId" json:"technical_team_members"`
}
