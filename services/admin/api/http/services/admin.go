package services

import (
	adminPort "admin/internal/port"
	"context"

	"github.com/goli-nababa/golibaba-backend/common"
)

type AdminService struct {
	svc adminPort.Service
}

func NewAdminService(svc adminPort.Service) *AdminService {
	return &AdminService{
		svc: svc,
	}
}

func (s *AdminService) BlockEntity(ctx context.Context, entityID string, entityType string) error {
	return s.svc.BlockEntity(ctx, entityID, entityType)
}

func (s *AdminService) UnblockEntity(ctx context.Context, entityID string, entityType string) error {
	return s.svc.UnblockEntity(ctx, entityID, entityType)
}

func (s *AdminService) AssignRole(ctx context.Context, userID uint, role string) error {
	return s.svc.AssignRole(ctx, userID, role)
}
func (s *AdminService) CancelRole(ctx context.Context, userID uint, role string) error {
	return s.svc.CancelRole(ctx, userID, role)
}
func (s *AdminService) AssignPermissionToRole(ctx context.Context, userID uint, role string, permissions []string) error {
	return s.svc.AssignPermissionToRole(ctx, userID, role, permissions)
}
func (s *AdminService) RevokePermissionFromRole(ctx context.Context, userID uint, role string, permissions []string) error {
	return s.svc.RevokePermissionFromRole(ctx, userID, role, permissions)
}
func (s *AdminService) PublishStatement(ctx context.Context, userIDs []common.UserID, action string, permissions []string) error {
	return s.svc.PublishStatement(ctx, userIDs, action, permissions)
}
func (s *AdminService) CancelStatement(ctx context.Context, userID uint, statementID uint) error {
	return s.svc.CancelStatement(ctx, userID, statementID)
}
