package client

import (
	"context"
	"crypto/tls"
	"fmt"
	pb "github.com/goli-nababa/golibaba-backend/proto/pb/buf/bank/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type Client struct {
	SessionClient          pb.SessionServiceClient
	WalletService          pb.WalletServiceClient
	TransactionService     pb.TransactionServiceClient
	PaymentService         pb.PaymentServiceClient
	BusinessService        pb.BusinessServiceClient
	CommissionService      pb.CommissionServiceClient
	AnalyticsService       pb.AnalyticsServiceClient
	FinancialReportService pb.FinancialReportServiceClient

	conn   *grpc.ClientConn
	config *Config
}

type Config struct {
	Host       string
	Port       int
	UseTLS     bool
	TLSConfig  *tls.Config
	Timeout    time.Duration
	MaxRetries int
}

func DefaultConfig() *Config {
	return &Config{
		Host:       "localhost",
		Port:       8082,
		UseTLS:     false,
		Timeout:    time.Second * 300,
		MaxRetries: 3,
	}
}

func NewClient(cfg *Config) (*Client, error) {
	opts := []grpc.DialOption{}

	if cfg.UseTLS {
		if cfg.TLSConfig == nil {
			return nil, fmt.Errorf("TLS config required when UseTLS is true")
		}
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(cfg.TLSConfig)))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		opts...,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to server: %w", err)
	}

	c := &Client{
		conn:   conn,
		config: cfg,
	}

	c.SessionClient = pb.NewSessionServiceClient(conn)
	c.WalletService = pb.NewWalletServiceClient(conn)
	c.TransactionService = pb.NewTransactionServiceClient(conn)
	c.PaymentService = pb.NewPaymentServiceClient(conn)
	c.BusinessService = pb.NewBusinessServiceClient(conn)
	c.CommissionService = pb.NewCommissionServiceClient(conn)
	c.AnalyticsService = pb.NewAnalyticsServiceClient(conn)
	c.FinancialReportService = pb.NewFinancialReportServiceClient(conn)

	return c, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) CreateWallet(ctx context.Context, ownerID uint64, walletType, currency string) (*pb.Wallet, error) {
	ctx, cancel := context.WithTimeout(ctx, c.config.Timeout)
	defer cancel()

	req := &pb.CreateWalletRequest{
		OwnerId:    ownerID,
		WalletType: walletType,
		Currency:   currency,
	}

	resp, err := c.WalletService.CreateWallet(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to create wallet: %w", err)
	}

	return resp.Wallet, nil
}

func (c *Client) GetWallet(ctx context.Context, walletID string) (*pb.Wallet, error) {
	ctx, cancel := context.WithTimeout(ctx, c.config.Timeout)
	defer cancel()

	req := &pb.GetWalletRequest{
		WalletId: walletID,
	}

	resp, err := c.WalletService.GetWallet(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get wallet: %w", err)
	}

	return resp.Wallet, nil
}

func (c *Client) Credit(ctx context.Context, walletID string, amount *pb.Money) (*pb.Wallet, error) {
	ctx, cancel := context.WithTimeout(ctx, c.config.Timeout)
	defer cancel()

	req := &pb.CreditRequest{
		WalletId: walletID,
		Amount:   amount,
	}

	resp, err := c.WalletService.Credit(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to credit wallet: %w", err)
	}

	return resp.UpdatedWallet, nil
}

func (c *Client) Debit(ctx context.Context, walletID string, amount *pb.Money) (*pb.Wallet, error) {
	ctx, cancel := context.WithTimeout(ctx, c.config.Timeout)
	defer cancel()

	req := &pb.DebitRequest{
		WalletId: walletID,
		Amount:   amount,
	}

	resp, err := c.WalletService.Debit(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to debit wallet: %w", err)
	}

	return resp.UpdatedWallet, nil
}

func (c *Client) Transfer(ctx context.Context, fromWalletID, toWalletID string, amount *pb.Money) (*pb.TransferResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, c.config.Timeout)
	defer cancel()

	req := &pb.TransferRequest{
		FromWalletId: fromWalletID,
		ToWalletId:   toWalletID,
		Amount:       amount,
	}

	return c.WalletService.Transfer(ctx, req)
}

