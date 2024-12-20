package migrations

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Location struct {
	ID        uint    `gorm:"primarykey"`
	Name      string  `gorm:"type:varchar(255);not null"`
	Type      string  `gorm:"type:location_type;not null"`
	Address   string  `gorm:"type:text;not null"`
	Latitude  float64 `gorm:"type:double precision;not null"`
	Longitude float64 `gorm:"type:double precision;not null"`
	Active    bool    `gorm:"default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Routing struct {
	ID           uint      `gorm:"primarykey"`
	UUID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Code         string    `gorm:"type:varchar(50);uniqueIndex;not null"`
	FromID       uint      `gorm:"not null"`
	ToID         uint      `gorm:"not null"`
	From         Location  `gorm:"foreignKey:FromID"`
	To           Location  `gorm:"foreignKey:ToID"`
	Distance     float64   `gorm:"type:double precision;not null"`
	VehicleTypes []string  `gorm:"type:vehicle_type[]"`
	Active       bool      `gorm:"default:true"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

// TableName overrides the table name
func (Location) TableName() string {
	return "locations"
}

func (Routing) TableName() string {
	return "routings"
}
