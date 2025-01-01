package migrations

import (
	"fmt"
	"gorm.io/gorm"
	"hotels-service/pkg/adapters/storage/types"
)

type Manager struct {
	db *gorm.DB
}

func NewManager(db *gorm.DB) *Manager {
	return &Manager{db: db}
}

func (m *Manager) RunMigrations() error {
	// Create UUID extension if not exists
	if err := m.db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Error; err != nil {
		return fmt.Errorf("failed to create uuid extension: %w", err)
	}

	// Create Custom types if needed
	if err := m.db.Exec(`DO $$
		BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'string_array') THEN
				CREATE TYPE string_array AS (values text[]);
			END IF;
		EXCEPTION WHEN duplicate_object THEN
			NULL;
		END$$;`).Error; err != nil {
		return fmt.Errorf("failed to create custom types: %w", err)
	}

	// Auto migrate tables
	if err := m.db.AutoMigrate(
		&types.Hotel{},
		&types.Room{},
		&types.Rate{},
		&types.Booking{},
	); err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	// Create indexes
	if err := m.createIndexes(); err != nil {
		return fmt.Errorf("failed to create indexes: %w", err)
	}

	return nil
}

func (m *Manager) createIndexes() error {
	// Hotel indexes
	if err := m.db.Exec(`CREATE INDEX IF NOT EXISTS idx_hotels_owner_id ON hotels(owner_id)`).Error; err != nil {
		return err
	}

	// Room indexes
	if err := m.db.Exec(`CREATE INDEX IF NOT EXISTS idx_rooms_hotel_id ON rooms(hotel_id)`).Error; err != nil {
		return err
	}
	if err := m.db.Exec(`CREATE INDEX IF NOT EXISTS idx_rooms_rate_id ON rooms(rate_id)`).Error; err != nil {
		return err
	}

	// Booking indexes
	if err := m.db.Exec(`CREATE INDEX IF NOT EXISTS idx_bookings_user_id ON bookings(user_id)`).Error; err != nil {
		return err
	}
	if err := m.db.Exec(`CREATE INDEX IF NOT EXISTS idx_bookings_room_id ON bookings(room_id)`).Error; err != nil {
		return err
	}

	return nil
}
