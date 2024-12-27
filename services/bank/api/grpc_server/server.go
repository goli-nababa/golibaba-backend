package grpc_server

import (
	"bank_service/app"
	pb "bank_service/proto/gen/go/bank/v1"
)

type Server struct {
	App app.App
	pb.UnimplementedWalletServiceServer
	pb.UnimplementedTransactionServiceServer
	pb.UnimplementedPaymentServiceServer
	pb.UnimplementedBusinessServiceServer
	pb.UnimplementedCommissionServiceServer
	pb.UnimplementedAnalyticsServiceServer
	pb.UnimplementedFinancialReportServiceServer
}

func NewServer(app app.App) *Server {
	return &Server{
		App: app,
	}
}
