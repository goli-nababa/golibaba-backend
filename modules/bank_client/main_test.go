package bank_client_test

import (
	"bank_service_client/client"
	"context"

	"google.golang.org/grpc/metadata"
	"testing"
	"time"

	pb "bank_service_client/proto/gen/go/bank/v1"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ClientTestSuite struct {
	suite.Suite
	client *client.Client
}

func (s *ClientTestSuite) SetupSuite() {
	cfg := client.DefaultConfig()
	cfg.Host = "localhost"
	cfg.Port = 8082
	newClient, err := client.NewClient(cfg)
	assert.NoError(s.T(), err)
	s.client = newClient
}

func (s *ClientTestSuite) TearDownSuite() {
	assert.NoError(s.T(), s.client.Close())
}

func (s *ClientTestSuite) TestWalletService() {
	ctx := context.Background()
	beginResp, err := s.client.SessionClient.BeginTransaction(ctx, &pb.BeginTransactionRequest{})

	// Add session ID to context
	ctx = metadata.AppendToOutgoingContext(ctx, "session_id", beginResp.SessionId)
	// Defer rollback in case of error
	var committed bool
	defer func() {
		if !committed {
			s.client.SessionClient.RollbackTransaction(ctx, &pb.RollbackTransactionRequest{
				SessionId: beginResp.SessionId,
			})

		}
	}()

	// Create wallet
	wallet, err := s.client.CreateWallet(ctx, 1, "user", "IRR")
	assert.NoError(s.T(), err)
	assert.NotEmpty(s.T(), wallet.Id)

	// Get wallet
	retrievedWallet, err := s.client.GetWallet(ctx, wallet.Id)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), wallet, retrievedWallet)

	// Credit wallet
	updatedWallet, err := s.client.Credit(ctx, wallet.Id, &pb.Money{Amount: 1000, Currency: "IRR"})
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), int64(1000), updatedWallet.Balance.Amount)

	// Commit transaction
	_, err = s.client.SessionClient.CommitTransaction(ctx, &pb.CommitTransactionRequest{
		SessionId: beginResp.SessionId,
	})

	committed = true

}

func (s *ClientTestSuite) TestTransactionService() {
	// Create wallets
	fromWallet, err := s.client.CreateWallet(context.Background(), 1, "user", "IRR")
	assert.NoError(s.T(), err)
	toWallet, err := s.client.CreateWallet(context.Background(), 2, "user", "IRR")
	assert.NoError(s.T(), err)

	s.client.Credit(context.Background(), fromWallet.Id, &pb.Money{Amount: 10000000, Currency: "IRR"})

	// Create transaction
	tx, err := s.client.CreateTransaction(context.Background(), fromWallet.Id, toWallet.Id, &pb.Money{Amount: 500, Currency: "IRR"}, "transfer", "Test transfer")
	assert.NoError(s.T(), err)
	assert.NotEmpty(s.T(), tx.Id)

	// Get transaction
	retrievedTx, err := s.client.GetTransaction(context.Background(), tx.Id)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), tx, retrievedTx)

	transfer, err := s.client.WalletService.Transfer(context.Background(), &pb.TransferRequest{FromWalletId: fromWallet.Id, ToWalletId: toWallet.Id, Amount: &pb.Money{Amount: 500, Currency: "IRR"}})
	assert.NoError(s.T(), err)
	assert.NotEmpty(s.T(), transfer)

	transaction, err := s.client.TransactionService.ProcessTransaction(context.Background(), &pb.ProcessTransactionRequest{
		TransactionId: tx.Id,
	})
	assert.NoError(s.T(), err)
	assert.NotEmpty(s.T(), transaction)
}

func (s *ClientTestSuite) TestPaymentService() {
	// Create wallets
	_, err := s.client.CreateWallet(context.Background(), 1, "user", "IRR")
	assert.NoError(s.T(), err)
	toWallet, err := s.client.CreateWallet(context.Background(), 2, "user", "IRR")
	assert.NoError(s.T(), err)

	// Create transaction
	tx, err := s.client.CreateTransaction(context.Background(), "", toWallet.Id, &pb.Money{Amount: 50000, Currency: "IRR"}, "payment", "Test payment")
	assert.NoError(s.T(), err)

	// Process payment
	refID, err := s.client.ProcessPayment(context.Background(), tx)
	assert.NoError(s.T(), err)
	assert.NotEmpty(s.T(), refID)
}

func (s *ClientTestSuite) TestBusinessService() {
	// Create business wallet
	businessWallet, err := s.client.CreateBusinessWallet(context.Background(), 1, "hotel", "IRR")
	assert.NoError(s.T(), err)
	assert.NotEmpty(s.T(), businessWallet.Id)
	assert.Equal(s.T(), uint64(1), businessWallet.BusinessId)
}

func (s *ClientTestSuite) TestCommissionService() {
	// Create wallets
	fromWallet, err := s.client.CreateWallet(context.Background(), 1, "user", "IRR")
	assert.NoError(s.T(), err)
	toWallet, err := s.client.CreateWallet(context.Background(), 2, "user", "IRR")
	assert.NoError(s.T(), err)

	// Create transaction
	tx, err := s.client.CreateTransaction(context.Background(), fromWallet.Id, toWallet.Id, &pb.Money{Amount: 1000, Currency: "IRR"}, "payment", "Test commission")
	assert.NoError(s.T(), err)

	// Calculate commission
	commission, err := s.client.CalculateCommission(context.Background(), tx, "hotel")
	assert.NoError(s.T(), err)
	assert.NotEmpty(s.T(), commission.Id)
	assert.Equal(s.T(), "hotel", commission.BusinessType)
}

func (s *ClientTestSuite) TestAnalyticsService() {
	// Generate business report
	report, err := s.client.GenerateBusinessReport(context.Background(), 1, time.Now().AddDate(0, 0, -7), time.Now())
	assert.NoError(s.T(), err)
	assert.NotEmpty(s.T(), report.Id)
	assert.Equal(s.T(), uint64(1), report.BusinessId)
}

func (s *ClientTestSuite) TestFinancialReportService() {
	// Generate daily report
	report, err := s.client.GenerateDailyReport(context.Background(), 1, time.Now())
	assert.NoError(s.T(), err)
	assert.NotEmpty(s.T(), report.Id)
	assert.Equal(s.T(), uint64(1), report.BusinessId)
}

func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}
