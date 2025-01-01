package port

import (
	"context"

	"github.com/goli-nababa/golibaba-backend/common"
	"github.com/google/uuid"
)

type Service interface {
	CreateUser(ctx context.Context, user *common.User) error
	GetUserByID(ctx context.Context, userID common.UserID) (*common.User, error)
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
}
