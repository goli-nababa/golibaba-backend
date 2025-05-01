package types

import (
	"time"

	"gorm.io/gorm"
)

type TechnicalTeamMember struct {
	ID              uint           `gorm:"primarykey" json:"id"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	TechnicalTeamId uint           `json:"technical_team_id"`
	MemberId        uint           `json:"member_id"`
	TechnicalTeam   TechnicalTeam  `gorm:"foreignKey:TechnicalTeamId" json:"technical_team"`
}
