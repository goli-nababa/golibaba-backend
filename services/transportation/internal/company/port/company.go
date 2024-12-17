package port

import (
	"context"
	"transportation/internal/company/domain"
)

type Repo interface {
	Create(ctx context.Context, order domain.Company) (*domain.Company, error)
}
