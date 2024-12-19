package user_service_client

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/goli-nababa/golibaba-backend/proto/pb"
)

type UserServiceClient struct {
	client  pb.UserServiceClient
	url     string
	port    uint64
	version uint32
}

func NewUserServiceClient(url string, version uint32, port uint64) (*UserServiceClient, error) {
	grpcClient, err := grpc.NewClient(fmt.Sprintf(":%v", port), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return &UserServiceClient{}, err
	}

	return &UserServiceClient{
		url:     url,
		version: version,
		port:    port,
		client:  pb.NewUserServiceClient(grpcClient),
	}, nil
}

var (
	ErrUserExists = errors.New("user already exists")
)

func (us *UserServiceClient) CreateUser(user *pb.User) (*pb.CreateUserResponse, error) {
	response, err := us.client.CreateUser(context.Background(), user)

	if err != nil {
		return nil, err
	}

	if response.Status == 409 {
		return nil, ErrUserExists
	}

	return response, nil
}
