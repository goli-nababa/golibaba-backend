package bank_client_test

import (
	pb "bank_service_client/proto/gen/go/bank/v1"
	"bank_service_client/session"
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"time"
)

func (s *ClientTestSuite) TestTransactionalOperations() {
	// Create session manager
	manager := session.NewManager(s.client)

	// Test successful transaction
	err := manager.WithTransaction(context.Background(), func(ctx context.Context) error {
		// Create wallet
		wallet, err := s.client.CreateWallet(ctx, 1, "user", "IRR")
		require.NoError(s.T(), err)
		require.NotEmpty(s.T(), wallet.Id)

		// Credit wallet
		amount := &pb.Money{Amount: 1000, Currency: "IRR"}
		_, err = s.client.Credit(ctx, wallet.Id, amount)
		require.NoError(s.T(), err)

		// Create another wallet
		wallet2, err := s.client.CreateWallet(ctx, 2, "user", "IRR")
		require.NoError(s.T(), err)

		// Transfer money
		_, err = s.client.Transfer(ctx, wallet.Id, wallet2.Id, &pb.Money{Amount: 500, Currency: "IRR"})
		require.NoError(s.T(), err)

		// Verify balances within transaction
		w1, err := s.client.GetWallet(ctx, wallet.Id)
		require.NoError(s.T(), err)
		assert.Equal(s.T(), int64(500), w1.Balance.Amount)

		w2, err := s.client.GetWallet(ctx, wallet2.Id)
		require.NoError(s.T(), err)
		assert.Equal(s.T(), int64(500), w2.Balance.Amount)

		return nil
	})
	require.NoError(s.T(), err)

	// Test transaction rollback
	err = manager.WithTransaction(context.Background(), func(ctx context.Context) error {
		_, err := s.client.CreateWallet(ctx, 3, "user", "IRR")
		require.NoError(s.T(), err)

		// This should cause rollback
		return fmt.Errorf("simulated error")
	})
	require.Error(s.T(), err)

	// Test nested transactions (should use same session)
	err = manager.WithTransaction(context.Background(), func(ctx context.Context) error {
		wallet1, err := s.client.CreateWallet(ctx, 4, "user", "IRR")
		require.NoError(s.T(), err)
		_, err = s.client.Credit(ctx, wallet1.Id, &pb.Money{Amount: 500000000, Currency: "IRR"})
		require.NoError(s.T(), err)

		return manager.WithTransaction(ctx, func(ctx2 context.Context) error {
			wallet2, err := s.client.CreateWallet(ctx2, 5, "user", "IRR")
			require.NoError(s.T(), err)

			// Transfer between wallets created in different "transactions"
			_, err = s.client.Transfer(ctx2, wallet1.Id, wallet2.Id, &pb.Money{Amount: 100, Currency: "IRR"})
			return err
		})
	})
	require.NoError(s.T(), err)
}

func (s *ClientTestSuite) TestComplexBusinessOperations() {
	ctx := context.Background()
	manager := session.NewManager(s.client)

	// Complex business operation with multiple service calls
	businessWallet, err := session.WithTransactionE(manager, ctx, func(ctx context.Context) (*pb.BusinessWallet, error) {
		// Create business wallet
		bw, err := s.client.CreateBusinessWallet(ctx, 1, "hotel", "IRR")
		if err != nil {
			return nil, err
		}
		amount := &pb.Money{Amount: 1000000000000, Currency: "IRR"}
		_, err = s.client.Credit(ctx, bw.Id, amount)

		// Set commission rate and payout schedule
		bw.CommissionRate = 0.1
		bw.PayoutSchedule = "weekly"

		bw, err = s.client.UpdateBusinessWallet(ctx, bw)
		if err != nil {
			return nil, err
		}

		// Create customer wallet
		customerWallet, err := s.client.CreateWallet(ctx, 100, "user", "IRR")
		if err != nil {
			return nil, err
		}

		// Credit customer wallet
		_, err = s.client.Credit(ctx, customerWallet.Id, &pb.Money{Amount: 100000000, Currency: "IRR"})
		if err != nil {
			return nil, err
		}

		// Create and process payment
		tx, err := s.client.CreateTransaction(ctx, customerWallet.Id, bw.Id,
			&pb.Money{Amount: 500000, Currency: "IRR"}, "payment", "Test payment")
		if err != nil {
			return nil, err
		}

		_, err = s.client.ProcessPayment(ctx, tx)
		if err != nil {
			return nil, err
		}

		// Calculate and process commission
		commission, err := s.client.CalculateCommission(ctx, tx, "hotel")
		if err != nil {
			return nil, err
		}

		success, err := s.client.ProcessCommission(ctx, commission.Id)
		if err != nil {
			return nil, err
		}
		require.True(s.T(), success)

		//Generate analytics report
		report, err := s.client.GenerateBusinessReport(ctx, bw.BusinessId,
			time.Now().AddDate(0, 0, -7), time.Now())
		if err != nil {
			return nil, err
		}
		require.NotNil(s.T(), report)

		// Verify final state
		wallet, err := s.client.BusinessService.GetBusinessWallet(ctx, &pb.GetBusinessWalletRequest{
			WalletId: bw.Id,
		})

		return wallet.GetWallet(), err
	})

	require.NoError(s.T(), err)
	require.NotNil(s.T(), businessWallet)

	// Verify the final state reflects all operations
	assert.Equal(s.T(), 0.1, businessWallet.CommissionRate)
	assert.Equal(s.T(), "weekly", businessWallet.PayoutSchedule)
	assert.True(s.T(), businessWallet.Balance.Amount > 0)
}

func (s *ClientTestSuite) TestErrorHandling() {
	ctx := context.Background()
	manager := session.NewManager(s.client)

	// Test invalid currency
	_, err := session.WithTransactionE(manager, ctx, func(ctx context.Context) (*pb.Wallet, error) {
		return s.client.CreateWallet(ctx, 1, "user", "INVALID")
	})
	require.Error(s.T(), err)

	// Test insufficient funds
	_, err = session.WithTransactionE(manager, ctx, func(ctx context.Context) (*pb.TransferResponse, error) {
		wallet1, err := s.client.CreateWallet(ctx, 1, "user", "IRR")
		require.NoError(s.T(), err)

		wallet2, err := s.client.CreateWallet(ctx, 2, "user", "IRR")
		require.NoError(s.T(), err)

		// Try to transfer without funds
		return s.client.Transfer(ctx, wallet1.Id, wallet2.Id, &pb.Money{Amount: 1000, Currency: "IRR"})
	})
	require.Error(s.T(), err)

	// Test concurrent modifications
	wallet1, err := s.client.CreateWallet(context.Background(), 1, "user", "IRR")
	require.NoError(s.T(), err)

	// Start two concurrent transactions
	errChan := make(chan error, 2)
	for i := 0; i < 2; i++ {
		go func() {
			err := manager.WithTransaction(context.Background(), func(ctx context.Context) error {
				_, err := s.client.Credit(ctx, wallet1.Id, &pb.Money{Amount: 1000, Currency: "IRR"})
				return err
			})
			errChan <- err
		}()
	}

	// At least one should fail due to concurrent modification
	err1 := <-errChan
	err2 := <-errChan
	assert.True(s.T(), err1 != nil || err2 != nil)
}
