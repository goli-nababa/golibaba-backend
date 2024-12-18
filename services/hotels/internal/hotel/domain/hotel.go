package domain

import (
	UserDomain "hotels-service/internal/user/domain"
	"time"
)

type HotelID string
type Hotel struct {
	ID         HotelID
	Name       string
	Address    string
	Rating     string
	Amentities []string
	OwnerID    UserDomain.UserID
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
