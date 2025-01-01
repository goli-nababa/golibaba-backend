package types

import (
	"time"
)

type Hotel struct {
	ID        string      `gorm:"type:uuid;primary_key"`
	Name      string      `gorm:"not null"`
	Address   string      `gorm:"not null"`
	Rating    uint        `gorm:"not null"`
	Amenities StringArray `gorm:"type:json"`
	OwnerID   string      `gorm:"type:uuid;not null"`
	CreatedAt time.Time
	UpdateAt  time.Time
	DeletedAt time.Time
}
