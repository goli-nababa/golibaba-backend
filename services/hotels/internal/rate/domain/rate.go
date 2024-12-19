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
	StartDate time.Time
	EndDate   time.Time
	CreateAt  time.Time
	EditedAt  time.Time
	DeletedAt time.Time
}

type RateFilterItem struct {
	Name      string
	Price     float64
	Currency  CurrencyType
	StartDate time.Time
	EndDate   time.Time
}

func ValidID(ID RateID) error {
	if ID == uuid.Nil {
		return ErrInvalidID
	}
	return nil
}

func ValidName(name string) error {
	if name == "" {
		return ErrInvalidName
	}
	return nil
}

func ValidPrice(price float64) error {
	if price <= 0 {
		return ErrInvalidPrice
	}
	return nil
}

func ValidCurrency(currency CurrencyType) error {
	if currency == CurrencyTypeUnknown {
		return ErrInvalidCurrency
	}
	return nil
}

func ValidStartDate(startDate time.Time) error {
	if startDate.IsZero() {
		return ErrInvalidStartDate
	}
	return nil
}

func ValidEndDate(endDate time.Time) error {
	if endDate.IsZero() {
		return ErrInvalidEndDate
	}
	return nil
}

func ValidDateRange(startDate, endDate time.Time) error {
	if endDate.Before(startDate) {
		return ErrInvalidDateRange
	}
	return nil
}

func (r *Rate) Validate() error {
	if err := ValidID(r.ID); err != nil {
		return err
	}

	if err := ValidName(r.Name); err != nil {
		return err
	}

	if err := ValidPrice(r.Price); err != nil {
		return err
	}

	if err := ValidCurrency(r.Currency); err != nil {
		return err
	}

	if err := ValidStartDate(r.StartDate); err != nil {
		return err
	}

	if err := ValidEndDate(r.EndDate); err != nil {
		return err
	}

	if err := ValidDateRange(r.StartDate, r.EndDate); err != nil {
		return err
	}

	return nil
}
