package domain

import "time"

type (
	AdminId uint
)

type Admin struct {
	ID          AdminId
	CreatedAt   time.Time
	DeletedAt   time.Time
	SubmittedAt time.Time
	Name        string
}
