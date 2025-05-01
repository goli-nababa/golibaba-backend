package types

import "time"

type Rate struct {
	ID        string
	Name      string
	Price     float64
	Currency  uint8
	StartDate time.Time
	EndDate   time.Time
	CreateAt  time.Time
	UpdateAt  time.Time
	DeletedAt time.Time
}
