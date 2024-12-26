package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidID        = errors.New("rate id is required")
	ErrInvalidName      = errors.New("name cannot be empty")
	ErrInvalidPrice     = errors.New("price must be greater than zero")
	ErrInvalidCurrency  = errors.New("invalid currency type")
	ErrInvalidStartDate = errors.New("start date cannot be empty")
	ErrInvalidEndDate   = errors.New("end date cannot be empty")
	ErrInvalidDateRange = errors.New("end date must be after start date")
)

type CurrencyType uint8

const (
	CurrencyTypeUnknown CurrencyType = iota
	CurrencyTypeIRR
	CurrencyTypeUSD
)

type RateID = uuid.UUID
type Rate struct {
	ID        RateID
	Name      string
	Price     float64
	Currency  CurrencyType
	CreateAt  time.Time
	DeletedAt time.Time
}

type RateFilterItem struct {
	Name     string
	Price    float64
	Currency CurrencyType
}

func ValidID(ID RateID) error {
	if err := uuid.Validate(ID.String()); err != nil {
		return ErrInvalidID
	}
	return nil
}

func validName(name string) error {
	if name == "" {
		return ErrInvalidName
	}
	return nil
}

func validPrice(price float64) error {
	if price <= 0 {
		return ErrInvalidPrice
	}
	return nil
}

func validCurrency(currency CurrencyType) error {
	if currency == CurrencyTypeUnknown {
		return ErrInvalidCurrency
	}
	return nil
}

func validStartDate(startDate time.Time) error {
	if startDate.IsZero() {
		return ErrInvalidStartDate
	}
	return nil
}

func validEndDate(endDate time.Time) error {
	if endDate.IsZero() {
		return ErrInvalidEndDate
	}
	return nil
}

func validDateRange(startDate, endDate time.Time) error {
	if endDate.Before(startDate) {
		return ErrInvalidDateRange
	}
	return nil
}

func (r *Rate) Validate() error {
	if err := ValidID(r.ID); err != nil {
		return err
	}

	if err := validName(r.Name); err != nil {
		return err
	}

	if err := validPrice(r.Price); err != nil {
		return err
	}

	if err := validCurrency(r.Currency); err != nil {
		return err
	}

	return nil
}
