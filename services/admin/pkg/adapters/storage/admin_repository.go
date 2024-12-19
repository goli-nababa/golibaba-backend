package storage

import (
	"admin/internal/admin/domain"
	"admin/internal/admin/port"
	"admin/pkg/adapters/storage/mapper"
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
	admin := mapper.AdminDomain2Storage(repoDomain)
	return mapper.AdminStorage2Domain(*admin), r.db.Table("users").WithContext(ctx).Create(admin).Error
}
