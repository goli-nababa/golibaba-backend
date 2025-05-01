package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Route struct {
	ID           uint         `gorm:"primarykey"`
	UUID         string       `gorm:"type:uuid;default:gen_random_uuid()"`
	Code         string       `gorm:"type:varchar(50);uniqueIndex;not null"`
	FromID       uint         `gorm:"not null"`
	ToID         uint         `gorm:"not null"`
	From         Location     `gorm:"foreignKey:FromID"`
	To           Location     `gorm:"foreignKey:ToID"`
	Distance     float64      `gorm:"type:double precision;not null"`
	VehicleTypes VehicleTypes `gorm:"type:jsonb;not null"`
	Active       bool         `gorm:"default:true"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type VehicleTypes []string

// Value implements the driver.Valuer interface
func (vt VehicleTypes) Value() (driver.Value, error) {
	if len(vt) == 0 {
		return json.Marshal([]string{})
	}
	return json.Marshal(vt)
}

// Scan implements the sql.Scanner interface
func (vt *VehicleTypes) Scan(value interface{}) error {
	if value == nil {
		vt = &VehicleTypes{}
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal JSONB value: %v", value)
	}

	return json.Unmarshal(bytes, &vt)
}

func (Route) TableName() string {
	return "routings"
}
