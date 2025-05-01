package middleware

import (
	"bank_service/pkg/transaction"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func TransactionInterceptor(sessions *transaction.SessionStore) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if isSessionMethod(info.FullMethod) {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return handler(ctx, req)
		}

		sessionIDs := md.Get("session_id")
		if len(sessionIDs) == 0 {
			return handler(ctx, req)
		}

		session, exists := sessions.Get(sessionIDs[0])
		if !exists {
			return nil, status.Errorf(codes.NotFound, "session not found")
		}

		txCtx := transaction.ContextWithTransaction(ctx, session.Tx)
		return handler(txCtx, req)
	}
}

func isSessionMethod(method string) bool {
	return method == "/bank.v1.SessionService/BeginTransaction" ||
		method == "/bank.v1.SessionService/CommitTransaction" ||
		method == "/bank.v1.SessionService/RollbackTransaction"
}