func (c *Client) UpdateWalletStatus(ctx context.Context, walletID, status string) (*pb.Wallet, error) {
	ctx, cancel := context.WithTimeout(ctx, c.config.Timeout)
	defer cancel()

	req := &pb.UpdateWalletStatusRequest{
		WalletId: walletID,
		Status:   status,
	}

	resp, err := c.WalletService.UpdateWalletStatus(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to update wallet status: %w", err)
	}

	return resp.Wallet, nil
}

func (c *Client) ProcessTransaction(ctx context.Context, transactionID string) (*pb.Transaction, error) {
	ctx, cancel := context.WithTimeout(ctx, c.config.Timeout)
	defer cancel()

	req := &pb.ProcessTransactionRequest{
		TransactionId: transactionID,
	}

	resp, err := c.TransactionService.ProcessTransaction(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to process transaction: %w", err)
	}

	return resp.Transaction, nil
}

func (c *Client) CancelTransaction(ctx context.Context, transactionID, reason string) (*pb.Transaction, error) {
	ctx, cancel := context.WithTimeout(ctx, c.config.Timeout)
	defer cancel()

	req := &pb.CancelTransactionRequest{
		TransactionId: transactionID,
		Reason:        reason,
	}

	resp, err := c.TransactionService.CancelTransaction(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to cancel transaction: %w", err)
	}

	return resp.Transaction, nil
}

func (c *Client) ListTransactions(ctx context.Context, req *pb.ListTransactionsRequest) ([]*pb.Transaction, error) {
	ctx, cancel := context.WithTimeout(ctx, c.config.Timeout)
	defer cancel()

	resp, err := c.TransactionService.ListTransactions(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to list transactions: %w", err)
	}

	return resp.Transactions, nil
}

func (c *Client) ProcessRefund(ctx context.Context, tx *pb.Transaction) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, c.config.Timeout)
	defer cancel()

	req := &pb.ProcessRefundRequest{
		Transaction: tx,
	}

	resp, err := c.PaymentService.ProcessRefund(ctx, req)
	if err != nil {
		return false, fmt.Errorf("failed to process refund: %w", err)
	}

	return resp.Success, nil
}

func (c *Client) VerifyPayment(ctx context.Context, referenceID string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, c.config.Timeout)
	defer cancel()

	req := &pb.VerifyPaymentRequest{
		ReferenceId: referenceID,
	}

	resp, err := c.PaymentService.VerifyPayment(ctx, req)
	if err != nil {
		return false, fmt.Errorf("failed to verify payment: %w", err)
	}

	return resp.Verified, nil
}

func (c *Client) RefundPayment(ctx context.Context, paymentID string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, c.config.Timeout)
	defer cancel()

	req := &pb.RefundPaymentRequest{
		PaymentId: paymentID,
	}

	resp, err := c.PaymentService.RefundPayment(ctx, req)
	if err != nil {
		return false, fmt.Errorf("failed to refund payment: %w", err)
	}

	return resp.Success, nil
}

func (c *Client) ChargeWallet(ctx context.Context, walletID string, amount *pb.Money) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, c.config.Timeout)
	defer cancel()

	req := &pb.ChargeWalletRequest{
		WalletId: walletID,
		Amount:   amount,
	}

	resp, err := c.PaymentService.ChargeWallet(ctx, req)
	if err != nil {
		return false, fmt.Errorf("failed to charge wallet: %w", err)
	}

	return resp.Success, nil
}

func (c *Client) WithdrawFromWallet(ctx context.Context, walletID string, amount *pb.Money) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, c.config.Timeout)
	defer cancel()

	req := &pb.WithdrawFromWalletRequest{
		WalletId: walletID,
		Amount:   amount,
	}

	resp, err := c.PaymentService.WithdrawFromWallet(ctx, req)
	if err != nil {
		return false, fmt.Errorf("failed to withdraw from wallet: %w", err)
	}

	return resp.Success, nil
}

