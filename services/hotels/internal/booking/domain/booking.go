package domain

import "time"

type StatusType uint8

const (
	StatusTypeUnknow StatusType = iota
	StatusTypeActive
	StatusTypeCancelled
	StatusTypeDone
)

type BookingID string
type Booking struct {
	ID           BookingID
	UserID       string
	CheckInDate  time.Time
	CheckOutDate time.Time
	TotalPrice   float64
	Status       StatusType
}

type BookingFilterItem struct{}

func (Booking) Validation() {
	//TODO
}
