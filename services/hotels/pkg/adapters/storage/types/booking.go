package types

import (
	"time"
)

type Booking struct {
	ID           string
	UserID       string
	RoomID       string
	CheckInDate  time.Time
	CheckOutDate time.Time
	TotalPrice   float64
	Status       uint8
	CreateAt     time.Time
	updatedAt    time.Time
	DeletedAt    time.Time
}