func (c *Client) GetTransactionHistory(ctx context.Context, walletID string) ([]*pb.Transaction, error) {
	ctx, cancel := context.WithTimeout(ctx, c.config.Timeout)
	defer cancel()

	req := &pb.GetTransactionHistoryRequest{
		WalletId: walletID,
	}

	resp, err := c.PaymentService.GetTransactionHistory(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction history: %w", err)
	}

	return resp.Transactions, nil
}

func (c *Client) UpdateBusinessWallet(ctx context.Context, wallet *pb.BusinessWallet) (*pb.BusinessWallet, error) {
	ctx, cancel := context.WithTimeout(ctx, c.config.Timeout)
	defer cancel()

	req := &pb.UpdateBusinessWalletRequest{
		Wallet: wallet,
	}

	resp, err := c.BusinessService.UpdateBusinessWallet(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to update business wallet: %w", err)
	}

	return resp.Wallet, nil
}

func (c *Client) SetPayoutSchedule(ctx context.Context, walletID, schedule string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, c.config.Timeout)
	defer cancel()

	req := &pb.SetPayoutScheduleRequest{
		WalletId: walletID,
		Schedule: schedule,
	}

	resp, err := c.BusinessService.SetPayoutSchedule(ctx, req)
	if err != nil {
		return false, fmt.Errorf("failed to set payout schedule: %w", err)
	}

	return resp.Success, nil
}

func (c *Client) RequestPayout(ctx context.Context, walletID string, amount *pb.Money) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, c.config.Timeout)
	defer cancel()

	req := &pb.RequestPayoutRequest{
		WalletId: walletID,
		Amount:   amount,
	}

	resp, err := c.BusinessService.RequestPayout(ctx, req)
	if err != nil {
		return false, fmt.Errorf("failed to request payout: %w", err)
	}

	return resp.Success, nil
}

func (c *Client) ProcessCommission(ctx context.Context, commissionID string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, c.config.Timeout)
	defer cancel()

	req := &pb.ProcessCommissionRequest{
		CommissionId: commissionID,
	}

	resp, err := c.CommissionService.ProcessCommission(ctx, req)
	if err != nil {
		return false, fmt.Errorf("failed to process commission: %w", err)
	}

	return resp.Success, nil
}

func (c *Client) GetPendingCommissions(ctx context.Context) ([]*pb.Commission, error) {
	ctx, cancel := context.WithTimeout(ctx, c.config.Timeout)
	defer cancel()

	resp, err := c.CommissionService.GetPendingCommissions(ctx, &pb.GetPendingCommissionsRequest{})
	if err != nil {
		return nil, fmt.Errorf("failed to get pending commissions: %w", err)
	}

	return resp.Commissions, nil
}

func (c *Client) GetFailedCommissions(ctx context.Context) ([]*pb.Commission, error) {
	ctx, cancel := context.WithTimeout(ctx, c.config.Timeout)
	defer cancel()

	resp, err := c.CommissionService.GetFailedCommissions(ctx, &pb.GetFailedCommissionsRequest{})
	if err != nil {
		return nil, fmt.Errorf("failed to get failed commissions: %w", err)
	}

	return resp.Commissions, nil
}

func (c *Client) RetryFailedCommissions(ctx context.Context) (int32, int32, error) {
	ctx, cancel := context.WithTimeout(ctx, c.config.Timeout)
	defer cancel()

	resp, err := c.CommissionService.RetryFailedCommissions(ctx, &pb.RetryFailedCommissionsRequest{})
	if err != nil {
		return 0, 0, fmt.Errorf("failed to retry failed commissions: %w", err)
	}

	return resp.RetriedCount, resp.SuccessCount, nil
}

func (c *Client) TrackTransaction(ctx context.Context, tx *pb.Transaction) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, c.config.Timeout)
	defer cancel()

	req := &pb.TrackTransactionRequest{
		Transaction: tx,
	}

	resp, err := c.AnalyticsService.TrackTransaction(ctx, req)
	if err != nil {
		return false, fmt.Errorf("failed to track transaction: %w", err)
	}

	return resp.Success, nil
}

