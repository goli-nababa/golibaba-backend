package storage

import (
	"context"
	"transportation/internal/company/domain"
	"transportation/internal/company/port"
	"transportation/pkg/adapters/storage/mapper"

	"gorm.io/gorm"
)

type companyRepo struct {
	db *gorm.DB
}

func NewCompanyRepo(db *gorm.DB) port.Repo {
	repo := &companyRepo{db}
	return repo
}

func (r *companyRepo) Create(ctx context.Context, companyDomain domain.Company) (*domain.Company, error) {
	company := mapper.CompanyDomain2Storage(companyDomain)
	return mapper.CompanyStorage2Domain(*company), r.db.Table("companies").WithContext(ctx).Create(company).Error
}
