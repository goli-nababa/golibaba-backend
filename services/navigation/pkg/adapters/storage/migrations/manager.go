package migrations

import (
	"fmt"
	"gorm.io/gorm"
	"navigation_service/pkg/adapters/storage/types"
)

type Manager struct {
	db *gorm.DB
}

func NewManager(db *gorm.DB) *Manager {
	return &Manager{db: db}
}

func (m *Manager) RunMigrations() error {
	// Enable UUID and JSONB support for PostgreSQL
	if err := m.db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Error; err != nil {
		return fmt.Errorf("failed to create uuid extension: %w", err)
	}

	// Run migrations
	if err := m.db.AutoMigrate(&types.Location{}, &types.Route{}); err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	// Create indexes
	return m.createIndexes()
}

func (m *Manager) createIndexes() error {

	if err := m.db.Exec(`
			DO $$
			BEGIN
			   IF NOT EXISTS (SELECT 1 FROM pg_extension WHERE extname = 'cube') THEN
				  CREATE EXTENSION cube;
			   END IF;
			   IF NOT EXISTS (SELECT 1 FROM pg_extension WHERE extname = 'earthdistance') THEN
				  CREATE EXTENSION earthdistance;
			   END IF;
			END $$;
		`).Error; err != nil {
		return err
	}

	// Location indexes
	if err := m.db.Exec(`
        CREATE INDEX IF NOT EXISTS idx_locations_type ON locations(type);
        CREATE INDEX IF NOT EXISTS idx_locations_active ON locations(active) WHERE active = true;
        CREATE INDEX IF NOT EXISTS idx_locations_coords ON locations USING gist (
            ll_to_earth(latitude, longitude)
        );
    `).Error; err != nil {
		return err
	}

	// Route indexes
	if err := m.db.Exec(`
        CREATE INDEX IF NOT EXISTS idx_routes_active ON routings(active) WHERE active = true;
        CREATE INDEX IF NOT EXISTS idx_routes_locations ON routings(from_id, to_id);
    `).Error; err != nil {
		return err
	}

	return nil
}