func (c *Client) GetCommissionHistory(ctx context.Context, businessID uint64, startDate, endDate time.Time) ([]*pb.CommissionEntry, error) {
	ctx, cancel := context.WithTimeout(ctx, c.config.Timeout)
	defer cancel()

	req := &pb.GetCommissionHistoryRequest{
		BusinessId: businessID,
		StartDate:  timestampFromTime(&startDate),
		EndDate:    timestampFromTime(&endDate),
	}

	resp, err := c.AnalyticsService.GetCommissionHistory(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get commission history: %w", err)
	}

	return resp.Entries, nil
}

func (c *Client) GetPayoutHistory(ctx context.Context, walletID string, startDate, endDate time.Time) ([]*pb.PayoutEntry, error) {
	ctx, cancel := context.WithTimeout(ctx, c.config.Timeout)
	defer cancel()

	req := &pb.GetPayoutHistoryRequest{
		WalletId:  walletID,
		StartDate: timestampFromTime(&startDate),
		EndDate:   timestampFromTime(&endDate),
	}

	resp, err := c.AnalyticsService.GetPayoutHistory(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get payout history: %w", err)
	}

	return resp.Entries, nil
}

func (c *Client) GenerateMonthlyReport(ctx context.Context, businessID uint64, year int32, month int32) (*pb.FinancialReport, error) {
	ctx, cancel := context.WithTimeout(ctx, c.config.Timeout)
	defer cancel()

	req := &pb.GenerateMonthlyReportRequest{
		BusinessId: businessID,
		Year:       year,
		Month:      month,
	}

	resp, err := c.FinancialReportService.GenerateMonthlyReport(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to generate monthly report: %w", err)
	}

	return resp.Report, nil
}

func (c *Client) GenerateCustomReport(ctx context.Context, businessID uint64, startDate, endDate time.Time) (*pb.FinancialReport, error) {
	ctx, cancel := context.WithTimeout(ctx, c.config.Timeout)
	defer cancel()

	req := &pb.GenerateCustomReportRequest{
		BusinessId: businessID,
		StartDate:  timestampFromTime(&startDate),
		EndDate:    timestampFromTime(&endDate),
	}

	resp, err := c.FinancialReportService.GenerateCustomReport(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to generate custom report: %w", err)
	}

	return resp.Report, nil
}

func (c *Client) GetReportByID(ctx context.Context, reportID string) (*pb.FinancialReport, error) {
	ctx, cancel := context.WithTimeout(ctx, c.config.Timeout)
	defer cancel()

	req := &pb.GetReportByIDRequest{
		ReportId: reportID,
	}

	resp, err := c.FinancialReportService.GetReportByID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get report: %w", err)
	}

	return resp.Report, nil
}

func (c *Client) ExportReport(ctx context.Context, reportID string, format string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(ctx, c.config.Timeout)
	defer cancel()

	req := &pb.ExportReportRequest{
		ReportId: reportID,
		Format:   format,
	}

	resp, err := c.FinancialReportService.ExportReport(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to export report: %w", err)
	}

	return resp.Data, nil
}

func (c *Client) BeginTransaction(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, c.config.Timeout)
	defer cancel()

	resp, err := c.SessionClient.BeginTransaction(ctx, &pb.BeginTransactionRequest{})
	if err != nil {
		return "", fmt.Errorf("failed to begin transaction: %w", err)
	}

	return resp.SessionId, nil
}

func (c *Client) CommitTransaction(ctx context.Context, sessionID string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, c.config.Timeout)
	defer cancel()

	resp, err := c.SessionClient.CommitTransaction(ctx, &pb.CommitTransactionRequest{
		SessionId: sessionID,
	})
	if err != nil {
		return false, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return resp.Success, nil
}

