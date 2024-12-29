package user

import (
	"context"
	"fmt"
	"user_service/internal/port"

	"github.com/goli-nababa/golibaba-backend/common"
)

type service struct {
	userRepo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{
		userRepo: repo,
	}
}

func (s *service) CreateUser(ctx context.Context, user *common.User) error {
	if err := s.userRepo.Create(ctx, user); err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}
