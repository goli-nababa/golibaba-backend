package domain

import (
	"errors"
	roomDomain "hotels-service/internal/room/domain"
	userDomain "hotels-service/internal/user/domain"
	"time"

	"github.com/google/uuid"
)

var (
	ErrBookingIDRequired           = errors.New("booking id is required")
	ErrBookingUserIDRequired       = errors.New("booking user id is required")
	ErrBookingRoomIDRequired       = errors.New("booking room id is required")
	ErrBookingCheckInDateRequired  = errors.New("booking check-in date is required")
	ErrBookingCheckOutDateRequired = errors.New("booking check-out date is required")
	ErrBookingInvalidDateRange     = errors.New("check-in date must be before check-out date")
	ErrBookingInvalidPrice         = errors.New("booking price must be greater than 0")
	ErrBookingStatusRequired       = errors.New("booking status is required")
)

type StatusType uint8

const (
	StatusTypeUnknow StatusType = iota
	StatusTypeActive
	StatusTypeCancelled
	StatusTypeDone
)

type BookingID = uuid.UUID
type Booking struct {
	ID           BookingID
	UserID       userDomain.UserID
	RoomID       roomDomain.RoomID
	CheckInDate  time.Time
	CheckOutDate time.Time
	TotalPrice   float64
	Status       StatusType
	CreateAt     time.Time
	DeletedAt    time.Time
}

type BookingFilterItem struct {
	UserID       userDomain.UserID
	RoomID       roomDomain.RoomID
	CheckInDate  time.Time
	CheckOutDate time.Time
	Status       StatusType
}

func ValidateID(ID BookingID) error {
	if err := uuid.Validate(ID.String()); err != nil {
		return ErrBookingIDRequired
	}
	return nil
}

func validateUserID(userID userDomain.UserID) error {
	if err := uuid.Validate(userID.String()); err != nil {
		return ErrBookingUserIDRequired
	}
	return nil
}

func validateRoomID(roomID roomDomain.RoomID) error {
	if err := uuid.Validate(roomID.String()); err != nil {
		return ErrBookingRoomIDRequired
	}
	return nil
}

func validateDates(checkInDate time.Time, checkOutDate time.Time) error {
	if checkInDate.IsZero() {
		return ErrBookingCheckInDateRequired
	}
	if checkOutDate.IsZero() {
		return ErrBookingCheckOutDateRequired
	}
	if checkInDate.After(checkOutDate) {
		return ErrBookingInvalidDateRange
	}
	return nil
}

func validatePrice(totalPrice float64) error {
	if totalPrice <= 0 {
		return ErrBookingInvalidPrice
	}
	return nil
}

func validateStatus(status StatusType) error {
	if status == StatusTypeUnknow {
		return ErrBookingStatusRequired
	}
	return nil
}

func (b *Booking) Validate() error {
	if err := ValidateID(b.ID); err != nil {
		return err
	}
	if err := validateUserID(b.UserID); err != nil {
		return err
	}
	if err := validateRoomID(b.RoomID); err != nil {
		return err
	}
	if err := validateDates(b.CheckInDate, b.CheckOutDate); err != nil {
		return err
	}
	if err := validatePrice(b.TotalPrice); err != nil {
		return err
	}
	if err := validateStatus(b.Status); err != nil {
		return err
	}
	return nil
}
