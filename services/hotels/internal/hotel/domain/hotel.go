package domain

import "time"

type HotelID string
type Hotel struct {
	ID         HotelID
	Name       string
	Address    string
	Rating     string
	Amentities []string
	OwnerID    uint
	CreatedAt  time.Time
	EditedAt   time.Time
	DeletedAt  time.Time
}

type HotelFilterItem struct {
	Name       string
	Address    string
	Rating     string
	Amentities []string
}

func (h *Hotel) Validation() {
	//TODO
}
