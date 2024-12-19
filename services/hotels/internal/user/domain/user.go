package domain

import "github.com/google/uuid"

type RoleType uint8

const (
	RoleTypeUnknown RoleType = iota
	RoleTypeAdmin
	RoleTypeUser
	RoleTypeOwner
)

type UserID = uuid.UUID
type User struct {
	ID    UserID
	Name  string
	Email string
	Role  string
}
