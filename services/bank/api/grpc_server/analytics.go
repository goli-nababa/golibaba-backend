package grpc_server

import (
	"bank_service/internal/common/types"
	domain2 "bank_service/internal/services/analytics/domain"
	txDomain "bank_service/internal/services/transaction/domain"
	walletDomain "bank_service/internal/services/wallet/domain"
	pb "bank_service/proto/gen/go/bank/v1"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) TrackTransaction(ctx context.Context, req *pb.TrackTransactionRequest) (*pb.TrackTransactionResponse, error) {
	err := s.App.AnalyticsService(ctx).TrackTransaction(ctx, convertProtoToTransactionDomain(req.Transaction))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to track transaction: %v", err)
	}

	return &pb.TrackTransactionResponse{
		Success: true,
	}, nil
}

func (s *Server) GenerateBusinessReport(ctx context.Context, req *pb.GenerateBusinessReportRequest) (*pb.GenerateBusinessReportResponse, error) {
	report, err := s.App.AnalyticsService(ctx).GenerateBusinessReport(
		ctx,
		req.BusinessId,
		req.StartDate.AsTime(),
		req.EndDate.AsTime(),
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate business report: %v", err)
	}

	return &pb.GenerateBusinessReportResponse{
		Report: convertAnalyticsReportToProto(report),
	}, nil
}

func (s *Server) GenerateBusinessStats(ctx context.Context, req *pb.GenerateBusinessStatsRequest) (*pb.GenerateBusinessStatsResponse, error) {
	stats, err := s.App.AnalyticsService(ctx).GenerateBusinessStats(
		ctx,
		req.BusinessId,
		req.StartDate.AsTime(),
		req.EndDate.AsTime(),
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate business stats: %v", err)
	}

	return &pb.GenerateBusinessStatsResponse{
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

func (s *Server) GetCommissionHistory(ctx context.Context, req *pb.GetCommissionHistoryRequest) (*pb.GetCommissionHistoryResponse, error) {
	entries, err := s.App.AnalyticsService(ctx).GetCommissionHistory(
		ctx,
		req.BusinessId,
		req.StartDate.AsTime(),
		req.EndDate.AsTime(),
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get commission history: %v", err)
	}

	protoEntries := make([]*pb.CommissionEntry, len(entries))
	for i, entry := range entries {
		protoEntries[i] = &pb.CommissionEntry{
			TransactionId: string(entry.TransactionID),
			Amount: &pb.Money{
				Amount:   entry.Amount.Amount,
				Currency: entry.Amount.Currency,
			},
			Rate:        entry.Rate,
			Status:      entry.Status,
			ProcessedAt: timestamppb.New(entry.ProcessedAt),
		}
	}

	return &pb.GetCommissionHistoryResponse{
		Entries: protoEntries,
	}, nil
}

func (s *Server) GetPayoutHistory(ctx context.Context, req *pb.GetPayoutHistoryRequest) (*pb.GetPayoutHistoryResponse, error) {
	entries, err := s.App.AnalyticsService(ctx).GetPayoutHistory(
		ctx,
		walletDomain.WalletID(req.WalletId),
		req.StartDate.AsTime(),
		req.EndDate.AsTime(),
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get payout history: %v", err)
	}

	protoEntries := make([]*pb.PayoutEntry, len(entries))
	for i, entry := range entries {
		protoEntries[i] = &pb.PayoutEntry{
			Id: entry.ID,
			Amount: &pb.Money{
				Amount:   entry.Amount.Amount,
				Currency: entry.Amount.Currency,
			},
			Status:      entry.Status,
			RequestedAt: timestamppb.New(entry.RequestedAt),
			//todo
			//ProcessedAt: entry.ProcessedAt != nil ? timestamppb.New(*entry.ProcessedAt) : nil,
			ReferenceId: entry.ReferenceID,
		}
	}

	return &pb.GetPayoutHistoryResponse{
		Entries: protoEntries,
	}, nil
}

func convertProtoToTransactionDomain(protoTx *pb.Transaction) *txDomain.Transaction {
	tx := &txDomain.Transaction{
		ID:           txDomain.TransactionID(protoTx.Id),
		FromWalletID: walletDomain.WalletID(protoTx.FromWalletId),
		ToWalletID:   walletDomain.WalletID(protoTx.ToWalletId),
		Amount: &types.Money{
			Amount:   protoTx.Amount.Amount,
			Currency: protoTx.Amount.Currency,
		},
		Type:          txDomain.TransactionType(protoTx.Type),
		Status:        txDomain.TransactionStatus(protoTx.Status),
		Description:   protoTx.Description,
		ReferenceID:   protoTx.ReferenceId,
		FailureReason: protoTx.FailureReason,
		Metadata:      make(map[string]interface{}),
		CreatedAt:     protoTx.CreatedAt.AsTime(),
		UpdatedAt:     protoTx.UpdatedAt.AsTime(),
		Version:       int(protoTx.Version),
	}

	if protoTx.CompletedAt != nil {
		completedAt := protoTx.CompletedAt.AsTime()
		tx.CompletedAt = &completedAt
	}

	for k, v := range protoTx.Metadata {
		tx.Metadata[k] = v
	}

	tx.StatusHistory = make([]txDomain.StatusChange, len(protoTx.StatusHistory))
	for i, change := range protoTx.StatusHistory {
		tx.StatusHistory[i] = txDomain.StatusChange{
			FromStatus: txDomain.TransactionStatus(change.FromStatus),
			ToStatus:   txDomain.TransactionStatus(change.ToStatus),
			Reason:     change.Reason,
			ChangedAt:  change.ChangedAt.AsTime(),
		}
	}

	return tx
}

func convertAnalyticsReportToProto(report *domain2.AnalyticsReport) *pb.AnalyticsReport {
	metrics := make(map[string]float64)
	for k, v := range report.Metrics {
		metrics[k] = v
	}

	trends := make(map[string]*pb.DataPointList)
	for k, points := range report.Trends {
		protoPoints := make([]*pb.DataPoint, len(points))
		for i, point := range points {
			protoPoints[i] = &pb.DataPoint{
				Timestamp: timestamppb.New(point.Timestamp),
				Value:     point.Value,
				Label:     point.Label,
			}
		}
		trends[k] = &pb.DataPointList{
			Points: protoPoints,
		}
	}

	comparisons := make(map[string]*pb.Comparison)
	for k, comp := range report.Comparisons {
		comparisons[k] = &pb.Comparison{
			CurrentValue:  comp.CurrentValue,
			PreviousValue: comp.PreviousValue,
			ChangePercent: comp.ChangePercent,
		}
	}

	return &pb.AnalyticsReport{
		Id:          report.ID,
		BusinessId:  report.BusinessID,
		StartDate:   timestamppb.New(report.StartDate),
		EndDate:     timestamppb.New(report.EndDate),
		Metrics:     metrics,
		Trends:      trends,
		Comparisons: comparisons,
		GeneratedAt: timestamppb.New(report.GeneratedAt),
	}
}

func convertDataPointToProto(point domain2.DataPoint) *pb.DataPoint {
	return &pb.DataPoint{
		Timestamp: timestamppb.New(point.Timestamp),
		Value:     point.Value,
		Label:     point.Label,
	}
}

func convertComparisonToProto(comp domain2.Comparison) *pb.Comparison {
	return &pb.Comparison{
		CurrentValue:  comp.CurrentValue,
		PreviousValue: comp.PreviousValue,
		ChangePercent: comp.ChangePercent,
	}
}
