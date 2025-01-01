package grpc

import (
	"context"

	di "user_service/app"

	"github.com/goli-nababa/golibaba-backend/common"
	pb "github.com/goli-nababa/golibaba-backend/proto/pb"
	"github.com/google/uuid"
)

type userServiceGRPCApi struct {
	pb.UnsafeUserServiceServer
	app di.App
}

func NewUserServiceGRPCApi(app di.App) pb.UserServiceServer {
	return &userServiceGRPCApi{
		app: app,
	}
}

func (u *userServiceGRPCApi) CreateUser(ctx context.Context, user *pb.User) (*pb.CreateUserResponse, error) {
	uuid, _ := uuid.Parse(user.Uuid)
	userDomain := &common.User{
		ID:        common.UserID(user.Id),
		UUID:      uuid,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Password:  user.Password,
		Phone:     user.Phone,
	}

	if err := u.app.UserService(ctx).CreateUser(ctx, userDomain); err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{
		User: convertToProtoUser(userDomain),
	}, nil
}

func (u *userServiceGRPCApi) GetUserByID(ctx context.Context, req *pb.GetUserByIDRequest) (*pb.GetUserByIDResponse, error) {
	user, err := u.app.UserService(ctx).GetUserByID(ctx, common.UserID(req.UserId))
	if err != nil {
		return nil, err
	}

	if user == nil {
		return &pb.GetUserByIDResponse{}, nil
	}

	return &pb.GetUserByIDResponse{
		User: convertToProtoUser(user),
	}, nil
}

func (u *userServiceGRPCApi) GetUserByUUID(ctx context.Context, req *pb.GetUserByUUIDRequest) (*pb.GetUserByUUIDResponse, error) {
	uuid, _ := uuid.Parse(req.Uuid)

	user, err := u.app.UserService(ctx).GetUserByUUID(ctx, uuid)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return &pb.GetUserByUUIDResponse{}, nil
	}

	return &pb.GetUserByUUIDResponse{
		User: convertToProtoUser(user),
	}, nil
}

func (u *userServiceGRPCApi) DeleteUserByID(ctx context.Context, req *pb.DeleteUserByIDRequest) (*pb.DeleteUserByIDResponse, error) {
	err := u.app.UserService(ctx).DeleteUserByID(ctx, common.UserID(req.UserId))
	if err != nil {
		return nil, err
	}

	return &pb.DeleteUserByIDResponse{}, nil
}

func (u *userServiceGRPCApi) DeleteUserByUUID(ctx context.Context, req *pb.DeleteUserByUUIDRequest) (*pb.DeleteUserByUUIDResponse, error) {
	uuid, _ := uuid.Parse(req.Uuid)

	err := u.app.UserService(ctx).DeleteUserByUUID(ctx, uuid)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteUserByUUIDResponse{}, nil
}

func (u *userServiceGRPCApi) AssignRole(ctx context.Context, req *pb.AssignRoleRequest) (*pb.AssignRoleResponse, error) {
	err := u.app.UserService(ctx).AssignRole(ctx, common.UserID(req.UserId), req.Role)
	if err != nil {
		return nil, err
	}

	return &pb.AssignRoleResponse{}, nil
}

func (u *userServiceGRPCApi) CancelRole(ctx context.Context, req *pb.CancelRoleRequest) (*pb.CancelRoleResponse, error) {
	err := u.app.UserService(ctx).CancelRole(ctx, common.UserID(req.UserId), req.Role)
	if err != nil {
		return nil, err
	}

	return &pb.CancelRoleResponse{}, nil
}

func (u *userServiceGRPCApi) AssignPermission(ctx context.Context, req *pb.AssignPermissionRequest) (*pb.AssignPermissionResponse, error) {
	err := u.app.UserService(ctx).AssignPermissionToRole(ctx, common.UserID(req.UserId), req.Role, req.Permissions)
	if err != nil {
		return nil, err
	}

	return &pb.AssignPermissionResponse{}, nil
}

func (u *userServiceGRPCApi) RevokePermission(ctx context.Context, req *pb.RevokePermissionRequest) (*pb.RevokePermissionResponse, error) {
	err := u.app.UserService(ctx).RevokePermissionFromRole(ctx, common.UserID(req.UserId), req.Role, req.Permissions)
	if err != nil {
		return nil, err
	}

	return &pb.RevokePermissionResponse{}, nil
}

func (u *userServiceGRPCApi) PublishStatement(ctx context.Context, req *pb.PublishStatementRequest) (*pb.PublishStatementResponse, error) {
	var userIds []common.UserID
	for _, id := range req.UserIds {
		userIds = append(userIds, common.UserID(id))
	}
	err := u.app.UserService(ctx).PublishStatement(ctx, userIds, common.TypeStatementAction(req.Action), req.Permissions)
	if err != nil {
		return nil, err
	}

	return &pb.PublishStatementResponse{}, nil
}

func (u *userServiceGRPCApi) CancelStatement(ctx context.Context, req *pb.CancelStatementRequest) (*pb.CancelStatementResponse, error) {
	err := u.app.UserService(ctx).CancelStatement(ctx, common.UserID(req.UserId), common.StatementID(req.StatementId))
	if err != nil {
		return nil, err
	}

	return &pb.CancelStatementResponse{}, nil
}

func (u *userServiceGRPCApi) CheckAccess(ctx context.Context, req *pb.CheckAccessRequest) (*pb.CheckAccessResponse, error) {
	hasAccess, err := u.app.UserService(ctx).CheckAccess(ctx, common.UserID(req.UserId), req.Permissions)
	if err != nil {
		return nil, err
	}

	return &pb.CheckAccessResponse{HasAccess: hasAccess}, nil
}

func (u *userServiceGRPCApi) BlockUser(ctx context.Context, req *pb.BlockUserRequest) (*pb.BlockUserResponse, error) {
	if err := u.app.UserService(ctx).BlockUser(ctx, uint(req.UserId)); err != nil {
		return nil, err
	}

	return &pb.BlockUserResponse{}, nil
}

func (u *userServiceGRPCApi) UnblockUser(ctx context.Context, req *pb.UnblockUserRequest) (*pb.UnblockUserResponse, error) {
	if err := u.app.UserService(ctx).UnblockUser(ctx, uint(req.UserId)); err != nil {
		return nil, err
	}

	return &pb.UnblockUserResponse{}, nil
}

func convertToProtoUser(user *common.User) *pb.User {
	return &pb.User{
		Id:        uint64(user.ID),
		Uuid:      user.UUID.String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Password:  user.Password,
		Phone:     user.Phone,
	}
}
