package domain

import (
	hotelDomain "hotels-service/internal/hotel/domain"
	RateDomain "hotels-service/internal/rate/domain"
	"time"
)

type RoomID string
type Room struct {
	ID         RoomID
	HotelID    hotelDomain.HotelId
	RateID     RateDomain.RateID
	RoomNumber string
	Capacity   uint
	Features   []string
	IsAvilabe  bool
	CreateAt   time.Time
	EditedAt   time.Time
	DeletedAt  time.Time
}

func (r *Room) Validation() {
	//TODO
}
