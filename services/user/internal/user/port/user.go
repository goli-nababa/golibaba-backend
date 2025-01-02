package user

import (
	"context"
	"user_service/internal/domain"

	"github.com/goli-nababa/golibaba-backend/common"
)

type Repo interface {
	FindByUsernamePassword(ctx context.Context, username string, password string) (*common.User, error)
	FindByEmail(ctx context.Context, email string) (*common.User, error)
	Insert(ctx context.Context, user *common.User) (common.UserID, error)
	RunMigrations() error
	FindByID(ctx context.Context, id common.UserID) (*common.User, error)
	Update(ctx context.Context, user *common.User) error
	
	Create(ctx context.Context, user *common.User) error
	GetByID(ctx context.Context, userID common.UserID) (*common.User, error)
	GetByUUID(ctx context.Context, userUUID uuid.UUID) (*common.User, error)
	DeleteByID(ctx context.Context, userID common.UserID) error
	DeleteByUUID(ctx context.Context, userUUID uuid.UUID) error

	Block(ctx context.Context, userId uint) error
	Unblock(ctx context.Context, userId uint) error
	AssignRole(ctx context.Context, userId common.UserID, role string) error
	CancelRole(ctx context.Context, userID common.UserID, role string) error
	AssignPermissionToRole(ctx context.Context, userID common.UserID, role string, permissions []string) error
	RevokePermissionFromRole(ctx context.Context, userID common.UserID, role string, permissions []string) error
	PublishStatement(ctx context.Context, userIDs []common.UserID, action common.TypeStatementAction, permissions []string) error
	CancelStatement(ctx context.Context, userIDs common.UserID, statementID common.StatementID) error
	CheckAccess(ctx context.Context, userID common.UserID, permissions []string) (bool, error)
	GetLogByUserId(ctx context.Context, userId uint, page int, pageSize int) ([]common.Log, error)
	SaveLog(ctx context.Context, log *common.Log) error
	ListNotif(ctx context.Context, userId uint) ([]domain.Notification, error)
	CreateNotif(ctx context.Context, notification *domain.Notification) error
}
