package types

import (
	"errors"
	"fmt"
	"math"
)

var (
	ErrInvalidAmount       = errors.New("invalid amount")
	ErrInvalidCurrency     = errors.New("invalid currency")
	ErrCurrencyMismatch    = errors.New("currency mismatch")
	ErrInvalidPrecision    = errors.New("invalid precision")
	ErrInvalidExchangeRate = errors.New("invalid exchange rate")
	ErrDivideByZero        = errors.New("divide by zero")
)

type Money struct {
	Amount    int64
	Currency  string // Currency code (e.g. IRR, USD)
	Precision int
}

func NewMoney(amount float64, currency string) (*Money, error) {
	if err := validateCurrency(currency); err != nil {
		return nil, err
	}

	precision := getCurrencyPrecision(currency)
	scaledAmount := int64(amount * math.Pow10(precision))

	if scaledAmount < 0 {
		return nil, ErrInvalidAmount
	}

	return &Money{
		Amount:    scaledAmount,
		Currency:  currency,
		Precision: precision,
	}, nil
}

func (m *Money) Add(other *Money) (*Money, error) {
	if err := m.validateSameCurrency(other); err != nil {
		return nil, err
	}

	return &Money{
		Amount:    m.Amount + other.Amount,
		Currency:  m.Currency,
		Precision: m.Precision,
	}, nil
}

func (m *Money) Subtract(other *Money) (*Money, error) {
	if err := m.validateSameCurrency(other); err != nil {
		return nil, err
	}

	result := m.Amount - other.Amount
	if result < 0 {
		return nil, ErrInvalidAmount
	}

	return &Money{
		Amount:    result,
		Currency:  m.Currency,
		Precision: m.Precision,
	}, nil
}

func (m *Money) Multiply(factor float64) (*Money, error) {
	if factor < 0 {
		return nil, ErrInvalidAmount
	}

	amount := float64(m.Amount) * factor
	if amount > math.MaxInt64 {
		return nil, ErrInvalidAmount
	}

	return &Money{
		Amount:    int64(amount),
		Currency:  m.Currency,
		Precision: m.Precision,
	}, nil
}

func (m *Money) Divide(factor float64) (*Money, error) {
	if factor == 0 {
		return nil, ErrDivideByZero
	}
	if factor < 0 {
		return nil, ErrInvalidAmount
	}

	amount := float64(m.Amount) / factor
	if amount < 1 {
		return nil, ErrInvalidPrecision
	}

	return &Money{
		Amount:    int64(amount),
		Currency:  m.Currency,
		Precision: m.Precision,
	}, nil
}

func (m *Money) Exchange(targetCurrency string, rate float64) (*Money, error) {
	if err := validateCurrency(targetCurrency); err != nil {
		return nil, err
	}
	if rate <= 0 {
		return nil, ErrInvalidExchangeRate
	}

	amount := float64(m.Amount) * rate
	precision := getCurrencyPrecision(targetCurrency)
	scaledAmount := amount * math.Pow10(precision-m.Precision)

	if scaledAmount > math.MaxInt64 {
		return nil, ErrInvalidAmount
	}

	return &Money{
		Amount:    int64(scaledAmount),
		Currency:  targetCurrency,
		Precision: precision,
	}, nil
}

func (m *Money) validateSameCurrency(other *Money) error {
	if m.Currency != other.Currency {
		return ErrCurrencyMismatch
	}
	return nil
}

func validateCurrency(currency string) error {
	if !IsValidCurrency(currency) {
		return ErrInvalidCurrency
	}
	return nil
}

func IsValidCurrency(currency string) bool {
	switch currency {
	case "IRR", "USD", "EUR":
		return true
	default:
		return false
	}
}

func getCurrencyPrecision(currency string) int {
	switch currency {
	case "IRR":
		return 0 // Iranian Rial has no decimal places
	case "USD", "EUR":
		return 2 // Most currencies use 2 decimal places
	default:
		return 2
	}
}

func (m *Money) String() string {
	factor := math.Pow10(m.Precision)
	amount := float64(m.Amount) / factor
	return fmt.Sprintf("%.2f %s", amount, m.Currency)
}

func (m *Money) Equals(other *Money) bool {
	if m.Currency != other.Currency {
		return false
	}
	return m.Amount == other.Amount
}

func (m *Money) GreaterThan(other *Money) bool {
	if m.Currency != other.Currency {
		return false
	}
	return m.Amount > other.Amount
}

func (m *Money) LessThan(other *Money) bool {
	if m.Currency != other.Currency {
		return false
	}
	return m.Amount < other.Amount
}

func (m *Money) IsPositive() bool {
	return m.Amount > 0
}
