package port

import (
	"context"

	"github.com/goli-nababa/golibaba-backend/common"
)

type Service interface {
	CreateUser(ctx context.Context, user *common.User) error
}
