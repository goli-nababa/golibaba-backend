package types

import "github.com/goli-nababa/golibaba-backend/common"

type Statement struct {
	ID          uint                       `gorm:"primaryKey"`
	UserID      uint                       `gorm:"index"`
	Action      common.TypeStatementAction `gorm:"index"`
	Permissions []string                   `gorm:"type:json"`
}

func (Statement) TableName() string {
	return "statements"
}
