package domain

import (
	"errors"
	hotelDomain "hotels-service/internal/hotel/domain"
	RateDomain "hotels-service/internal/rate/domain"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidRoomID     = errors.New("invalid room ID")
	ErrInvalidHotelID    = errors.New("invalid hotel ID")
	ErrInvalidRateID     = errors.New("invalid rate ID")
	ErrInvalidRoomNumber = errors.New("invalid room number")
	ErrInvalidCapacity   = errors.New("invalid room capacity")
	ErrInvalidFeatures   = errors.New("invalid room features")
)

type StatusType uint8

const (
	StatusTypeUnknow StatusType = iota
	StatusTypeAvailable
	StatusTypeBooked
	StatusTypeMaintenance
	StatusTypeCleaning
	StatusTypeReserved
	StatusTypeBlocked
	StatusTypeOutOfOrder
)

type RoomID = uuid.UUID
type Room struct {
	ID         RoomID
	HotelID    hotelDomain.HotelID
	RateID     RateDomain.RateID
	RoomNumber string
	Capacity   uint
	Features   []string
	Status     StatusType
	CreateAt   time.Time
	DeletedAt  time.Time
}

type RoomFilterItem struct {
	HotelID    hotelDomain.HotelID
	RateID     RateDomain.RateID
	RoomNumber string
	Capacity   uint
	Features   []string
	Status     StatusType
}

func ValidateID(ID uuid.UUID) error {
	if err := uuid.Validate(ID.String()); err != nil {
		return ErrInvalidRoomID
	}
	return nil
}

func validateHotelID(hotelID hotelDomain.HotelID) error {
	if hotelID == uuid.Nil {
		return ErrInvalidHotelID
	}
	return nil
}

func validateRateID(rateID RateDomain.RateID) error {
	if rateID == uuid.Nil {
		return ErrInvalidRateID
	}
	return nil
}

func validateRoomNumber(roomNumber string) error {
	if roomNumber == "" {
		return ErrInvalidRoomNumber
	}
	return nil
}

func validateCapacity(capacity uint) error {
	if capacity <= 0 {
		return ErrInvalidCapacity
	}
	return nil
}

func validateFeatures(features []string) error {
	if features == nil {
		return ErrInvalidFeatures
	}
	return nil
}

func (r *Room) Validate() error {
	if err := ValidateID(r.ID); err != nil {
		return err
	}

	if err := validateHotelID(r.HotelID); err != nil {
		return err
	}

	if err := validateRateID(r.RateID); err != nil {
		return err
	}

	if err := validateRoomNumber(r.RoomNumber); err != nil {
		return err
	}

	if err := validateCapacity(r.Capacity); err != nil {
		return err
	}

	if err := validateFeatures(r.Features); err != nil {
		return err
	}

	return nil
}
