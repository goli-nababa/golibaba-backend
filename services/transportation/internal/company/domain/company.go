package domain

import "time"

type (
	CompanyId uint
)

type Company struct {
	ID          CompanyId
	CreatedAt   time.Time
	DeletedAt   time.Time
	SubmittedAt time.Time
	Name        string
}
