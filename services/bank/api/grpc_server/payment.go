package grpc_server

import (
	"bank_service/internal/common/types"
	txDomain "bank_service/internal/services/transaction/domain"
	walletDomain "bank_service/internal/services/wallet/domain"
	pb "bank_service/proto/gen/go/bank/v1"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ProcessPayment(ctx context.Context, req *pb.ProcessPaymentRequest) (*pb.ProcessPaymentResponse, error) {
	tx := &txDomain.Transaction{
		ID:            txDomain.TransactionID(req.Transaction.Id),
		FromWalletID:  walletDomain.WalletID(req.Transaction.FromWalletId),
		ToWalletID:    walletDomain.WalletID(req.Transaction.ToWalletId),
		Amount:        &types.Money{Amount: req.Transaction.Amount.Amount, Currency: req.Transaction.Amount.Currency},
		Type:          txDomain.TransactionType(req.Transaction.Type),
		Status:        txDomain.TransactionStatus(req.Transaction.Status),
		Description:   req.Transaction.Description,
		ReferenceID:   req.Transaction.ReferenceId,
		FailureReason: req.Transaction.FailureReason,
		Version:       int(req.Transaction.Version),
	}

	refID, err := s.App.PaymentService(ctx).ProcessPayment(ctx, tx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to process payment: %v", err)
	}

	return &pb.ProcessPaymentResponse{
		ReferenceId: refID,
	}, nil
}

func (s *Server) ProcessRefund(ctx context.Context, req *pb.ProcessRefundRequest) (*pb.ProcessRefundResponse, error) {
	tx := &txDomain.Transaction{
		FromWalletID: walletDomain.WalletID(req.Transaction.FromWalletId),
		ToWalletID:   walletDomain.WalletID(req.Transaction.ToWalletId),
		Amount:       &types.Money{Amount: req.Transaction.Amount.Amount, Currency: req.Transaction.Amount.Currency},
		Type:         txDomain.TransactionType(req.Transaction.Type),
		Description:  req.Transaction.Description,
		ReferenceID:  req.Transaction.ReferenceId,
		Status:       txDomain.TransactionStatus(req.Transaction.Status),
	}

	err := s.App.PaymentService(ctx).ProcessRefund(ctx, tx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to process refund: %v", err)
	}

	return &pb.ProcessRefundResponse{
		Success: true,
	}, nil
}

func (s *Server) VerifyPayment(ctx context.Context, req *pb.VerifyPaymentRequest) (*pb.VerifyPaymentResponse, error) {
	verified, err := s.App.PaymentService(ctx).VerifyPayment(ctx, req.ReferenceId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to verify payment: %v", err)
	}

	return &pb.VerifyPaymentResponse{
		Verified: verified,
	}, nil
}

func (s *Server) RefundPayment(ctx context.Context, req *pb.RefundPaymentRequest) (*pb.RefundPaymentResponse, error) {
	err := s.App.PaymentService(ctx).RefundPayment(ctx, req.PaymentId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to refund payment: %v", err)
	}

	return &pb.RefundPaymentResponse{
		Success: true,
	}, nil
}

func (s *Server) ChargeWallet(ctx context.Context, req *pb.ChargeWalletRequest) (*pb.ChargeWalletResponse, error) {
	amount := &types.Money{
		Amount:   req.Amount.Amount,
		Currency: req.Amount.Currency,
	}

	err := s.App.PaymentService(ctx).ChargeWallet(ctx, walletDomain.WalletID(req.WalletId), amount)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to charge wallet: %v", err)
	}

	return &pb.ChargeWalletResponse{
		Success: true,
	}, nil
}

func (s *Server) WithdrawFromWallet(ctx context.Context, req *pb.WithdrawFromWalletRequest) (*pb.WithdrawFromWalletResponse, error) {
	amount := &types.Money{
		Amount:   req.Amount.Amount,
		Currency: req.Amount.Currency,
	}

	err := s.App.PaymentService(ctx).WithdrawFromWallet(ctx, walletDomain.WalletID(req.WalletId), amount)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to withdraw from wallet: %v", err)
	}

	return &pb.WithdrawFromWalletResponse{
		Success: true,
	}, nil
}

func (s *Server) GetTransactionHistory(ctx context.Context, req *pb.GetTransactionHistoryRequest) (*pb.GetTransactionHistoryResponse, error) {
	transactions, err := s.App.PaymentService(ctx).GetTransactionHistory(ctx, walletDomain.WalletID(req.WalletId))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get transaction history: %v", err)
	}

	protoTxs := make([]*pb.Transaction, len(transactions))
	for i, tx := range transactions {
		protoTxs[i] = convertDomainTxToProto(tx)
	}

	return &pb.GetTransactionHistoryResponse{
		Transactions: protoTxs,
	}, nil
}
