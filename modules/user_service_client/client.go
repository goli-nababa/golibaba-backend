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
		return nil, err
	}

	if response.Status == 409 {
		return nil, ErrUserExists
	}

	return response, nil
}

func (us *userServiceClient) BlockUser(userID common.UserID) error {
	//TODO implement me
	panic("implement me")
}

func (us *userServiceClient) GetUserByID(userID common.UserID) (*pb.User, error) {
	//TODO implement me
	panic("implement me")
}

func (us *userServiceClient) GetUserByUUID(userID uuid.UUID) (*pb.User, error) {
	//TODO implement me
	panic("implement me")
}

func (us *userServiceClient) DeleteUserByID(userID common.UserID) (*pb.User, error) {
	//TODO implement me
	panic("implement me")
}

func (us *userServiceClient) DeleteUserByUUID(userID uuid.UUID) (*pb.User, error) {
	//TODO implement me
	panic("implement me")
}

func (us *userServiceClient) AssignRole(userID common.UserID, role string) error {
	//TODO implement me
	panic("implement me")
}

func (us *userServiceClient) CancelRole(userID common.UserID, role string) error {
	//TODO implement me
	panic("implement me")
}

func (us *userServiceClient) UnblockUser(userID common.UserID) error {
	//TODO implement me
	panic("implement me")
}

func (us *userServiceClient) AssignPermissionToRole(userID common.UserID, role string, permissions []string) error {
	//TODO implement me
	panic("implement me")
}

func (us *userServiceClient) RevokePermissionFromRole(userID common.UserID, role string, permissions []string) error {
	//TODO implement me
	panic("implement me")
}

func (us *userServiceClient) PublishStatement(userID []common.UserID, action common.TypeStatementAction, permissions []string) error {
	//TODO implement me
	panic("implement me")
}

func (us *userServiceClient) CancelStatement(userID common.UserID, statementID common.StatementID) error {
	//TODO implement me
	panic("implement me")
}

func (us *userServiceClient) CheckAccess(userID common.UserID, permission []string) (bool, error) {
	//TODO implement me
	panic("implement me")
}
