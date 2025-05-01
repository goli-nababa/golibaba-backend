package port

import (
	"context"

	"github.com/goli-nababa/golibaba-backend/common"
)

type Service interface {
	BlockEntity(ctx context.Context, entityID string, entityType string) error
	UnblockEntity(ctx context.Context, entityID string, entityType string) error
	AssignRole(ctx context.Context, userID uint, role string) error
	CancelRole(ctx context.Context, userID uint, role string) error
	AssignPermissionToRole(ctx context.Context, userID uint, role string, permissions []string) error
	RevokePermissionFromRole(ctx context.Context, userID uint, role string, permissions []string) error
	PublishStatement(ctx context.Context, userIDs []common.UserID, action string, permissions []string) error
	CancelStatement(ctx context.Context, userID uint, statementID uint) error
}
