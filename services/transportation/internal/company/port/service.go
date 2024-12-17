package port

import (
	"context"
	"transportation/internal/company/domain"
)

type Service interface {
	CreateCompany(ctx context.Context, company domain.Company) (*domain.Company, error)
}
