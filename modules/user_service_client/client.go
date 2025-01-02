package user_service_client

import (
	"context"
	"errors"
	"fmt"

	"github.com/goli-nababa/golibaba-backend/common"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/goli-nababa/golibaba-backend/proto/pb"
)

type userServiceClient struct {
	client  pb.UserServiceClient
	url     string
	port    uint64
	version uint32
}

func NewUserServiceClient(url string, version uint32, port uint64) (UserServiceClient, error) {
	grpcClient, err := grpc.NewClient(fmt.Sprintf(":%v", port), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return &userServiceClient{}, err
	}

	return &userServiceClient{
		url:     url,
		version: version,
		port:    port,
		client:  pb.NewUserServiceClient(grpcClient),
	}, nil
}

var (
	ErrUserExists = errors.New("user already exists")
)

func (us *userServiceClient) CreateUser(user *pb.User) (*pb.CreateUserResponse, error) {
	response, err := us.client.CreateUser(context.Background(), user)

	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return response, nil
}

func (us *userServiceClient) BlockUser(userID common.UserID) error {
	_, err := us.client.BlockUser(context.Background(), &pb.BlockUserRequest{UserId: uint64(userID)})
	if err != nil {
		return fmt.Errorf("failed to block user: %w", err)
	}

	return nil
}

func (us *userServiceClient) UnblockUser(userID common.UserID) error {
	_, err := us.client.UnblockUser(context.Background(), &pb.UnblockUserRequest{UserId: uint64(userID)})
	if err != nil {
		return fmt.Errorf("failed to unblock user: %w", err)
	}

	return nil
}

func (us *userServiceClient) GetUserByID(userID common.UserID) (*pb.User, error) {
	resp, err := us.client.GetUserByID(context.Background(), &pb.GetUserByIDRequest{UserId: uint64(userID)})
	if err != nil {
		return nil, fmt.Errorf("failed to get user by ID: %w", err)
	}

	return resp.User, nil
}

func (us *userServiceClient) GetUserByUUID(userUUID uuid.UUID) (*pb.User, error) {
	resp, err := us.client.GetUserByUUID(context.Background(), &pb.GetUserByUUIDRequest{Uuid: userUUID.String()})
	if err != nil {
		return nil, fmt.Errorf("failed to get user by UUID: %w", err)
	}

	return resp.User, nil
}

func (us *userServiceClient) DeleteUserByID(userID common.UserID) error {
	_, err := us.client.DeleteUserByID(context.Background(), &pb.DeleteUserByIDRequest{UserId: uint64(userID)})
	if err != nil {
		return fmt.Errorf("failed to delete user by ID: %w", err)
	}

	return nil
}

func (us *userServiceClient) DeleteUserByUUID(userUUID uuid.UUID) error {
	_, err := us.client.DeleteUserByUUID(context.Background(), &pb.DeleteUserByUUIDRequest{Uuid: userUUID.String()})
	if err != nil {
		return fmt.Errorf("failed to delete user by ID: %w", err)
	}

	return nil
}

func (us *userServiceClient) AssignRole(userID common.UserID, role string) error {
	_, err := us.client.AssignRole(context.Background(), &pb.AssignRoleRequest{UserId: uint64(userID), Role: role})
	if err != nil {
		return fmt.Errorf("failed to assign role: %w", err)
	}

	return nil
}

func (us *userServiceClient) CancelRole(userID common.UserID, role string) error {
	_, err := us.client.CancelRole(context.Background(), &pb.CancelRoleRequest{UserId: uint64(userID), Role: role})
	if err != nil {
		return fmt.Errorf("failed to cancel role: %w", err)
	}

	return nil
}

func (us *userServiceClient) AssignPermissionToRole(userID common.UserID, role string, permissions []string) error {
	_, err := us.client.AssignPermission(context.Background(), &pb.AssignPermissionRequest{
		UserId:      uint64(userID),
		Role:        role,
		Permissions: permissions,
	})
	if err != nil {
		return fmt.Errorf("failed to assign permissions to role: %w", err)
	}

	return nil
}

func (us *userServiceClient) RevokePermissionFromRole(userID common.UserID, role string, permissions []string) error {
	_, err := us.client.RevokePermission(context.Background(), &pb.RevokePermissionRequest{
		UserId:      uint64(userID),
		Role:        role,
		Permissions: permissions,
	})
	if err != nil {
		return fmt.Errorf("failed to revoke permissions from role: %w", err)
	}

	return nil
}

func (us *userServiceClient) PublishStatement(userIDs []common.UserID, action common.TypeStatementAction, permissions []string) error {
	var ids []uint64
	for _, id := range userIDs {
		ids = append(ids, uint64(id))
	}

	_, err := us.client.PublishStatement(context.Background(), &pb.PublishStatementRequest{
		UserIds:     ids,
		Action:      pb.TypeStatementAction(action),
		Permissions: permissions,
	})
	if err != nil {
		return fmt.Errorf("failed to publish statement: %w", err)
	}

	return nil
}

func (us *userServiceClient) CancelStatement(userID common.UserID, statementID common.StatementID) error {
	_, err := us.client.CancelStatement(context.Background(), &pb.CancelStatementRequest{
		UserId:      uint64(userID),
		StatementId: uint64(statementID),
	})
	if err != nil {
		return fmt.Errorf("failed to cancel statement: %w", err)
	}

	return nil
}

func (us *userServiceClient) CheckAccess(userID common.UserID, permissions []string) (bool, error) {
	resp, err := us.client.CheckAccess(context.Background(), &pb.CheckAccessRequest{
		UserId:      uint64(userID),
		Permissions: permissions,
	})
	if err != nil {
		return false, fmt.Errorf("failed to check access: %w", err)
	}

	return resp.HasAccess, nil
}
