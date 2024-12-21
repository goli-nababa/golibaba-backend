package typs

import "time"

type Room struct {
	ID         string
	HotelID    string
	RateID     string
	RoomNumber string
	Capacity   uint
	Features   []string
	IsAvilabe  bool
	CreateAt   time.Time
	UpdateAt   time.Time
	DeletedAt  time.Time
}
