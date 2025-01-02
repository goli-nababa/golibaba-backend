package storage

import (
	"context"
	"errors"
	"fmt"
	"user_service/pkg/adapters/storage/mapper"

	storageTypes "user_service/pkg/adapters/storage/types"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/goli-nababa/golibaba-backend/common"
	"github.com/google/uuid"
	userService "user_service/internal/user"
	userPort "user_service/internal/user/port"
	"user_service/pkg/adapters/storage/migrations"
	"user_service/pkg/adapters/storage/types"
	"user_service/pkg/hash"

	"gorm.io/gorm"
)

type userRepo struct {
	db     *gorm.DB
	secret string
}

func NewUserRepo(db *gorm.DB, secret string) userPort.Repo {
	return &userRepo{db, secret}
}

func (r *userRepo) FindByUsernamePassword(ctx context.Context, username string, password string) (*common.User, error) {
	var user types.User

	// Retrieve the user by username
	err := r.db.WithContext(ctx).Where("email = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}

	// Validate the plain password against the hashed password
	bcryptHasher := hash.NewBcryptHasher()
	if !bcryptHasher.Validate(user.Password, password) {
		return nil, userService.ErrInvalidPassword
	}

	// Map the user to domain and return
	return mapper.UserFromStorage(&user), nil
}

func (r *userRepo) Create(ctx context.Context, user *common.User) error {
	storageUser := mapper.UserToStorage(user)

	result := r.db.WithContext(ctx).Create(storageUser)
	if result.Error != nil {
		return fmt.Errorf("failed to create user: %w", result.Error)
	}
	*user = *mapper.UserFromStorage(storageUser)
	return nil
}

func (r *userRepo) GetByID(ctx context.Context, id common.UserID) (*common.User, error) {
	var user storageTypes.User

	result := r.db.WithContext(ctx).First(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get user by id: %w", result.Error)
	}

	return mapper.UserFromStorage(&user), nil
}

func (r *userRepo) GetByUUID(ctx context.Context, uuid uuid.UUID) (*common.User, error) {
	var user storageTypes.User

	result := r.db.WithContext(ctx).Where("uuid = ?", uuid).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get user by uuid: %w", result.Error)
	}

	return mapper.UserFromStorage(&user), nil
}

func (r *userRepo) DeleteByID(ctx context.Context, id common.UserID) error {
	result := r.db.WithContext(ctx).Delete(&storageTypes.User{}, id)
	if result.Error != nil {
		return fmt.Errorf("failed to delete user by id: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("user not found: %d", id)
	}

	return nil
}

func (r *userRepo) DeleteByUUID(ctx context.Context, uuid uuid.UUID) error {
	result := r.db.WithContext(ctx).Delete(&storageTypes.User{}, uuid)
	if result.Error != nil {
		return fmt.Errorf("failed to delete user by uuid: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("user not found: %d", uuid)
	}

	return nil
}

func (r *userRepo) Block(ctx context.Context, userId uint) error {
	query := r.db.WithContext(ctx).WithContext(ctx).Model(&storageTypes.User{}).
		Where("id = ?", userId).Update("blocked", true)

	if query.Error != nil {
		return fmt.Errorf("failed to delete location: %w", query.Error)
	}

	return nil
}

func (r *userRepo) Unblock(ctx context.Context, userId uint) error {
	query := r.db.WithContext(ctx).WithContext(ctx).Model(&storageTypes.User{}).
		Where("id = ?", userId).Update("blocked", false)

	if query.Error != nil {
		return fmt.Errorf("failed to delete location: %w", query.Error)
	}

	return nil
}

func (r *userRepo) AssignRole(ctx context.Context, userId common.UserID, role string) error {
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

func (r *userRepo) CancelRole(ctx context.Context, userID common.UserID, role string) error {
	result := r.db.WithContext(ctx).Where("user_id = ? AND role = ?", userID, role).Delete(&storageTypes.UserRole{})
	if result.Error != nil {
		return fmt.Errorf("failed to cancel role: %w", result.Error)
	}
	return nil
}

func (r *userRepo) RunMigrations() error {
	migrator := gormigrate.New(r.db, gormigrate.DefaultOptions, migrations.GetUserMigrations())
	return migrator.Migrate()
}

func (r *userRepo) FindByEmail(ctx context.Context, email string) (*common.User, error) {
	var user types.User

	// Retrieve the user by username
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return mapper.UserFromStorage(&user), nil
}

func (r *userRepo) FindByID(ctx context.Context, id common.UserID) (*common.User, error) {
	var user common.User
	if err := r.db.WithContext(ctx).First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, fmt.Errorf("failed to find user by ID: %w", err)
	}
	return &user, nil
}

func (r *userRepo) Update(ctx context.Context, user *common.User) error {
	if err := r.db.WithContext(ctx).Save(user).Error; err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

func (r *userRepo) AssignPermissionToRole(ctx context.Context, userID common.UserID, role string, permissions []string) error {
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

func (r *userRepo) RevokePermissionFromRole(ctx context.Context, userID common.UserID, role string, permissions []string) error {
	result := r.db.WithContext(ctx).Where("user_id = ? AND role = ? AND permission IN ?", userID, role, permissions).Delete(&storageTypes.RolePermission{})
	if result.Error != nil {
		return fmt.Errorf("failed to revoke permissions from role: %w", result.Error)
	}
	return nil
}

func (r *userRepo) PublishStatement(ctx context.Context, userIDs []common.UserID, action common.TypeStatementAction, permissions []string) error {
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

func (r *userRepo) CancelStatement(ctx context.Context, userID common.UserID, statementID common.StatementID) error {
	result := r.db.WithContext(ctx).
		Where("user_id = ? AND id = ?", userID, uint64(statementID)).
		Delete(&storageTypes.Statement{})
	if result.Error != nil {
		return fmt.Errorf("failed to cancel statements: %w", result.Error)
	}

	return nil
}

func (r *userRepo) CheckAccess(ctx context.Context, userID common.UserID, permissions []string) (bool, error) {
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

func (r *userRepo) Insert(ctx context.Context, user *common.User) (common.UserID, error) {
	newU := mapper.UserToStorage(user)

	return common.UserID(newU.ID), r.db.WithContext(ctx).Create(newU).Error
}
