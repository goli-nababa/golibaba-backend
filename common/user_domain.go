package common

import "github.com/google/uuid"

type UserID uint64

type User struct {
	ID        UserID
	UUID      uuid.UUID
	FirstName string
	LastName  string
	Email     string
	Password  string
	Phone     string
	WalletID  WalletID
	// ToDo: Implement me
}
