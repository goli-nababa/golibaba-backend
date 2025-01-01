package domain

import (
	"time"
)

type NotificationID uint

type Notification struct {
	ID        NotificationID
	UserID    uint
	Message   string
	Seen      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
