package migrations

import (
	"vehicles/pkg/adapters/storage/types"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&types.Vehicle{},
	)
}
