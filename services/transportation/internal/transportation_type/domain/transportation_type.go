package domain

import (
	"time"
)

type (
	TransportationTypeId uint
)

type TransportationType struct {
	ID        TransportationTypeId `json:"id"`
	CreatedAt time.Time            `json:"created_at"`
	UpdatedAt time.Time            `json:"updated_at"`
	DeletedAt *time.Time           `json:"deleted_at,omitempty"`
	Name      string               `json:"name"`
}
