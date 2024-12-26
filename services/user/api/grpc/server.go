package grpc

import (
	"context"
	"user_service/config"

	pb "github.com/goli-nababa/golibaba-backend/proto/pb"
	di "user_service/app"
)

type userServiceGRPCApi struct {
	pb.UnsafeUserServiceServer
	appContainer di.App
	config       config.Config
}

func NewUserServiceGRPCApi(appContainer di.App, cfg config.Config) pb.UserServiceServer {
	return &userServiceGRPCApi{
		appContainer: appContainer,
		config:       cfg,
	}
}

func (u *userServiceGRPCApi) CreateUser(ctx context.Context, user *pb.User) (*pb.CreateUserResponse, error) {
	return &pb.CreateUserResponse{
		Status:  200,
		Message: "Created successfully",
	}, nil
}
