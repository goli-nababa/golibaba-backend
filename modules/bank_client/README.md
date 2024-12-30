# راهنمای فرایندهای مالی قلی‌بابا

## مدل‌های پایه

### Money
```
message Money {
  int64 amount = 1;
  string currency = 2;
}
```

## سناریوهای اصلی

### 1. کیف پول مرکزی قلی‌بابا

کیف پول مرکزی با شناسه ثابت `CENTRAL_WALLET_ID` در سیستم تعریف شده که تمام تراکنش‌ها از طریق آن انجام می‌شود.

### 2. حالت‌های مختلف فرایند مالی

#### خرید بلیط:
1. کاربر درخواست خرید می‌دهد
2. مبلغ از کیف پول کاربر به کیف پول مرکزی منتقل می‌شود
3. بعد از تایید سفر:
    - سهم شرکت حمل و نقل پرداخت می‌شود
    - کمیسیون قلی‌بابا (۱٪) برداشت می‌شود
    - در صورت خرید از طریق آژانس، سهم آژانس پرداخت می‌شود

#### رزرو هتل:
1. مبلغ از کیف پول کاربر به کیف پول مرکزی منتقل می‌شود
2. بعد از پایان اقامت و تایید هتل:
    - سهم هتل پرداخت می‌شود
    - کمیسیون قلی‌بابا برداشت می‌شود
    - در صورت رزرو از طریق آژانس، سهم آژانس پرداخت می‌شود

#### خرید تور:
تور شامل سه فاکتور مجزاست:
1. پرواز/وسیله رفت
2. هتل
3. پرواز/وسیله برگشت

هر فاکتور باید جداگانه پردازش شود ولی خرید باید پکیجی انجام شود.

### 3. مدیریت تراکنش‌ها

از پکیج `pkg/transaction` برای مدیریت تراکنش‌ها استفاده می‌شود:

```
// شروع تراکنش
ctx := transaction.BeginTransaction(baseCtx)

// کامیت تراکنش
err := transaction.Commit(ctx)

// رولبک تراکنش
err := transaction.Rollback(ctx)
```

 
## فرایندهای عملیاتی

### 1. فرایند خرید بلیط:

```
func ProcessTicketPurchase(ctx context.Context, ticketData *TicketData) error {
    // شروع تراکنش
    txCtx := transaction.BeginTransaction(ctx)
    defer transaction.Rollback(txCtx)

    // بررسی موجودی کیف پول
    wallet, err := walletRepo.FindByID(txCtx, ticketData.UserWalletID)
    if err != nil {
        return fmt.Errorf("failed to get wallet: %w", err)
    }

    if wallet.Balance < ticketData.Amount {
        return errors.New("insufficient funds")
    }

    // انتقال به کیف پول مرکزی
    err = walletRepo.Transfer(
        txCtx,
        ticketData.UserWalletID,
        CENTRAL_WALLET_ID,
        ticketData.Amount,
    )
    if err != nil {
        return fmt.Errorf("transfer failed: %w", err)
    }

    // ثبت تراکنش
    tx := &Transaction{
        ID:          uuid.New().String(),
        FromWallet:  ticketData.UserWalletID,
        ToWallet:    CENTRAL_WALLET_ID,
        Amount:      ticketData.Amount,
        Type:        TransactionTypePayment,
        Status:      TransactionStatusPending,
        Description: "Ticket purchase",
    }
    
    err = txRepo.Create(txCtx, tx)
    if err != nil {
        return fmt.Errorf("failed to create transaction: %w", err)
    }

    return transaction.Commit(txCtx)
}
```

### 2. فرایند توزیع درآمد:

```
func DistributeRevenue(ctx context.Context, tripData *TripData) error {
    txCtx := transaction.BeginTransaction(ctx)
    defer transaction.Rollback(txCtx)

    // محاسبه سهم‌ها
    companyShare := calculateCompanyShare(tripData.Amount)
    gholibabaShare := calculateGholibabaCommission(tripData.Amount)
    agencyShare := calculateAgencyShare(tripData.Amount)

    // پرداخت به شرکت
    err = walletRepo.Transfer(
        txCtx,
        CENTRAL_WALLET_ID,
        tripData.CompanyWalletID,
        companyShare,
    )
    if err != nil {
        return fmt.Errorf("company payment failed: %w", err)
    }

    // برداشت کمیسیون قلی‌بابا
    err = walletRepo.Transfer(
        txCtx,
        CENTRAL_WALLET_ID,
        GHOLIBABA_WALLET_ID,
        gholibabaShare,
    )
    if err != nil {
        return fmt.Errorf("commission payment failed: %w", err)
    }

    // پرداخت به آژانس (در صورت وجود)
    if tripData.AgencyWalletID != "" {
        err = walletRepo.Transfer(
            txCtx,
            CENTRAL_WALLET_ID,
            tripData.AgencyWalletID,
            agencyShare,
        )
        if err != nil {
            return fmt.Errorf("agency payment failed: %w", err)
        }
    }

    return transaction.Commit(txCtx)
}
```

### 3. فرایند استرداد:

```
func ProcessRefund(ctx context.Context, refundData *RefundData) error {
    txCtx := transaction.BeginTransaction(ctx)
    defer transaction.Rollback(txCtx)

    // محاسبه مبلغ قابل استرداد با توجه به جریمه
    refundAmount := calculateRefundAmount(
        refundData.OriginalAmount,
        refundData.CancellationTime,
    )

    // برگشت پول به کیف پول کاربر
    err = walletRepo.Transfer(
        txCtx,
        CENTRAL_WALLET_ID,
        refundData.UserWalletID,
        refundAmount,
    )
    if err != nil {
        return fmt.Errorf("refund failed: %w", err)
    }

    // ثبت تراکنش استرداد
    tx := &Transaction{
        ID:          uuid.New().String(),
        FromWallet:  CENTRAL_WALLET_ID,
        ToWallet:    refundData.UserWalletID,
        Amount:      refundAmount,
        Type:        TransactionTypeRefund,
        Status:      TransactionStatusCompleted,
        Description: "Refund for cancellation",
    }

    err = txRepo.Create(txCtx, tx)
    if err != nil {
        return fmt.Errorf("failed to create refund transaction: %w", err)
    }

    return transaction.Commit(txCtx)
}
```

## نکات مهم

1. هر تراکنش باید در یک transaction انجام شود
4. در صورت خطا در هر مرحله، کل تراکنش باید rollback شود
