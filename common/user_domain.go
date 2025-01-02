package common

import (
	"github.com/google/uuid"
	"time"
)

type UserID uint64

type User struct {
	ID        UserID
	UUID      uuid.UUID
	FirstName string
	LastName  string
	Email     string
	Password  string
	Phone     string
	Blocked   bool
	WalletID  WalletID
	Role      string
	Birthday  time.Time
	CreatedAt time.Time
	DeletedAt time.Time
}
