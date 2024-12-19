package domain

import (
	hotelDomain "hotels-service/internal/hotel/domain"
	RateDomain "hotels-service/internal/rate/domain"
	"time"

	"github.com/google/uuid"
)

type RoomID = uuid.UUID
type Room struct {
	ID         RoomID
	HotelID    hotelDomain.HotelID
	RateID     RateDomain.RateID
	RoomNumber string
	Capacity   uint
	Features   []string
	IsAvilabe  bool
	CreateAt   time.Time
	EditedAt   time.Time
	DeletedAt  time.Time
}

type RoomFilterItem struct{}

func (r *Room) Validation() {
	//TODO
}
