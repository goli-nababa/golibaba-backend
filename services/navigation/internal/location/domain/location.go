package domain

import (
	"errors"
	"navigation_service/internal/common/types"
	"time"
)

type LocationID uint

type Location struct {
	ID        LocationID
	Name      string
	Type      types.LocationType
	Address   string
	Latitude  float64
	Longitude float64
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewLocation(name string, locationType types.LocationType, address string, lat, lng float64) (*Location, error) {
	location := &Location{
		Name:      name,
		Type:      locationType,
		Address:   address,
		Latitude:  lat,
		Longitude: lng,
		Active:    true,
	}

	if err := location.Validate(); err != nil {
		return nil, err
	}

	return location, nil
}

func (l *Location) Validate() error {
	if l.Name == "" {
		return errors.New("name is required")
	}

	if err := l.Type.Validate(); err != nil {
		return err
	}

	if l.Address == "" {
		return errors.New("address is required")
	}

	if l.Latitude < -90 || l.Latitude > 90 {
		return errors.New("latitude must be between -90 and 90")
	}

	if l.Longitude < -180 || l.Longitude > 180 {
		return errors.New("longitude must be between -180 and 180")
	}

	return nil
}

func (l *Location) Update(name string, locationType types.LocationType, address string, lat, lng float64, active bool) error {
	l.Name = name
	l.Type = locationType
	l.Address = address
	l.Latitude = lat
	l.Longitude = lng
	l.Active = active

	return l.Validate()
}
