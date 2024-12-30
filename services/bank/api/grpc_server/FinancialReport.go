package grpc_server

import (
	"bank_service/internal/services/financial_report/domain"
	pb "bank_service/proto/gen/go/bank/v1"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func (s *Server) GenerateDailyReport(ctx context.Context, req *pb.GenerateDailyReportRequest) (*pb.GenerateDailyReportResponse, error) {
	report, err := s.App.FinancialReportService(ctx).GenerateDailyReport(
		ctx,
		req.BusinessId,
		req.Date.AsTime(),
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate daily report: %v", err)
	}

	return &pb.GenerateDailyReportResponse{
		Report: convertFinancialReportToProto(report),
	}, nil
}

func (s *Server) GenerateMonthlyReport(ctx context.Context, req *pb.GenerateMonthlyReportRequest) (*pb.GenerateMonthlyReportResponse, error) {
	report, err := s.App.FinancialReportService(ctx).GenerateMonthlyReport(
		ctx,
		req.BusinessId,
		int(req.Year),
		time.Month(int(req.Month)),
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate monthly report: %v", err)
	}

	return &pb.GenerateMonthlyReportResponse{
		Report: convertFinancialReportToProto(report),
	}, nil
}

func (s *Server) GenerateCustomReport(ctx context.Context, req *pb.GenerateCustomReportRequest) (*pb.GenerateCustomReportResponse, error) {
	report, err := s.App.FinancialReportService(ctx).GenerateCustomReport(
		ctx,
		req.BusinessId,
		req.StartDate.AsTime(),
		req.EndDate.AsTime(),
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate custom report: %v", err)
	}

	return &pb.GenerateCustomReportResponse{
		Report: convertFinancialReportToProto(report),
	}, nil
}

func (s *Server) GetReportByID(ctx context.Context, req *pb.GetReportByIDRequest) (*pb.GetReportByIDResponse, error) {
	report, err := s.App.FinancialReportService(ctx).GetReportByID(ctx, req.ReportId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get report: %v", err)
	}

	if report == nil {
		return nil, status.Error(codes.NotFound, "report not found")
	}

	return &pb.GetReportByIDResponse{
		Report: convertFinancialReportToProto(report),
	}, nil
}

func (s *Server) ExportReport(ctx context.Context, req *pb.ExportReportRequest) (*pb.ExportReportResponse, error) {
	data, err := s.App.FinancialReportService(ctx).ExportReport(
		ctx,
		req.ReportId,
		domain.ReportFormat(req.Format),
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to export report: %v", err)
	}

	return &pb.ExportReportResponse{
		Data: data,
	}, nil
}

func convertFinancialReportToProto(report *domain.FinancialReport) *pb.FinancialReport {
	metrics := make(map[string]*pb.Money)
	for key, value := range report.Metrics {
		metrics[key] = &pb.Money{
			Amount:   value.Amount,
			Currency: value.Currency,
		}
	}

	return &pb.FinancialReport{
		Id:          report.ID,
		ReportType:  string(report.ReportType),
		BusinessId:  report.BusinessID,
		StartDate:   timestamppb.New(report.StartDate),
		EndDate:     timestamppb.New(report.EndDate),
		Granularity: string(report.Granularity),
		Status:      string(report.Status),
		Metrics:     metrics,
		GeneratedAt: timestamppb.New(report.GeneratedAt),
		GeneratedBy: report.GeneratedBy,
		Format:      string(report.Format),
		DataUrl:     report.DataURL,
	}
}
