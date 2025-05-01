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
	ID        HotelID
	Name      string
	Address   string
	Rating    uint
	Amenities []string
	OwnerID   userDomain.UserID
	CreatedAt time.Time
	DeletedAt time.Time
}

type HotelFilterItem struct {
	Name      string
	Address   string
	Rating    string
	Amenities []string
}

func ValidateID(ID uuid.UUID) error {
	if err := uuid.Validate(ID.String()); err != nil {
		return ErrInvalidHotelID
	}
	return nil
}

func validateName(name string) error {
	if name == "" {
		return ErrInvalidHotelName
	}
	return nil
}

func validateAddress(address string) error {
	if address == "" {
		return ErrInvalidHotelAddress
	}
	return nil
}

func validateRating(rating uint) error {
	if rating <= 0 {
		return ErrInvalidHotelRating
	}
	return nil
}

func validateOwnerID(ownerID userDomain.UserID) error {
	if ownerID == uuid.Nil {
		return ErrInvalidHotelOwnerID
	}
	return nil
}

func (h *Hotel) Validate() error {
	if err := ValidateID(h.ID); err != nil {
		return err
	}

	if err := validateName(h.Name); err != nil {
		return err
	}

	if err := validateAddress(h.Address); err != nil {
		return err
	}

	if err := validateRating(h.Rating); err != nil {
		return err
	}

	if err := validateOwnerID(h.OwnerID); err != nil {
		return err
	}

	return nil
}
