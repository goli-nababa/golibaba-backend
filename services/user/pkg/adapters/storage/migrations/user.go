package migrations

import (
	"gorm.io/gorm"
	"user_service/pkg/adapters/storage/helpers"
	"user_service/pkg/adapters/storage/types"

	"github.com/go-gormigrate/gormigrate/v2"
)

func GetUserMigrations() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		{
			ID: helpers.GenerateMigrationID("add_users_table"),
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&types.User{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("users")
			},
		},
	}
}
