package domain

import (
	"errors"
	userDomain "hotels-service/internal/user/domain"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidHotelID        = errors.New("invalid hotel id")
	ErrInvalidHotelName      = errors.New("invalid hotel name")
	ErrInvalidHotelAddress   = errors.New("invalid hotel address")
	ErrInvalidHotelRating    = errors.New("invalid hotel rating")
	ErrInvalidHotelAmenities = errors.New("invalid hotel amenities")
	ErrInvalidHotelOwnerID   = errors.New("invalid hotel owner id")
)

type HotelID = uuid.UUID
type Hotel struct {
	ID         HotelID
	Name       string
	Address    string
	Rating     uint
	Amentities []string
	OwnerID    userDomain.UserID
	CreatedAt  time.Time
	DeletedAt  time.Time
}

type HotelFilterItem struct {
	Name       string
	Address    string
	Rating     string
	Amentities []string
}

func ValidateID(ID uuid.UUID) error {
	if ID == uuid.Nil {
		return ErrInvalidHotelID
	}
	return nil
}

func ValidateName(name string) error {
	if name == "" {
		return ErrInvalidHotelName
	}
	return nil
}

func ValidateAddress(address string) error {
	if address == "" {
		return ErrInvalidHotelAddress
	}
	return nil
}

func ValidateRating(rating uint) error {
	if rating <= 0 {
		return ErrInvalidHotelRating
	}
	return nil
}

func ValidateOwnerID(ownerID userDomain.UserID) error {
	if ownerID == uuid.Nil {
		return ErrInvalidHotelOwnerID
	}
	return nil
}

func (h *Hotel) Validate() error {
	if err := ValidateID(h.ID); err != nil {
		return err
	}

	if err := ValidateName(h.Name); err != nil {
		return err
	}

	if err := ValidateAddress(h.Address); err != nil {
		return err
	}

	if err := ValidateRating(h.Rating); err != nil {
		return err
	}

	if err := ValidateOwnerID(h.OwnerID); err != nil {
		return err
	}

	return nil
}
