package grpc_server

import (
	moneyDomain "bank_service/internal/common/types"
	"bank_service/internal/services/wallet/domain"
	pb "bank_service/proto/gen/go/bank/v1"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateWallet(ctx context.Context, req *pb.CreateWalletRequest) (*pb.CreateWalletResponse, error) {
	wallet, err := s.App.WalletService(ctx).CreateWallet(
		ctx,
		req.OwnerId,
		domain.WalletType(req.WalletType),
		req.Currency,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create wallet: %v", err)
	}

	return &pb.CreateWalletResponse{
		Wallet: &pb.Wallet{
			Id:        string(wallet.ID),
			UserId:    wallet.UserID,
			Balance:   &pb.Money{Amount: wallet.Balance.Amount, Currency: wallet.Balance.Currency},
			Status:    string(wallet.Status),
			CreatedAt: nil,
			UpdatedAt: nil,
			Version:   int32(wallet.Version),
		},
	}, nil
}

func (s *Server) GetWallet(ctx context.Context, req *pb.GetWalletRequest) (*pb.GetWalletResponse, error) {
	wallet, err := s.App.WalletService(ctx).GetWallet(ctx, domain.WalletID(req.WalletId))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get wallet: %v", err)
	}

	if wallet == nil {
		return nil, status.Error(codes.NotFound, "wallet not found")
	}

	return &pb.GetWalletResponse{
		Wallet: &pb.Wallet{
			Id:        string(wallet.ID),
			UserId:    wallet.UserID,
			Balance:   &pb.Money{Amount: wallet.Balance.Amount, Currency: wallet.Balance.Currency},
			Status:    string(wallet.Status),
			CreatedAt: nil,
			UpdatedAt: nil,
			Version:   int32(wallet.Version),
		},
	}, nil
}

func (s *Server) GetWalletsByUser(ctx context.Context, req *pb.GetWalletsByUserRequest) (*pb.GetWalletsByUserResponse, error) {
	wallets, err := s.App.WalletService(ctx).GetWalletsByUser(ctx, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user wallets: %v", err)
	}

	pbWallets := make([]*pb.Wallet, len(wallets))
	for i, wallet := range wallets {
		pbWallets[i] = &pb.Wallet{
			Id:        string(wallet.ID),
			UserId:    wallet.UserID,
			Balance:   &pb.Money{Amount: wallet.Balance.Amount, Currency: wallet.Balance.Currency},
			Status:    string(wallet.Status),
			CreatedAt: nil,
			UpdatedAt: nil,
			Version:   int32(wallet.Version),
		}
	}

	return &pb.GetWalletsByUserResponse{
		Wallets: pbWallets,
	}, nil
}

func (s *Server) Credit(ctx context.Context, req *pb.CreditRequest) (*pb.CreditResponse, error) {
	amount := &moneyDomain.Money{
		Amount:   req.Amount.Amount,
		Currency: req.Amount.Currency,
	}

	err := s.App.WalletService(ctx).Credit(ctx, domain.WalletID(req.WalletId), amount)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to credit wallet: %v", err)
	}

	wallet, err := s.App.WalletService(ctx).GetWallet(ctx, domain.WalletID(req.WalletId))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get updated wallet: %v", err)
	}

	return &pb.CreditResponse{
		UpdatedWallet: &pb.Wallet{
			Id:        string(wallet.ID),
			UserId:    wallet.UserID,
			Balance:   &pb.Money{Amount: wallet.Balance.Amount, Currency: wallet.Balance.Currency},
			Status:    string(wallet.Status),
			CreatedAt: nil,
			UpdatedAt: nil,
			Version:   int32(wallet.Version),
		},
	}, nil
}

func (s *Server) Debit(ctx context.Context, req *pb.DebitRequest) (*pb.DebitResponse, error) {
	amount := &moneyDomain.Money{
		Amount:   req.Amount.Amount,
		Currency: req.Amount.Currency,
	}

	err := s.App.WalletService(ctx).Debit(ctx, domain.WalletID(req.WalletId), amount)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to debit wallet: %v", err)
	}

	wallet, err := s.App.WalletService(ctx).GetWallet(ctx, domain.WalletID(req.WalletId))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get updated wallet: %v", err)
	}

	return &pb.DebitResponse{
		UpdatedWallet: &pb.Wallet{
			Id:        string(wallet.ID),
			UserId:    wallet.UserID,
			Balance:   &pb.Money{Amount: wallet.Balance.Amount, Currency: wallet.Balance.Currency},
			Status:    string(wallet.Status),
			CreatedAt: nil,
			UpdatedAt: nil,
			Version:   int32(wallet.Version),
		},
	}, nil
}

func (s *Server) Transfer(ctx context.Context, req *pb.TransferRequest) (*pb.TransferResponse, error) {
	if req.Amount == nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid amount")
	}
	amount := &moneyDomain.Money{
		Amount:   req.Amount.Amount,
		Currency: req.Amount.Currency,
	}

	err := s.App.WalletService(ctx).Transfer(
		ctx,
		domain.WalletID(req.FromWalletId),
		domain.WalletID(req.ToWalletId),
		amount,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to transfer: %v", err)
	}

	fromWallet, err := s.App.WalletService(ctx).GetWallet(ctx, domain.WalletID(req.FromWalletId))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get source wallet: %v", err)
	}

	toWallet, err := s.App.WalletService(ctx).GetWallet(ctx, domain.WalletID(req.ToWalletId))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get destination wallet: %v", err)
	}

	return &pb.TransferResponse{
		FromWallet: &pb.Wallet{
			Id:        string(fromWallet.ID),
			UserId:    fromWallet.UserID,
			Balance:   &pb.Money{Amount: fromWallet.Balance.Amount, Currency: fromWallet.Balance.Currency},
			Status:    string(fromWallet.Status),
			CreatedAt: nil,
			UpdatedAt: nil,
			Version:   int32(fromWallet.Version),
		},
		ToWallet: &pb.Wallet{
			Id:        string(toWallet.ID),
			UserId:    toWallet.UserID,
			Balance:   &pb.Money{Amount: toWallet.Balance.Amount, Currency: toWallet.Balance.Currency},
			Status:    string(toWallet.Status),
			CreatedAt: nil,
			UpdatedAt: nil,
			Version:   int32(toWallet.Version),
		},
	}, nil
}

func (s *Server) UpdateWalletStatus(ctx context.Context, req *pb.UpdateWalletStatusRequest) (*pb.UpdateWalletStatusResponse, error) {
	err := s.App.WalletService(ctx).UpdateWalletStatus(
		ctx,
		domain.WalletID(req.WalletId),
		domain.WalletStatus(req.Status),
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update wallet status: %v", err)
	}

	wallet, err := s.App.WalletService(ctx).GetWallet(ctx, domain.WalletID(req.WalletId))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get updated wallet: %v", err)
	}

	return &pb.UpdateWalletStatusResponse{
		Wallet: &pb.Wallet{
			Id:        string(wallet.ID),
			UserId:    wallet.UserID,
			Balance:   &pb.Money{Amount: wallet.Balance.Amount, Currency: wallet.Balance.Currency},
			Status:    string(wallet.Status),
			CreatedAt: nil,
			UpdatedAt: nil,
			Version:   int32(wallet.Version),
		},
	}, nil
}
