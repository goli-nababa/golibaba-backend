package storage

import (
	"context"
	"fmt"
	"user_service/internal/port"
	"user_service/pkg/adapters/storage/mapper"

	storageTypes "user_service/pkg/adapters/storage/types"

	"github.com/goli-nababa/golibaba-backend/common"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *common.User) error {
	return nil
}

func (r *userRepository) Block(ctx context.Context, userId uint) error {
	query := r.db.WithContext(ctx).WithContext(ctx).Model(&storageTypes.User{}).
		Where("id = ?", userId).Update("blocked", true)

	if query.Error != nil {
		return fmt.Errorf("failed to delete location: %w", query.Error)
	}

	return nil
}

func (r *userRepository) Unblock(ctx context.Context, userId uint) error {
	query := r.db.WithContext(ctx).WithContext(ctx).Model(&storageTypes.User{}).
		Where("id = ?", userId).Update("blocked", false)

	if query.Error != nil {
		return fmt.Errorf("failed to delete location: %w", query.Error)
	}

	return nil
}

func (r *userRepository) AssignRole(ctx context.Context, userId common.UserID, role string) error {
	userRole := storageTypes.UserRole{
		UserID: uint(userId),
		Role:   role,
	}
	result := r.db.WithContext(ctx).Create(&userRole)
	if result.Error != nil {
		return fmt.Errorf("failed to assign role: %w", result.Error)
	}
	return nil
}

func (r *userRepository) CancelRole(ctx context.Context, userID common.UserID, role string) error {
	result := r.db.WithContext(ctx).Where("user_id = ? AND role = ?", userID, role).Delete(&storageTypes.UserRole{})
	if result.Error != nil {
		return fmt.Errorf("failed to cancel role: %w", result.Error)
	}
	return nil
}

func (r *userRepository) AssignPermissionToRole(ctx context.Context, userID common.UserID, role string, permissions []string) error {
	var rolePermissions []storageTypes.RolePermission
	for _, perm := range permissions {
		rolePermissions = append(rolePermissions, storageTypes.RolePermission{
			UserID:     uint(userID),
			Role:       role,
			Permission: perm,
		})
	}

	result := r.db.WithContext(ctx).Create(&rolePermissions)
	if result.Error != nil {
		return fmt.Errorf("failed to assign permissions to role: %w", result.Error)
	}
	return nil
}

func (r *userRepository) RevokePermissionFromRole(ctx context.Context, userID common.UserID, role string, permissions []string) error {
	result := r.db.WithContext(ctx).Where("user_id = ? AND role = ? AND permission IN ?", userID, role, permissions).Delete(&storageTypes.RolePermission{})
	if result.Error != nil {
		return fmt.Errorf("failed to revoke permissions from role: %w", result.Error)
	}
	return nil
}

func (r *userRepository) PublishStatement(ctx context.Context, userIDs []common.UserID, action common.TypeStatementAction, permissions []string) error {
	var statements []storageTypes.Statement

	for _, userID := range userIDs {
		statements = append(statements, storageTypes.Statement{
			UserID:      uint(userID),
			Action:      action,
			Permissions: permissions,
		})
	}

	result := r.db.WithContext(ctx).Create(&statements)
	if result.Error != nil {
		return fmt.Errorf("failed to publish statements: %w", result.Error)
	}

	return nil
}

func (r *userRepository) CancelStatement(ctx context.Context, userID common.UserID, statementID common.StatementID) error {
	result := r.db.WithContext(ctx).
		Where("user_id = ? AND id = ?", userID, uint64(statementID)).
		Delete(&storageTypes.Statement{})
	if result.Error != nil {
		return fmt.Errorf("failed to cancel statements: %w", result.Error)
	}

	return nil
}

func (r *userRepository) CheckAccess(ctx context.Context, userID common.UserID, permissions []string) (bool, error) {
	var roleCount int64
	result := r.db.WithContext(ctx).Model(&storageTypes.RolePermission{}).
		Where("user_id = ? AND permission IN ?", userID, permissions).Count(&roleCount)
	if result.Error != nil {
		return false, fmt.Errorf("failed to check role permissions: %w", result.Error)
	}

	if roleCount > 0 {
		return true, nil
	}

	var statementCount int64
	result = r.db.WithContext(ctx).
		Model(&storageTypes.Statement{}).
		Where("user_id = ? AND action = ? AND JSON_CONTAINS(permissions, ?)", userID, common.StatementActionAllow, permissions).
		Count(&statementCount)
	if result.Error != nil {
		return false, fmt.Errorf("failed to check statement permissions: %w", result.Error)
	}

	if statementCount > 0 {
		return true, nil
	}

	return false, nil
}