func (c *Client) RollbackTransaction(ctx context.Context, sessionID string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, c.config.Timeout)
	defer cancel()

	resp, err := c.SessionClient.RollbackTransaction(ctx, &pb.RollbackTransactionRequest{
		SessionId: sessionID,
	})
	if err != nil {
		return false, fmt.Errorf("failed to rollback transaction: %w", err)
	}

	return resp.Success, nil
}

func (c *Client) CreateTransaction(ctx context.Context, fromWalletID, toWalletID string, amount *pb.Money, txType, description string) (*pb.Transaction, error) {
	ctx, cancel := context.WithTimeout(ctx, c.config.Timeout)
	defer cancel()

	req := &pb.CreateTransactionRequest{
		FromWalletId: fromWalletID,
		ToWalletId:   toWalletID,
		Amount:       amount,
		Type:         txType,
		Description:  description,
	}

	resp, err := c.TransactionService.CreateTransaction(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to create transaction: %w", err)
	}

	return resp.Transaction, nil
}

func (c *Client) GetTransaction(ctx context.Context, transactionID string) (*pb.Transaction, error) {
	ctx, cancel := context.WithTimeout(ctx, c.config.Timeout)
	defer cancel()

	req := &pb.GetTransactionRequest{
		TransactionId: transactionID,
	}

	resp, err := c.TransactionService.GetTransaction(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction: %w", err)
	}

	return resp.Transaction, nil
}

func (c *Client) ProcessPayment(ctx context.Context, tx *pb.Transaction) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, c.config.Timeout)
	defer cancel()

	req := &pb.ProcessPaymentRequest{
		Transaction: tx,
	}

	resp, err := c.PaymentService.ProcessPayment(ctx, req)
	if err != nil {
		return "", fmt.Errorf("failed to process payment: %w", err)
	}

	return resp.ReferenceId, nil
}

func (c *Client) CreateBusinessWallet(ctx context.Context, businessID uint64, businessType, currency string) (*pb.BusinessWallet, error) {
	ctx, cancel := context.WithTimeout(ctx, c.config.Timeout)
	defer cancel()

	req := &pb.CreateBusinessWalletRequest{
		BusinessId:   businessID,
		BusinessType: businessType,
		Currency:     currency,
	}

	resp, err := c.BusinessService.CreateBusinessWallet(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to create business wallet: %w", err)
	}

	return resp.Wallet, nil
}

func (c *Client) CalculateCommission(ctx context.Context, tx *pb.Transaction, businessType string) (*pb.Commission, error) {
	ctx, cancel := context.WithTimeout(ctx, c.config.Timeout)
	defer cancel()

	req := &pb.CalculateCommissionRequest{
		Transaction:  tx,
		BusinessType: businessType,
	}

	resp, err := c.CommissionService.CalculateCommission(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate commission: %w", err)
	}

	return resp.Commission, nil
}

func (c *Client) GenerateBusinessReport(ctx context.Context, businessID uint64, startDate, endDate time.Time) (*pb.AnalyticsReport, error) {
	ctx, cancel := context.WithTimeout(ctx, c.config.Timeout)
	defer cancel()

	req := &pb.GenerateBusinessReportRequest{
		BusinessId: businessID,
		StartDate:  timestampFromTime(&startDate),
		EndDate:    timestampFromTime(&endDate),
	}

	resp, err := c.AnalyticsService.GenerateBusinessReport(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to generate business report: %w", err)
	}

	return resp.Report, nil
}

func (c *Client) GenerateDailyReport(ctx context.Context, businessID uint64, date time.Time) (*pb.FinancialReport, error) {
	ctx, cancel := context.WithTimeout(ctx, c.config.Timeout)
	defer cancel()

	req := &pb.GenerateDailyReportRequest{
		BusinessId: businessID,
		Date:       timestampFromTime(&date),
	}

	resp, err := c.FinancialReportService.GenerateDailyReport(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to generate daily report: %w", err)
	}

	return resp.Report, nil
}

func timestampFromTime(t *time.Time) *timestamppb.Timestamp {
	if t == nil {
		return nil
	}
	return timestamppb.New(*t)
}
