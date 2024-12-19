package port

import (
	"context"
	"hotels-service/internal/user/domain"
)

type Hotel interface {
	getUserByID(ctx context.Context, userID domain.UserID) (domain.User, error)
}
