package domain

import (
	"time"

	"github.com/google/uuid"
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

type RateFilterItem struct{}
