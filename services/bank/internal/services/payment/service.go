package payment

import (
	moneyDomain "bank_service/internal/common/types"
	domain2 "bank_service/internal/services/business/domain"
	CommissionServicePort "bank_service/internal/services/commission/port"
	"bank_service/internal/services/payment/domain"
	"bank_service/internal/services/payment/port"
	"bank_service/internal/services/payment/port/gateway"
	txDomain "bank_service/internal/services/transaction/domain"
	TransactionRepositoryPort "bank_service/internal/services/transaction/port"
	walletDomain "bank_service/internal/services/wallet/domain"
	WalletServicePort "bank_service/internal/services/wallet/port"
	"bank_service/pkg/cache"
	"context"
	"fmt"
	"time"
)

type paymentService struct {
	txRepo         TransactionRepositoryPort.TransactionRepository
	walletService  WalletServicePort.Service
	paymentGateway gateway.PaymentGateway
	commissionsvc  CommissionServicePort.Service
	cache          cache.Provider
}

func NewPaymentService(
	txRepo TransactionRepositoryPort.TransactionRepository,
	walletService WalletServicePort.Service,
	paymentGateway gateway.PaymentGateway,
	commissionsvc CommissionServicePort.Service,
	cache cache.Provider,
) port.Service {
	return &paymentService{
		txRepo:         txRepo,
		walletService:  walletService,
		paymentGateway: paymentGateway,
		commissionsvc:  commissionsvc,
		cache:          cache,
	}
}

func (s *paymentService) HandlePaymentCallback(ctx context.Context, callback *domain.PaymentCallback) (bool, string, error) {
	lockKey := fmt.Sprintf("payment_callback:%s", callback.Authority)
	acquired, err := s.cache.GetLock(ctx, lockKey, 30*time.Second)
	if err != nil || !acquired {
		return false, "", fmt.Errorf("failed to acquire lock: %w", err)
	}
	defer s.cache.ReleaseLock(ctx, lockKey)

	tx, err := s.txRepo.FindByReferenceID(ctx, callback.Authority)
	if err != nil {
		return false, "", fmt.Errorf("failed to find transaction: %w", err)
	}

	if tx.Status != txDomain.TransactionStatusPending {
		return false, "/payment/duplicate", nil
	}

	if callback.Status != "OK" {
		if err := tx.Fail("payment failed: " + callback.Status); err != nil {
			return false, "", err
		}
		if err := s.txRepo.Update(ctx, tx); err != nil {
			return false, "", fmt.Errorf("failed to update failed transaction: %w", err)
		}
		return false, "/payment/failed", nil
	}

	verified, err := s.paymentGateway.VerifyPayment(ctx, callback.Authority, tx.Amount.Amount)
	if err != nil {
		return false, "", fmt.Errorf("failed to verify payment: %w", err)
	}

	if !verified {
		if err := tx.Fail("payment verification failed"); err != nil {
			return false, "", err
		}
		if err := s.txRepo.Update(ctx, tx); err != nil {
			return false, "", fmt.Errorf("failed to update failed transaction: %w", err)
		}
		return false, "/payment/failed", nil
	}

	if err := s.updateWalletBalances(ctx, tx, false); err != nil {
		return false, "", fmt.Errorf("failed to update wallet balances: %w", err)
	}

	if err := tx.Process(); err != nil {
		return false, "", fmt.Errorf("failed to complete transaction: %w", err)
	}

	if err := tx.Complete(); err != nil {
		return false, "", fmt.Errorf("failed to complete transaction: %w", err)
	}

	if err := s.txRepo.Update(ctx, tx); err != nil {
		return false, "", fmt.Errorf("failed to update completed transaction: %w", err)
	}

	return true, "/payment/success", nil
}

