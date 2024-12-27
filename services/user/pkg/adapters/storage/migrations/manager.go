package migrations

import (
	"fmt"
	"user_service/pkg/adapters/storage/types"

	"gorm.io/gorm"
)

type Manager struct {
	db *gorm.DB
}

func NewManager(db *gorm.DB) *Manager {
	return &Manager{db: db}
}

func (m *Manager) RunMigrations() error {
	if err := m.db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Error; err != nil {
		return fmt.Errorf("failed to create uuid extension: %w", err)
	}

	if err := m.db.AutoMigrate(&types.User{}); err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	return nil
}
