package domain

import (
	"time"
)

type CurrencyType uint8

const (
	CurrencyTypeUnknown CurrencyType = iota
	CurrencyTypeIRR
	CurrencyTypeUSD
)

type RateID string
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

func (r *Rate) Validation() {
	//TODO
}