func (s *paymentService) RefundPayment(ctx context.Context, paymentID string) error {
	lockKey := fmt.Sprintf("refund:%s", paymentID)
	acquired, err := s.cache.GetLock(ctx, lockKey, 30*time.Second)
	if err != nil || !acquired {
		return fmt.Errorf("failed to acquire refund lock: %w", err)
	}
	defer s.cache.ReleaseLock(ctx, lockKey)

	tx, err := s.txRepo.FindByID(ctx, txDomain.TransactionID(paymentID))
	if err != nil {
		return fmt.Errorf("failed to find transaction: %w", err)
	}

	if tx.Status != txDomain.TransactionStatusSuccess {
		return fmt.Errorf("transaction is not in success status")
	}

	if tx.ReferenceID != "" {
		if err := s.paymentGateway.RefundPayment(ctx, tx.ReferenceID, tx.Amount); err != nil {
			return fmt.Errorf("gateway refund failed: %w", err)
		}
	}

	refundTx, err := txDomain.NewTransaction(
		tx.ToWalletID,
		tx.FromWalletID,
		tx.Amount,
		txDomain.TransactionTypeRefund,
		fmt.Sprintf("Refund for transaction %s", tx.ID),
	)
	if err != nil {
		return fmt.Errorf("failed to create refund transaction: %w", err)
	}

	if err := s.txRepo.Create(ctx, refundTx); err != nil {
		return fmt.Errorf("failed to save refund transaction: %w", err)
	}

	if err := s.updateWalletBalances(ctx, refundTx, true); err != nil {
		return fmt.Errorf("failed to update wallet balances: %w", err)
	}

	tx.Status = txDomain.TransactionStatusRefunded
	return s.txRepo.Update(ctx, tx)
}

func (s *paymentService) ProcessCommission(ctx context.Context, amount *moneyDomain.Money, businessID uint64) error {
	lockKey := fmt.Sprintf("commission:%d", businessID)
	acquired, err := s.cache.GetLock(ctx, lockKey, 30*time.Second)
	if err != nil || !acquired {
		return fmt.Errorf("failed to acquire commission lock: %w", err)
	}
	defer s.cache.ReleaseLock(ctx, lockKey)

	tx, err := txDomain.NewTransaction(
		walletDomain.WalletID(fmt.Sprintf("business_%d", businessID)),
		domain2.CENTRAL_WALLET_ID,
		amount,
		txDomain.TransactionTypeCommission,
		fmt.Sprintf("Commission charge for business %d", businessID),
	)
	if err != nil {
		return fmt.Errorf("failed to create commission transaction: %w", err)
	}

	if err := s.txRepo.Create(ctx, tx); err != nil {
		return fmt.Errorf("failed to save commission transaction: %w", err)
	}

	if err := s.commissionsvc.ProcessCommission(ctx, string(tx.ID)); err != nil {
		return fmt.Errorf("failed to process commission: %w", err)
	}

	return nil
}

func (s *paymentService) ChargeWallet(ctx context.Context, walletID walletDomain.WalletID, amount *moneyDomain.Money) error {
	lockKey := fmt.Sprintf("wallet:%s", walletID)
	acquired, err := s.cache.GetLock(ctx, lockKey, 30*time.Second)
	if err != nil || !acquired {
		return fmt.Errorf("failed to acquire wallet lock: %w", err)
	}
	defer s.cache.ReleaseLock(ctx, lockKey)

	tx, err := txDomain.NewTransaction(
		domain2.CENTRAL_WALLET_ID, // Platform's wallet
		walletID,
		amount,
		txDomain.TransactionTypeDeposit,
		"Wallet charge",
	)
	if err != nil {
		return fmt.Errorf("failed to create deposit transaction: %w", err)
	}

	if err := s.txRepo.Create(ctx, tx); err != nil {
		return fmt.Errorf("failed to save deposit transaction: %w", err)
	}

	if err := s.walletService.Credit(ctx, walletID, amount); err != nil {
		return fmt.Errorf("failed to credit wallet: %w", err)
	}

	return nil
}

