package grpc_server

import (
	"bank_service/internal/common/types"
	businessDomain "bank_service/internal/services/business/domain"
	txDomain "bank_service/internal/services/transaction/domain"
	walletDomain "bank_service/internal/services/wallet/domain"
	pb "bank_service/proto/gen/go/bank/v1"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) CalculateCommission(ctx context.Context, req *pb.CalculateCommissionRequest) (*pb.CalculateCommissionResponse, error) {
	tx := &txDomain.Transaction{
		ID:           txDomain.TransactionID(req.Transaction.Id),
		FromWalletID: walletDomain.WalletID(req.Transaction.FromWalletId),
		ToWalletID:   walletDomain.WalletID(req.Transaction.ToWalletId),
		Amount: &types.Money{
			Amount:   req.Transaction.Amount.Amount,
			Currency: req.Transaction.Amount.Currency,
		},
		Type:        txDomain.TransactionType(req.Transaction.Type),
		Status:      txDomain.TransactionStatus(req.Transaction.Status),
		Description: req.Transaction.Description,
		ReferenceID: req.Transaction.ReferenceId,
	}

	commission, err := s.App.CommissionService(ctx).CalculateCommission(ctx, tx, businessDomain.BusinessType(req.BusinessType))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to calculate commission: %v", err)
	}

	protoCommission := &pb.Commission{
		Id:            commission.ID,
		TransactionId: string(commission.TransactionID),
		Amount: &pb.Money{
			Amount:   commission.Amount.Amount,
			Currency: commission.Amount.Currency,
		},
		Rate:         commission.Rate,
		RecipientId:  string(commission.RecipientID),
		BusinessType: string(commission.BusinessType),
		Status:       string(commission.Status),
		CreatedAt:    timestamppb.New(commission.CreatedAt),
		Description:  commission.Description,
	}

	if commission.PaidAt != nil {
		protoCommission.PaidAt = timestamppb.New(*commission.PaidAt)
	}

	return &pb.CalculateCommissionResponse{
		Commission: protoCommission,
	}, nil
}

func (s *Server) ProcessCommission(ctx context.Context, req *pb.ProcessCommissionRequest) (*pb.ProcessCommissionResponse, error) {
	err := s.App.CommissionService(ctx).ProcessCommission(ctx, req.CommissionId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to process commission: %v", err)
	}

	return &pb.ProcessCommissionResponse{
		Success: true,
	}, nil
}

func (s *Server) GetCommission(ctx context.Context, req *pb.GetCommissionRequest) (*pb.GetCommissionResponse, error) {
	commission, err := s.App.CommissionService(ctx).GetCommission(ctx, req.CommissionId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get commission: %v", err)
	}

	if commission == nil {
		return nil, status.Error(codes.NotFound, "commission not found")
	}

	protoCommission := &pb.Commission{
		Id:            commission.ID,
		TransactionId: string(commission.TransactionID),
		Amount: &pb.Money{
			Amount:   commission.Amount.Amount,
			Currency: commission.Amount.Currency,
		},
		Rate:         commission.Rate,
		RecipientId:  string(commission.RecipientID),
		BusinessType: string(commission.BusinessType),
		Status:       string(commission.Status),
		CreatedAt:    timestamppb.New(commission.CreatedAt),
		Description:  commission.Description,
	}

	if commission.PaidAt != nil {
		protoCommission.PaidAt = timestamppb.New(*commission.PaidAt)
	}

	return &pb.GetCommissionResponse{
		Commission: protoCommission,
	}, nil
}

func (s *Server) GetPendingCommissions(ctx context.Context, req *pb.GetPendingCommissionsRequest) (*pb.GetPendingCommissionsResponse, error) {
	commissions, err := s.App.CommissionService(ctx).GetPendingCommissions(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get pending commissions: %v", err)
	}

	protoCommissions := make([]*pb.Commission, len(commissions))
	for i, commission := range commissions {
		protoCommissions[i] = &pb.Commission{
			Id:            commission.ID,
			TransactionId: string(commission.TransactionID),
			Amount: &pb.Money{
				Amount:   commission.Amount.Amount,
				Currency: commission.Amount.Currency,
			},
			Rate:         commission.Rate,
			RecipientId:  string(commission.RecipientID),
			BusinessType: string(commission.BusinessType),
			Status:       string(commission.Status),
			CreatedAt:    timestamppb.New(commission.CreatedAt),
			Description:  commission.Description,
		}

		if commission.PaidAt != nil {
			protoCommissions[i].PaidAt = timestamppb.New(*commission.PaidAt)
		}
	}

	return &pb.GetPendingCommissionsResponse{
		Commissions: protoCommissions,
	}, nil
}

func (s *Server) GetFailedCommissions(ctx context.Context, req *pb.GetFailedCommissionsRequest) (*pb.GetFailedCommissionsResponse, error) {
	commissions, err := s.App.CommissionService(ctx).GetFailedCommissions(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get failed commissions: %v", err)
	}

	protoCommissions := make([]*pb.Commission, len(commissions))
	for i, commission := range commissions {
		protoCommissions[i] = &pb.Commission{
			Id:            commission.ID,
			TransactionId: string(commission.TransactionID),
			Amount: &pb.Money{
				Amount:   commission.Amount.Amount,
				Currency: commission.Amount.Currency,
			},
			Rate:         commission.Rate,
			RecipientId:  string(commission.RecipientID),
			BusinessType: string(commission.BusinessType),
			Status:       string(commission.Status),
			CreatedAt:    timestamppb.New(commission.CreatedAt),
			Description:  commission.Description,
		}

		if commission.PaidAt != nil {
			protoCommissions[i].PaidAt = timestamppb.New(*commission.PaidAt)
		}

	}

	return &pb.GetFailedCommissionsResponse{
		Commissions: protoCommissions,
	}, nil
}

func (s *Server) RetryFailedCommissions(ctx context.Context, req *pb.RetryFailedCommissionsRequest) (*pb.RetryFailedCommissionsResponse, error) {
	err := s.App.CommissionService(ctx).RetryFailedCommissions(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to retry failed commissions: %v", err)
	}

	return &pb.RetryFailedCommissionsResponse{
		RetriedCount: 0,
		SuccessCount: 0,
	}, nil
}
