package common

import "time"

type LogID uint64
type Log struct {
	ID        LogID
	UserID    int64
	CompanyID int64
	Action    string
	Path      string
	CreatedAt time.Time
}