func (s *paymentService) WithdrawFromWallet(ctx context.Context, walletID walletDomain.WalletID, amount *moneyDomain.Money) error {
	lockKey := fmt.Sprintf("wallet:%s", walletID)
	acquired, err := s.cache.GetLock(ctx, lockKey, 30*time.Second)
	if err != nil || !acquired {
		return fmt.Errorf("failed to acquire wallet lock: %w", err)
	}
	defer s.cache.ReleaseLock(ctx, lockKey)

	tx, err := txDomain.NewTransaction(
		walletID,
		domain2.CENTRAL_WALLET_ID, // Platform's wallet
		amount,
		txDomain.TransactionTypeWithdrawal,
		"Wallet withdrawal",
	)
	if err != nil {
		return fmt.Errorf("failed to create withdrawal transaction: %w", err)
	}

	if err := s.txRepo.Create(ctx, tx); err != nil {
		return fmt.Errorf("failed to save withdrawal transaction: %w", err)
	}

	if err := s.walletService.Debit(ctx, walletID, amount); err != nil {
		return fmt.Errorf("failed to debit wallet: %w", err)
	}

	return nil
}

func (s *paymentService) GetTransactionHistory(ctx context.Context, walletID walletDomain.WalletID) ([]*txDomain.Transaction, error) {
	transactions, err := s.txRepo.FindByWallet(ctx, walletID, &TransactionRepositoryPort.TransactionFilter{
		WalletID: walletID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction history: %w", err)
	}

	return transactions, nil
}

func (s *paymentService) ProcessPayment(ctx context.Context, tx *txDomain.Transaction) (string, error) {
	lockKey := fmt.Sprintf("payment:%s", tx.ID)
	acquired, err := s.cache.GetLock(ctx, lockKey, 30*time.Second)
	if !acquired {
		return "", fmt.Errorf("concurrent payment in progress")
	}
	defer s.cache.ReleaseLock(ctx, lockKey)

	gatewayRef, err := s.paymentGateway.InitiatePayment(ctx, tx.Amount, tx.Metadata)
	if err != nil {
		return "", fmt.Errorf("gateway payment failed: %w", err)
	}

	tx.ReferenceID = gatewayRef
	if err := s.txRepo.Update(ctx, tx); err != nil {
		return "", err
	}

	return gatewayRef, nil
}

func (s *paymentService) ProcessRefund(ctx context.Context, tx *txDomain.Transaction) error {
	if tx.ReferenceID == "" {
		return fmt.Errorf("no gateway reference found for refund")
	}

	if err := s.paymentGateway.RefundPayment(ctx, tx.ReferenceID, tx.Amount); err != nil {
		return fmt.Errorf("gateway refund failed: %w", err)
	}

	if err := s.updateWalletBalances(ctx, tx, true); err != nil {
		return err
	}

	tx.Status = txDomain.TransactionStatusRefunded
	return s.txRepo.Update(ctx, tx)
}

func (s *paymentService) VerifyPayment(ctx context.Context, referenceID string) (bool, error) {
	return s.paymentGateway.VerifyPayment(ctx, referenceID, 0)
}

func (s *paymentService) updateWalletBalances(ctx context.Context, tx *txDomain.Transaction, isRefund bool) error {
	var fromID, toID = tx.FromWalletID, tx.ToWalletID
	if isRefund {
		fromID, toID = toID, fromID
	}

	if tx.Type == txDomain.TransactionTypePayment {
		if err := s.walletService.Credit(ctx, toID, tx.Amount); err != nil {
			if reverseErr := s.walletService.Credit(ctx, fromID, tx.Amount); reverseErr != nil {
				return fmt.Errorf("failed to reverse debit: %v (original error: %v)", reverseErr, err)
			}
			return err
		}
		return nil
	}

	if err := s.walletService.Debit(ctx, fromID, tx.Amount); err != nil {
		return err
	}

	if err := s.walletService.Credit(ctx, toID, tx.Amount); err != nil {
		if reverseErr := s.walletService.Credit(ctx, fromID, tx.Amount); reverseErr != nil {
			return fmt.Errorf("failed to reverse debit: %v (original error: %v)", reverseErr, err)
		}
		return err
	}

	return nil
}
