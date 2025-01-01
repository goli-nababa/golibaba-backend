package types

import "time"

type Room struct {
	ID         string      `gorm:"type:uuid;primary_key"`
	HotelID    string      `gorm:"type:uuid;not null"`
	RateID     string      `gorm:"type:uuid;not null"`
	RoomNumber string      `gorm:"not null"`
	Capacity   uint        `gorm:"not null"`
	Features   StringArray `gorm:"type:json"`
	IsAvilabe  bool        `gorm:"default:true"`
	CreateAt   time.Time
	UpdateAt   time.Time
	DeletedAt  time.Time
}
