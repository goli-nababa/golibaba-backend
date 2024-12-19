package port

import (
	"admin/internal/admin/domain"
	"context"
)

type Repo interface {
	Create(ctx context.Context, order domain.Admin) (*domain.Admin, error)
}
