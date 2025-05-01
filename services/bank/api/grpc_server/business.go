package grpc_server

import (
	"bank_service/internal/common/types"
	businessDomain "bank_service/internal/services/business/domain"
	walletDomain "bank_service/internal/services/wallet/domain"
	pb "bank_service/proto/gen/go/bank/v1"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) CreateBusinessWallet(ctx context.Context, req *pb.CreateBusinessWalletRequest) (*pb.CreateBusinessWalletResponse, error) {
	wallet, err := s.App.BusinessService(ctx).CreateBusinessWallet(
		ctx,
		req.BusinessId,
		businessDomain.BusinessType(req.BusinessType),
		req.Currency,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create business wallet: %v", err)
	}

	return &pb.CreateBusinessWalletResponse{
		Wallet: convertBusinessWalletToProto(wallet),
	}, nil
}

func (s *Server) GetBusinessWallet(ctx context.Context, req *pb.GetBusinessWalletRequest) (*pb.GetBusinessWalletResponse, error) {
	wallet, err := s.App.BusinessService(ctx).GetBusinessWallet(ctx, walletDomain.WalletID(req.WalletId))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get business wallet: %v", err)
	}

	if wallet == nil {
		return nil, status.Error(codes.NotFound, "business wallet not found")
	}

	return &pb.GetBusinessWalletResponse{
		Wallet: convertBusinessWalletToProto(wallet),
	}, nil
}

func (s *Server) UpdateBusinessWallet(ctx context.Context, req *pb.UpdateBusinessWalletRequest) (*pb.UpdateBusinessWalletResponse, error) {
	wallet := &businessDomain.BusinessWallet{
		Wallet: &walletDomain.Wallet{
			ID: walletDomain.WalletID(req.Wallet.Id),
			Balance: &types.Money{
				Amount:   req.Wallet.Balance.Amount,
				Currency: req.Wallet.Balance.Currency,
			},
			Status: walletDomain.WalletStatus(req.Wallet.Status),
		},
		BusinessID:     req.Wallet.BusinessId,
		BusinessType:   businessDomain.BusinessType(req.Wallet.BusinessType),
		CommissionRate: req.Wallet.CommissionRate,
		PayoutSchedule: req.Wallet.PayoutSchedule,
	}

	if req.Wallet.LastPayoutDate != nil {
		lastPayoutDate := req.Wallet.LastPayoutDate.AsTime()
		wallet.LastPayoutDate = &lastPayoutDate
	}

	if req.Wallet.MinimumPayout != nil {
		wallet.MinimumPayout = &types.Money{
			Amount:   req.Wallet.MinimumPayout.Amount,
			Currency: req.Wallet.MinimumPayout.Currency,
		}
	}

	if req.Wallet.BankInfo != nil {
		wallet.BankInfo = &businessDomain.BankAccountInfo{
			AccountNumber: req.Wallet.BankInfo.AccountNumber,
			IBAN:          req.Wallet.BankInfo.Iban,
			BankName:      req.Wallet.BankInfo.BankName,
			AccountName:   req.Wallet.BankInfo.AccountName,
			CardNumber:    req.Wallet.BankInfo.CardNumber,
		}
	}

	err := s.App.BusinessService(ctx).UpdateBusinessWallet(ctx, wallet)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update business wallet: %v", err)
	}

	updatedWallet, err := s.App.BusinessService(ctx).GetBusinessWallet(ctx, wallet.ID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get updated business wallet: %v", err)
	}

	return &pb.UpdateBusinessWalletResponse{
		Wallet: convertBusinessWalletToProto(updatedWallet),
	}, nil
}

func (s *Server) SetPayoutSchedule(ctx context.Context, req *pb.SetPayoutScheduleRequest) (*pb.SetPayoutScheduleResponse, error) {
	err := s.App.BusinessService(ctx).SetPayoutSchedule(ctx, walletDomain.WalletID(req.WalletId), req.Schedule)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to set payout schedule: %v", err)
	}

	return &pb.SetPayoutScheduleResponse{
		Success: true,
	}, nil
}

func (s *Server) RequestPayout(ctx context.Context, req *pb.RequestPayoutRequest) (*pb.RequestPayoutResponse, error) {
	amount := &types.Money{
		Amount:   req.Amount.Amount,
		Currency: req.Amount.Currency,
	}

	err := s.App.BusinessService(ctx).RequestPayout(ctx, walletDomain.WalletID(req.WalletId), amount)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to request payout: %v", err)
	}

	return &pb.RequestPayoutResponse{
		Success: true,
	}, nil
}

func (s *Server) GetBusinessStats(ctx context.Context, req *pb.GetBusinessStatsRequest) (*pb.GetBusinessStatsResponse, error) {
	stats, err := s.App.BusinessService(ctx).GetBusinessStats(
		ctx,
		req.BusinessId,
		req.StartDate.AsTime(),
		req.EndDate.AsTime(),
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get business stats: %v", err)
	}

	return &pb.GetBusinessStatsResponse{
		Stats: &pb.BusinessStats{
			TotalRevenue: &pb.Money{
				Amount:   stats.TotalRevenue.Amount,
				Currency: stats.TotalRevenue.Currency,
			},
			TotalCommission: &pb.Money{
				Amount:   stats.TotalCommission.Amount,
				Currency: stats.TotalCommission.Currency,
			},
			TotalPayouts: &pb.Money{
				Amount:   stats.TotalPayouts.Amount,
				Currency: stats.TotalPayouts.Currency,
			},
			TransactionCount:  stats.TransactionCount,
			SuccessfulTxCount: stats.SuccessfulTxCount,
			FailedTxCount:     stats.FailedTxCount,
			AverageOrderValue: &pb.Money{
				Amount:   stats.AverageOrderValue.Amount,
				Currency: stats.AverageOrderValue.Currency,
			},
			CommissionRate: stats.CommissionRate,
			PeriodStart:    timestamppb.New(stats.Period.Start),
			PeriodEnd:      timestamppb.New(stats.Period.End),
		},
	}, nil
}

func convertBusinessWalletToProto(wallet *businessDomain.BusinessWallet) *pb.BusinessWallet {
	protoWallet := &pb.BusinessWallet{
		Id:           string(wallet.ID),
		BusinessId:   wallet.BusinessID,
		BusinessType: string(wallet.BusinessType),
		Balance: &pb.Money{
			Amount:   wallet.Balance.Amount,
			Currency: wallet.Balance.Currency,
		},
		Status:         string(wallet.Status),
		CommissionRate: wallet.CommissionRate,
		PayoutSchedule: wallet.PayoutSchedule,
	}

	if wallet.LastPayoutDate != nil {
		protoWallet.LastPayoutDate = timestamppb.New(*wallet.LastPayoutDate)
	}

	if wallet.MinimumPayout != nil {
		protoWallet.MinimumPayout = &pb.Money{
			Amount:   wallet.MinimumPayout.Amount,
			Currency: wallet.MinimumPayout.Currency,
		}
	}

	if wallet.BankInfo != nil {
		protoWallet.BankInfo = &pb.BankAccountInfo{
			AccountNumber: wallet.BankInfo.AccountNumber,
			Iban:          wallet.BankInfo.IBAN,
			BankName:      wallet.BankInfo.BankName,
			AccountName:   wallet.BankInfo.AccountName,
			CardNumber:    wallet.BankInfo.CardNumber,
		}
	}

	return protoWallet
}
