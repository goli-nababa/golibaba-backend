package user_service_client

import (
	"github.com/goli-nababa/golibaba-backend/common"
	pb "github.com/goli-nababa/golibaba-backend/proto/pb"
	"github.com/google/uuid"
)

type UserServiceClient interface {
	CreateUser(user *pb.User) (*pb.CreateUserResponse, error)
	BlockUser(userID common.UserID) error
	UnblockUser(userID common.UserID) error

	GetUserByID(userID common.UserID) (*pb.User, error)
	GetUserByUUID(userID uuid.UUID) (*pb.User, error)
	DeleteUserByID(userID common.UserID) (*pb.User, error)
	DeleteUserByUUID(userID uuid.UUID) (*pb.User, error)

	AssignRole(userID common.UserID, role string) error
	CancelRole(userID common.UserID, role string) error
	AssignPermissionToRole(userID common.UserID, role string, permissions []string) error
	RevokePermissionFromRole(userID common.UserID, role string, permissions []string) error
	PublishStatement(
		userID []common.UserID,
		action common.TypeStatementAction,
		permissions []string,
	) error
	CancelStatement(userID common.UserID, statementID common.StatementID) error
	CheckAccess(userID common.UserID, permission []string) (bool, error)
}
