package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/goli-nababa/golibaba-backend/common"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
	"log"
	port "user_service/internal/user/port"
)

var (
	ErrUserOnCreate      = errors.New("error on creating new user")
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrInvalidPassword   = errors.New("password is invalid")
	ErrPasswordTooLong   = errors.New("password too long")
)

type service struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetUserByUsernamePassword(ctx context.Context, username string, password string) (*common.User, error) {
	user, err := s.repo.FindByUsernamePassword(ctx, username, password)

	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, ErrInvalidPassword):
			return nil, ErrUserNotFound
		default:
			return nil, errors.New(fmt.Sprintf("failed to authenticate user: %s", err))
		}
	}

	return user, nil
}

func (s *service) RunMigrations() error {
	return s.repo.RunMigrations()
}

func (s *service) GetUserByEmail(ctx context.Context, email string) (*common.User, error) {
	user, err := s.repo.FindByEmail(ctx, email)

	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, ErrInvalidPassword):
			return nil, ErrUserNotFound
		default:
			return nil, errors.New(fmt.Sprintf("failed to authenticate user: %s", err))
		}
	}

	return user, nil
}

func (s *service) CreateUser(ctx context.Context, user *common.User) (common.UserID, error) {
	userID, err := s.repo.Insert(ctx, user)

	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return 0, ErrUserAlreadyExists
		}

		log.Println("error on creating new user : ", err.Error())

		return 0, ErrUserOnCreate
	}

	return userID, nil
}

func (s *service) GetUserByID(ctx context.Context, id common.UserID) (*common.User, error) {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found: %w", err)
		}
		return nil, fmt.Errorf("error retrieving user: %w", err)
	}
	return user, nil
}

func (s *service) BlockUser(ctx context.Context, userID uint) error {
	if err := s.repo.Block(ctx, userID); err != nil {
		return fmt.Errorf("failed to block user: %w", err)
	}
	return nil
}

func (s *service) UnblockUser(ctx context.Context, userID uint) error {
	if err := s.repo.Unblock(ctx, userID); err != nil {
		return fmt.Errorf("failed to unblock user: %w", err)
	}
	return nil
}

func (s *service) AssignRole(ctx context.Context, userID common.UserID, role string) error {
	if err := s.repo.AssignRole(ctx, userID, role); err != nil {
		return fmt.Errorf("failed to assign role to user: %w", err)
	}
	return nil
}

func (s *service) CancelRole(ctx context.Context, userID common.UserID, role string) error {
	if err := s.repo.CancelRole(ctx, userID, role); err != nil {
		return fmt.Errorf("failed to cancel role from user: %w", err)
	}
	return nil
}

func (s *service) AssignPermissionToRole(ctx context.Context, userID common.UserID, role string, permissions []string) error {
	if err := s.repo.AssignPermissionToRole(ctx, userID, role, permissions); err != nil {
		return fmt.Errorf("failed to assign permissions to role for user: %w", err)
	}
	return nil
}

func (s *service) RevokePermissionFromRole(ctx context.Context, userID common.UserID, role string, permissions []string) error {
	if err := s.repo.RevokePermissionFromRole(ctx, userID, role, permissions); err != nil {
		return fmt.Errorf("failed to revoke permissions from role for user: %w", err)
	}
	return nil
}

func (s *service) PublishStatement(ctx context.Context, userIDs []common.UserID, action common.TypeStatementAction, permissions []string) error {
	if err := s.repo.PublishStatement(ctx, userIDs, action, permissions); err != nil {
		return fmt.Errorf("failed to publish statement for users: %w", err)
	}
	return nil
}

func (s *service) CancelStatement(ctx context.Context, userIDs common.UserID, statementID common.StatementID) error {
	if err := s.repo.CancelStatement(ctx, userIDs, statementID); err != nil {
		return fmt.Errorf("failed to cancel statement for user: %w", err)
	}
	return nil
}

func (s *service) CheckAccess(ctx context.Context, userID common.UserID, permissions []string) (bool, error) {
	hasAccess, err := s.repo.CheckAccess(ctx, userID, permissions)

	if err != nil {
		return false, fmt.Errorf("failed to cancel statement for user: %w", err)
	}

	return hasAccess, nil
}

func (s *service) GetNotifications(ctx context.Context, userId uint) ([]domain.Notification, error) {
	locations, err := s.repo.ListNotif(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to list notifications: %w", err)
	}
	return locations, nil
}

func (s *service) CreateNotification(ctx context.Context, notification *domain.Notification) error {
	if err := s.repo.CreateNotif(ctx, notification); err != nil {
		return fmt.Errorf("failed to create notification: %w", err)
	}
	return nil
}
