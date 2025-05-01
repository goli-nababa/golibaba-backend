package storage

import (
	"admin/internal/domain"
	"admin/internal/port"
	"context"

	"gorm.io/gorm"
)

type adminRepo struct {
	db *gorm.DB
}

func NewAdminRepo(db *gorm.DB) port.Repo {
	repo := &adminRepo{db}
	return repo
}

func (r *adminRepo) Create(ctx context.Context, repoDomain domain.Admin) (*domain.Admin, error) {
	return &repoDomain, nil
}
