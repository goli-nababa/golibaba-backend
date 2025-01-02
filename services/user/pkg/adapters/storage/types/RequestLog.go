package types

import "time"

type RequestLog struct {
	ID        uint `gorm:"primarykey"`
	UserID    uint `gorm:"type:uuid;default:gen_random_uuid()"`
	CompanyID uint
	Action    string `gorm:"not null"`
	Path      string `gorm:"not null"`
	CreatedAt time.Time
}

func (RequestLog) TableName() string {
	return "request_logs"
}
