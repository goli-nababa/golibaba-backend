package port

import (
	"context"

	"github.com/goli-nababa/golibaba-backend/common"
)

type Repo interface {
	Create(ctx context.Context, user *common.User) error
}
