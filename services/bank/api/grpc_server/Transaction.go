package grpc_server

import (
	"bank_service/internal/common/types"
	txDomain "bank_service/internal/services/transaction/domain"
	"bank_service/internal/services/transaction/port"
	walletDomain "bank_service/internal/services/wallet/domain"
	pb "bank_service/proto/gen/go/bank/v1"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) CreateTransaction(ctx context.Context, req *pb.CreateTransactionRequest) (*pb.CreateTransactionResponse, error) {
	amount := &types.Money{
		Amount:   req.Amount.Amount,
		Currency: req.Amount.Currency,
	}

	txReq := &port.TransactionRequest{
		FromWalletID: walletDomain.WalletID(req.FromWalletId),
		ToWalletID:   walletDomain.WalletID(req.ToWalletId),
		Amount:       amount,
		Type:         txDomain.TransactionType(req.Type),
		Description:  req.Description,
	}

	tx, err := s.App.TransactionService(ctx).CreateTransaction(ctx, txReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create transaction: %v", err)
	}

	return &pb.CreateTransactionResponse{
		Transaction: convertDomainTxToProto(tx),
	}, nil
}

func (s *Server) ProcessTransaction(ctx context.Context, req *pb.ProcessTransactionRequest) (*pb.ProcessTransactionResponse, error) {
	err := s.App.TransactionService(ctx).ProcessTransaction(ctx, txDomain.TransactionID(req.TransactionId))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to process transaction: %v", err)
	}

	tx, err := s.App.TransactionService(ctx).GetTransaction(ctx, txDomain.TransactionID(req.TransactionId))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get updated transaction: %v", err)
	}

	return &pb.ProcessTransactionResponse{
		Transaction: convertDomainTxToProto(tx),
	}, nil
}

func (s *Server) CancelTransaction(ctx context.Context, req *pb.CancelTransactionRequest) (*pb.CancelTransactionResponse, error) {
	err := s.App.TransactionService(ctx).CancelTransaction(ctx, txDomain.TransactionID(req.TransactionId), req.Reason)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to cancel transaction: %v", err)
	}

	tx, err := s.App.TransactionService(ctx).GetTransaction(ctx, txDomain.TransactionID(req.TransactionId))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get updated transaction: %v", err)
	}

	return &pb.CancelTransactionResponse{
		Transaction: convertDomainTxToProto(tx),
	}, nil
}

func (s *Server) GetTransaction(ctx context.Context, req *pb.GetTransactionRequest) (*pb.GetTransactionResponse, error) {
	tx, err := s.App.TransactionService(ctx).GetTransaction(ctx, txDomain.TransactionID(req.TransactionId))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get transaction: %v", err)
	}

	if tx == nil {
		return nil, status.Error(codes.NotFound, "transaction not found")
	}

	return &pb.GetTransactionResponse{
		Transaction: convertDomainTxToProto(tx),
	}, nil
}

func (s *Server) ListTransactions(ctx context.Context, req *pb.ListTransactionsRequest) (*pb.ListTransactionsResponse, error) {
	filter := &port.TransactionFilter{
		WalletID: walletDomain.WalletID(req.WalletId),
	}

	if len(req.Status) > 0 {
		statuses := make([]txDomain.TransactionStatus, len(req.Status))
		for i, status := range req.Status {
			statuses[i] = txDomain.TransactionStatus(status)
		}
		filter.Status = statuses
	}

	if len(req.Types) > 0 {
		types := make([]txDomain.TransactionType, len(req.Types))
		for i, txType := range req.Types {
			types[i] = txDomain.TransactionType(txType)
		}
		filter.Types = types
	}

	if req.FromDate != nil {
		filter.From = req.FromDate.AsTime()
	}
	if req.ToDate != nil {
		filter.To = req.ToDate.AsTime()
	}

	transactions, err := s.App.TransactionService(ctx).ListTransactions(ctx, filter)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list transactions: %v", err)
	}

	protoTxs := make([]*pb.Transaction, len(transactions))
	for i, tx := range transactions {
		protoTxs[i] = convertDomainTxToProto(tx)
	}

	return &pb.ListTransactionsResponse{
		Transactions: protoTxs,
	}, nil
}

func convertDomainTxToProto(tx *txDomain.Transaction) *pb.Transaction {
	statusHistory := make([]*pb.StatusChange, len(tx.StatusHistory))
	for i, change := range tx.StatusHistory {
		statusHistory[i] = &pb.StatusChange{
			FromStatus: string(change.FromStatus),
			ToStatus:   string(change.ToStatus),
			Reason:     change.Reason,
			ChangedAt:  timestamppb.New(change.ChangedAt),
		}
	}

	protoTx := &pb.Transaction{
		Id:            string(tx.ID),
		FromWalletId:  string(tx.FromWalletID),
		ToWalletId:    string(tx.ToWalletID),
		Amount:        &pb.Money{Amount: tx.Amount.Amount, Currency: tx.Amount.Currency},
		Type:          string(tx.Type),
		Status:        string(tx.Status),
		Description:   tx.Description,
		ReferenceId:   tx.ReferenceID,
		FailureReason: tx.FailureReason,
		Metadata:      make(map[string]string),
		StatusHistory: statusHistory,
		CreatedAt:     timestamppb.New(tx.CreatedAt),
		UpdatedAt:     timestamppb.New(tx.UpdatedAt),
		Version:       int32(tx.Version),
	}

	if tx.CompletedAt != nil {
		protoTx.CompletedAt = timestamppb.New(*tx.CompletedAt)
	}

	for k, v := range tx.Metadata {
		if str, ok := v.(string); ok {
			protoTx.Metadata[k] = str
		}
	}

	return protoTx
}
