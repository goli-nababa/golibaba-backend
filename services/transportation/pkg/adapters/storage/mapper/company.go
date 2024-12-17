package mapper

import (
	"transportation/internal/company/domain"
	"transportation/pkg/adapters/storage/types"

	"gorm.io/gorm"
)

func CompanyDomain2Storage(companyDomain domain.Company) *types.Company {
	return &types.Company{
		Model: gorm.Model{
			ID:        uint(companyDomain.ID),
			CreatedAt: companyDomain.CreatedAt,
			DeletedAt: gorm.DeletedAt(ToNullTime(companyDomain.DeletedAt)),
		},
		Name: companyDomain.Name,
	}
}

func CompanyStorage2Domain(company types.Company) *domain.Company {
	return &domain.Company{
		ID:        domain.CompanyId(company.ID),
		CreatedAt: company.CreatedAt,
		DeletedAt: company.DeletedAt.Time,
		Name:      company.Name,
	}
}
