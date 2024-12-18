package domain

type RoleType uint8

const (
	RoleTypeUnknown RoleType = iota
	RoleTypeAdmin
	RoleTypeUser
	RoleTypeOwner
)

type UserID string
type User struct {
	ID    UserID
	Name  string
	Email string
	Role  string
}
