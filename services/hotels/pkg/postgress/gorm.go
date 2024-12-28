package postgress

import (
	"fmt"

	PostgresDrive "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBOpt struct {
	Host     string
	Port     uint
	User     string
	Password string
	DBName   string
	SSLMode  string
	Schema   string
}

func (db *DBOpt) dsn() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s search_path=%s sslmode=%s",
		db.Host, db.Port, db.User, db.Password, db.DBName, db.Schema, db.SSLMode)
}

func NewConnection(psg DBOpt) (*gorm.DB, error) {
	return gorm.Open(PostgresDrive.Open(psg.dsn()), &gorm.Config{})
}
