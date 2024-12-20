package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBConnOptions struct {
	User   string
	Pass   string
	Host   string
	Port   uint
	Name   string
	Schema string
}

func (o DBConnOptions) PostgresDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s search_path=%s sslmode=disable",
		o.Host, o.Port, o.User, o.Pass, o.Name, o.Schema)
}

func NewPsqlGormConnection(opt DBConnOptions) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(opt.PostgresDSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return db, err
	}

	return db, nil
}
