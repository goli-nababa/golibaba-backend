package user

import (
	"context"
	"fmt"
	"user_service/internal/port"

	"github.com/google/uuid"

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

func (s *service) GetUserByID(ctx context.Context, userID common.UserID) (*common.User, error) {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}
	return user, nil
}

func (s *service) GetUserByUUID(ctx context.Context, userUUID uuid.UUID) (*common.User, error) {
	user, err := s.userRepo.GetByUUID(ctx, userUUID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by uuid: %w", err)
	}
	return user, nil
}

func (s *service) DeleteUserByID(ctx context.Context, userID common.UserID) error {
	if err := s.userRepo.DeleteByID(ctx, userID); err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}

func (s *service) DeleteUserByUUID(ctx context.Context, userUUID uuid.UUID) error {
	if err := s.userRepo.DeleteByUUID(ctx, userUUID); err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}

func (s *service) BlockUser(ctx context.Context, userID uint) error {
	if err := s.userRepo.Block(ctx, userID); err != nil {
		return fmt.Errorf("failed to block user: %w", err)
	}
	return nil
}

func (s *service) UnblockUser(ctx context.Context, userID uint) error {
	if err := s.userRepo.Unblock(ctx, userID); err != nil {
		return fmt.Errorf("failed to unblock user: %w", err)
	}
	return nil
}

func (s *service) AssignRole(ctx context.Context, userID common.UserID, role string) error {
	if err := s.userRepo.AssignRole(ctx, userID, role); err != nil {
		return fmt.Errorf("failed to assign role to user: %w", err)
	}
	return nil
}

func (s *service) CancelRole(ctx context.Context, userID common.UserID, role string) error {
	if err := s.userRepo.CancelRole(ctx, userID, role); err != nil {
		return fmt.Errorf("failed to cancel role from user: %w", err)
	}
	return nil
}

func (s *service) AssignPermissionToRole(ctx context.Context, userID common.UserID, role string, permissions []string) error {
	if err := s.userRepo.AssignPermissionToRole(ctx, userID, role, permissions); err != nil {
		return fmt.Errorf("failed to assign permissions to role for user: %w", err)
	}
	return nil
}

func (s *service) RevokePermissionFromRole(ctx context.Context, userID common.UserID, role string, permissions []string) error {
	if err := s.userRepo.RevokePermissionFromRole(ctx, userID, role, permissions); err != nil {
		return fmt.Errorf("failed to revoke permissions from role for user: %w", err)
	}
	return nil
}

func (s *service) PublishStatement(ctx context.Context, userIDs []common.UserID, action common.TypeStatementAction, permissions []string) error {
	if err := s.userRepo.PublishStatement(ctx, userIDs, action, permissions); err != nil {
		return fmt.Errorf("failed to publish statement for users: %w", err)
	}
	return nil
}
func (s *service) CancelStatement(ctx context.Context, userIDs common.UserID, statementID common.StatementID) error {
	if err := s.userRepo.CancelStatement(ctx, userIDs, statementID); err != nil {
		return fmt.Errorf("failed to cancel statement for user: %w", err)
	}
	return nil
}
func (s *service) CheckAccess(ctx context.Context, userID common.UserID, permissions []string) (bool, error) {
	hasAccess, err := s.userRepo.CheckAccess(ctx, userID, permissions)

	if err != nil {
		return false, fmt.Errorf("failed to cancel statement for user: %w", err)
	}

	return hasAccess, nil
}
