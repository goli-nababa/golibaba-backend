package migrations

import (
	"transportation/pkg/adapters/storage/types"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&types.Company{},
		&types.TechnicalTeam{},
		&types.TechnicalTeamMember{},
		&types.TransportationType{},
		&types.Trip{},
		&types.VehicleRequest{},
	)
}
