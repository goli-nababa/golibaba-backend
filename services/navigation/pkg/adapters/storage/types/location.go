package types

import (
	"gorm.io/gorm"
	"time"
)

type Location struct {
	ID        uint    `gorm:"primarykey"`
	Name      string  `gorm:"type:varchar(255);not null"`
	Type      string  `gorm:"type:varchar(50);not null"`
	Address   string  `gorm:"type:text;not null"`
	Latitude  float64 `gorm:"type:double precision;not null"`
	Longitude float64 `gorm:"type:double precision;not null"`
	Active    bool    `gorm:"default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Location) TableName() string {
	return "locations"
}
