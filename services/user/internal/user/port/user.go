package user

import (
	"context"
	"github.com/goli-nababa/golibaba-backend/common"
)

type Repo interface {
	FindByUsernamePassword(ctx context.Context, username string, password string) (*common.User, error)
	FindByEmail(ctx context.Context, email string) (*common.User, error)
	Insert(ctx context.Context, user *common.User) (common.UserID, error)
	RunMigrations() error
	FindByID(ctx context.Context, id common.UserID) (*common.User, error)
	Update(ctx context.Context, user *common.User) error

	Block(ctx context.Context, userId uint) error
	Unblock(ctx context.Context, userId uint) error
	AssignRole(ctx context.Context, userId common.UserID, role string) error
	CancelRole(ctx context.Context, userID common.UserID, role string) error
	AssignPermissionToRole(ctx context.Context, userID common.UserID, role string, permissions []string) error
	RevokePermissionFromRole(ctx context.Context, userID common.UserID, role string, permissions []string) error
	PublishStatement(ctx context.Context, userIDs []common.UserID, action common.TypeStatementAction, permissions []string) error
	CancelStatement(ctx context.Context, userIDs common.UserID, statementID common.StatementID) error
	CheckAccess(ctx context.Context, userID common.UserID, permissions []string) (bool, error)
}
