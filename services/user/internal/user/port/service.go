package user

import (
	"context"
	"user_service/internal/domain"
	"github.com/google/uuid"
	"github.com/goli-nababa/golibaba-backend/common"
)

type Service interface {
	GetUserByUsernamePassword(ctx context.Context, username string, password string) (*common.User, error)
	GetUserByEmail(ctx context.Context, email string) (*common.User, error)
	CreateUser(ctx context.Context, user *common.User) (common.UserID, error)
	RunMigrations() error
	GetUserByID(ctx context.Context, id common.UserID) (*common.User, error)
	GetUserByUUID(ctx context.Context, userUUID uuid.UUID) (*common.User, error)
	DeleteUserByID(ctx context.Context, userID common.UserID) error
	DeleteUserByUUID(ctx context.Context, userID uuid.UUID) error

	BlockUser(ctx context.Context, userId uint) error
	UnblockUser(ctx context.Context, userId uint) error
	AssignRole(ctx context.Context, userId common.UserID, role string) error
	CancelRole(ctx context.Context, userID common.UserID, role string) error
	AssignPermissionToRole(ctx context.Context, userID common.UserID, role string, permissions []string) error
	RevokePermissionFromRole(ctx context.Context, userID common.UserID, role string, permissions []string) error
	PublishStatement(ctx context.Context, userIDs []common.UserID, action common.TypeStatementAction, permissions []string) error
	CancelStatement(ctx context.Context, userIDs common.UserID, statementID common.StatementID) error
	CheckAccess(ctx context.Context, userID common.UserID, permissions []string) (bool, error)
	GetHistory(ctx context.Context, id uint, page int, pageSize int) ([]common.Log, error)
	SaveLog(ctx context.Context, log *common.Log) error

	GetNotifications(ctx context.Context, userId uint) ([]domain.Notification, error)
	CreateNotification(ctx context.Context, notification *domain.Notification) error
}
