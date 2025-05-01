package types

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primarykey"`
	UUID      string `gorm:"type:uuid;default:gen_random_uuid()"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Email     string `gorm:"unique;not null" json:"email"`
	Password  string `gorm:"not null"`
	Phone     string
	Blocked   bool `gorm:"default:false"`
	WalletID  uint `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (User) TableName() string {
	return "users"
}
