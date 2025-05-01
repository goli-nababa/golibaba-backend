package grpc_server

import (
	"bank_service/app"
	"bank_service/pkg/transaction"
	pb "bank_service/proto/gen/go/bank/v1"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SessionServer struct {
	pb.UnimplementedSessionServiceServer
	app      app.App
	Sessions *transaction.SessionStore
}

func NewSessionServer(app app.App) *SessionServer {
	return &SessionServer{
		app:      app,
		Sessions: transaction.NewSessionStore(),
	}
}

func (s *SessionServer) BeginTransaction(ctx context.Context, req *pb.BeginTransactionRequest) (*pb.BeginTransactionResponse, error) {
	session, err := s.Sessions.Create(s.app.DB())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to begin transaction: %v", err)
	}

	return &pb.BeginTransactionResponse{
		SessionId: session.ID,
	}, nil
}

func (s *SessionServer) CommitTransaction(ctx context.Context, req *pb.CommitTransactionRequest) (*pb.CommitTransactionResponse, error) {
	session, exists := s.Sessions.Get(req.SessionId)
	if !exists {
		return nil, status.Errorf(codes.NotFound, "session not found")
	}

	session.Mu.Lock()
	defer session.Mu.Unlock()

	if err := session.Tx.Commit().Error; err != nil {
		return nil, status.Errorf(codes.Internal, "failed to commit: %v", err)
	}

	s.Sessions.Remove(req.SessionId)

	return &pb.CommitTransactionResponse{
		Success: true,
	}, nil
}

func (s *SessionServer) RollbackTransaction(ctx context.Context, req *pb.RollbackTransactionRequest) (*pb.RollbackTransactionResponse, error) {
	session, exists := s.Sessions.Get(req.SessionId)
	if !exists {
		return nil, status.Errorf(codes.NotFound, "session not found")
	}

	session.Mu.Lock()
	defer session.Mu.Unlock()

	if err := session.Tx.Rollback().Error; err != nil {
		return nil, status.Errorf(codes.Internal, "failed to rollback: %v", err)
	}

	s.Sessions.Remove(req.SessionId)

	return &pb.RollbackTransactionResponse{
		Success: true,
	}, nil
}
