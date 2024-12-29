package migrations

import (
	"bank_service/internal/services/business/domain"
	walletDomain "bank_service/internal/services/wallet/domain"
	"bank_service/pkg/adapters/storage/postgres/model"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type JSONMap map[string]interface{}

func (m JSONMap) Value() (driver.Value, error) {
	if m == nil {
		return nil, nil
	}
	return json.Marshal(m)
}

func (m *JSONMap) Scan(value interface{}) error {
	if value == nil {
		*m = nil
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal JSONB value: %v", value)
	}

	result := make(JSONMap)
	err := json.Unmarshal(bytes, &result)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSONB value: %v", err)
	}

	*m = result
	return nil
}

type StatusChange struct {
	FromStatus string    `json:"from_status"`
	ToStatus   string    `json:"to_status"`
	Reason     string    `json:"reason"`
	ChangedAt  time.Time `json:"changed_at"`
}

func Migrate(db *gorm.DB) error {

	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Error; err != nil {
		return fmt.Errorf("failed to create uuid extension: %w", err)
	}

	if err := createEnumTypes(db); err != nil {
		return err
	}

	err := db.AutoMigrate(
		&model.WalletModel{},
		&model.BusinessWalletModel{},
		&model.TransactionModel{},
		&model.CommissionModel{},
		&model.FinancialReportModel{},
		&model.AnalyticsModel{},
	)
	if err != nil {
		return err
	}

	return createCentralWallets(db)
}

func createCentralWallets(db *gorm.DB) error {
	baseWallet := &model.WalletModel{
		ID:        domain.CENTRAL_WALLET_ID,
		UserID:    0, // System user
		Balance:   0,
		Currency:  "IRR",
		Status:    string(walletDomain.WalletStatusActive),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Version:   1,
	}
	var count int64
	db.Model(&model.WalletModel{}).Where("id = ?", domain.CENTRAL_WALLET_ID).Count(&count)

	if count != 0 {
		return nil
	}
	if err := db.Create(baseWallet).Error; err != nil {
		return fmt.Errorf("failed to create central base wallet: %w", err)
	}

	businessWallet := &model.BusinessWalletModel{
		WalletModel:    model.WalletModel{ID: domain.CENTRAL_WALLET_ID},
		BusinessID:     0,
		BusinessType:   string(domain.BusinessTypeHotel),
		CommissionRate: 0,
		PayoutSchedule: "never",
	}

	if err := db.Create(businessWallet).Error; err != nil {
		return fmt.Errorf("failed to create central business wallet: %w", err)
	}

	return nil
}

func createEnumTypes(db *gorm.DB) error {

	if err := db.Exec(`DO $$ 
        BEGIN 
            IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'business_type') THEN
                CREATE TYPE business_type AS ENUM (
                    'hotel',
                    'airline',
                    'travel_agency',
                    'ship',
                    'train',
                    'bus'
                );
            END IF;
        END 
        $$;`).Error; err != nil {
		return err
	}

	if err := db.Exec(`DO $$ 
		BEGIN 
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'transaction_status') THEN
				CREATE TYPE transaction_status AS ENUM (
					'pending', 
					'processing', 
					'success', 
					'failed', 
					'cancelled', 
					'refunded'
				);
			END IF;
		END 
		$$;`).Error; err != nil {
		return err
	}

	if err := db.Exec(`DO $$ 
		BEGIN 
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'transaction_type') THEN
				CREATE TYPE transaction_type AS ENUM (
					'deposit',
					'withdrawal',
					'transfer',
					'payment',
					'refund',
					'commission'
				);
			END IF;
		END 
		$$;`).Error; err != nil {
		return err
	}

	if err := db.Exec(`DO $$ 
		BEGIN 
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'wallet_status') THEN
				CREATE TYPE wallet_status AS ENUM (
					'active',
					'inactive',
					'blocked',
					'locked'
				);
			END IF;
		END 
		$$;`).Error; err != nil {
		return err
	}

	if err := db.Exec(`DO $$ 
		BEGIN 
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'business_type') THEN
				CREATE TYPE business_type AS ENUM (
					'hotel',
					'airline',
					'travel_agency',
					'ship',
					'train',
					'bus'
				);
			END IF;
		END 
		$$;`).Error; err != nil {
		return err
	}

	return nil
}
