package client

import (
	pb "bank_service_client/proto/gen/go/bank/v1"
	"context"

	"time"
)

// WalletService provides wallet management operations
type WalletService interface {
	CreateWallet(ctx context.Context, ownerID uint64, walletType, currency string) (*pb.Wallet, error)
	GetWallet(ctx context.Context, walletID string) (*pb.Wallet, error)
	GetWalletsByUser(ctx context.Context, userID uint64) ([]*pb.Wallet, error)
	Credit(ctx context.Context, walletID string, amount *pb.Money) (*pb.Wallet, error)
	Debit(ctx context.Context, walletID string, amount *pb.Money) (*pb.Wallet, error)
	Transfer(ctx context.Context, fromWalletID, toWalletID string, amount *pb.Money) (*pb.TransferResponse, error)
	UpdateWalletStatus(ctx context.Context, walletID, status string) (*pb.Wallet, error)
}

// TransactionService provides transaction management operations
type TransactionService interface {
	CreateTransaction(ctx context.Context, fromWalletID, toWalletID string, amount *pb.Money, txType, description string) (*pb.Transaction, error)
	ProcessTransaction(ctx context.Context, transactionID string) (*pb.Transaction, error)
	CancelTransaction(ctx context.Context, transactionID, reason string) (*pb.Transaction, error)
	GetTransaction(ctx context.Context, transactionID string) (*pb.Transaction, error)
	ListTransactions(ctx context.Context, req *pb.ListTransactionsRequest) ([]*pb.Transaction, error)
}

// PaymentService provides payment processing operations
type PaymentService interface {
	ProcessPayment(ctx context.Context, tx *pb.Transaction) (string, error)
	ProcessRefund(ctx context.Context, tx *pb.Transaction) (bool, error)
	VerifyPayment(ctx context.Context, referenceID string) (bool, error)
	RefundPayment(ctx context.Context, paymentID string) (bool, error)
	ChargeWallet(ctx context.Context, walletID string, amount *pb.Money) (bool, error)
	WithdrawFromWallet(ctx context.Context, walletID string, amount *pb.Money) (bool, error)
	GetTransactionHistory(ctx context.Context, walletID string) ([]*pb.Transaction, error)
}

// BusinessService provides business wallet operations
type BusinessService interface {
	CreateBusinessWallet(ctx context.Context, businessID uint64, businessType, currency string) (*pb.BusinessWallet, error)
	GetBusinessWallet(ctx context.Context, walletID string) (*pb.BusinessWallet, error)
	UpdateBusinessWallet(ctx context.Context, wallet *pb.BusinessWallet) (*pb.BusinessWallet, error)
	SetPayoutSchedule(ctx context.Context, walletID, schedule string) (bool, error)
	RequestPayout(ctx context.Context, walletID string, amount *pb.Money) (bool, error)
	GetBusinessStats(ctx context.Context, businessID uint64, startDate, endDate time.Time) (*pb.BusinessStats, error)
}

// CommissionService provides commission management operations
type CommissionService interface {
	CalculateCommission(ctx context.Context, tx *pb.Transaction, businessType string) (*pb.Commission, error)
	ProcessCommission(ctx context.Context, commissionID string) (bool, error)
	GetCommission(ctx context.Context, commissionID string) (*pb.Commission, error)
	GetPendingCommissions(ctx context.Context) ([]*pb.Commission, error)
	GetFailedCommissions(ctx context.Context) ([]*pb.Commission, error)
	RetryFailedCommissions(ctx context.Context) (retried, success int32, err error)
}

// AnalyticsService provides analytics operations
type AnalyticsService interface {
	TrackTransaction(ctx context.Context, tx *pb.Transaction) (bool, error)
	GenerateBusinessReport(ctx context.Context, businessID uint64, startDate, endDate time.Time) (*pb.AnalyticsReport, error)
	GenerateBusinessStats(ctx context.Context, businessID uint64, startDate, endDate time.Time) (*pb.BusinessStats, error)
	GetCommissionHistory(ctx context.Context, businessID uint64, startDate, endDate time.Time) ([]*pb.CommissionEntry, error)
	GetPayoutHistory(ctx context.Context, walletID string, startDate, endDate time.Time) ([]*pb.PayoutEntry, error)
}

// FinancialReportService provides financial reporting operations
type FinancialReportService interface {
	GenerateDailyReport(ctx context.Context, businessID uint64, date time.Time) (*pb.FinancialReport, error)
	GenerateMonthlyReport(ctx context.Context, businessID uint64, year int32, month int32) (*pb.FinancialReport, error)
	GenerateCustomReport(ctx context.Context, businessID uint64, startDate, endDate time.Time) (*pb.FinancialReport, error)
	GetReportByID(ctx context.Context, reportID string) (*pb.FinancialReport, error)
	ExportReport(ctx context.Context, reportID string, format string) ([]byte, error)
}

// SessionService provides transaction session management
type SessionService interface {
	BeginTransaction(ctx context.Context) (string, error)
	CommitTransaction(ctx context.Context, sessionID string) (bool, error)
	RollbackTransaction(ctx context.Context, sessionID string) (bool, error)
}
