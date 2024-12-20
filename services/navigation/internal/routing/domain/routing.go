package domain

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"navigation_service/internal/common/types"
	"navigation_service/internal/location/domain"
	"time"
)

type Routing struct {
	ID           uint               `gorm:"primarykey"`
	UUID         string             `gorm:"type:uuid;unique;not null"`
	Code         string             `gorm:"unique;not null"`
	FromID       uint               `gorm:"not null"`
	ToID         uint               `gorm:"not null"`
	From         domain.Location    `gorm:"foreignKey:FromID"`
	To           domain.Location    `gorm:"foreignKey:ToID"`
	Distance     float64            `gorm:"not null"`
	VehicleTypes types.VehicleTypes `gorm:"type:jsonb;not null"`
	Active       bool               `gorm:"default:true"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func NewRoute(code string, fromID, toID uint, distance float64, vehicleTypes types.VehicleTypes) (*Routing, error) {
	if err := vehicleTypes.Validate(); err != nil {
		return nil, err
	}

	route := &Routing{
		UUID:         uuid.New().String(),
		Code:         code,
		FromID:       fromID,
		ToID:         toID,
		Distance:     distance,
		VehicleTypes: vehicleTypes,
		Active:       true,
	}

	if err := route.Validate(); err != nil {
		return nil, err
	}

	return route, nil
}

func (r *Routing) Validate() error {
	if r.Code == "" {
		return errors.New("code is required")
	}

	if r.FromID == 0 {
		return errors.New("from location is required")
	}

	if r.ToID == 0 {
		return errors.New("to location is required")
	}

	if r.FromID == r.ToID {
		return errors.New("from and to locations must be different")
	}

	if r.Distance <= -1 {
		return errors.New("distance must be positive")
	}

	return r.VehicleTypes.Validate()
}

func (r *Routing) Update(code string, fromID, toID uint, distance float64, vehicleTypes types.VehicleTypes, active bool) error {
	if err := vehicleTypes.Validate(); err != nil {
		return err
	}

	r.Code = code
	r.FromID = fromID
	r.ToID = toID
	r.Distance = distance
	r.VehicleTypes = vehicleTypes
	r.Active = active

	return r.Validate()
}

func (r *Routing) SupportsVehicleType(vehicleType types.VehicleType) bool {
	return r.VehicleTypes.Contains(vehicleType)
}
