package types

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID      string `gorm:"type:uuid;default:gen_random_uuid()"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Email     string `gorm:"unique;not null" json:"email"`
	Password  string `gorm:"not null"`
	Phone     string
	WalletID  uint `gorm:"not null"`
}

func (User) TableName() string {
	return "users"
}
