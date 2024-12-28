package types

import "time"

type Hotel struct {
	ID        string
	Name      string
	Address   string
	Rating    uint
	Amenities []string
	OwnerID   string
	CreatedAt time.Time
	UpdateAt  time.Time
	DeletedAt time.Time
}
