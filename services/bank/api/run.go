package api

import (
	"bank_service/api/grpc_server"
	"bank_service/api/middleware"
	"bank_service/app"
	"bank_service/config"
	"bank_service/pkg/errors"
	pb "bank_service/proto/gen/go/bank/v1"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func RunGRPCServer(ctx context.Context, cfg config.Config) error {
	application := app.NewMustApp(cfg)

	srv := grpc_server.NewServer(application)
	sessionServer := grpc_server.NewSessionServer(application)

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			middleware.TransactionInterceptor(sessionServer.Sessions),
			errors.UnaryServerInterceptor(),
		),
	)

	pb.RegisterSessionServiceServer(grpcServer, sessionServer)

	registerServices(grpcServer, srv)

	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Server.Port))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	log.Printf("Starting gRPC server on port %d", cfg.Server.Port)
	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}

func registerServices(grpcServer *grpc.Server, srv *grpc_server.Server) {
	pb.RegisterWalletServiceServer(grpcServer, srv)
	pb.RegisterTransactionServiceServer(grpcServer, srv)
	pb.RegisterPaymentServiceServer(grpcServer, srv)
	pb.RegisterBusinessServiceServer(grpcServer, srv)
	pb.RegisterCommissionServiceServer(grpcServer, srv)
	pb.RegisterAnalyticsServiceServer(grpcServer, srv)
	pb.RegisterFinancialReportServiceServer(grpcServer, srv)
}
