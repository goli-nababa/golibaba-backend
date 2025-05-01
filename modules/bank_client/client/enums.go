package client

type BusinessType string

const (
	BusinessTypeHotel        BusinessType = "hotel"
	BusinessTypeAirline      BusinessType = "airline"
	BusinessTypeTravelAgency BusinessType = "travel_agency"
	BusinessTypeShip         BusinessType = "ship"
	BusinessTypeTrain        BusinessType = "train"
	BusinessTypeBus          BusinessType = "bus"
)

type TransactionStatus string

const (
	TransactionStatusPending    TransactionStatus = "pending"
	TransactionStatusProcessing TransactionStatus = "processing"
	TransactionStatusSuccess    TransactionStatus = "success"
	TransactionStatusFailed     TransactionStatus = "failed"
	TransactionStatusCancelled  TransactionStatus = "cancelled"
	TransactionStatusRefunded   TransactionStatus = "refunded"
)

type TransactionType string

const (
	TransactionTypeDeposit    TransactionType = "deposit"
	TransactionTypeWithdrawal TransactionType = "withdrawal"
	TransactionTypeTransfer   TransactionType = "transfer"
	TransactionTypePayment    TransactionType = "payment"
	TransactionTypeRefund     TransactionType = "refund"
	TransactionTypeCommission TransactionType = "commission"
)

type WalletStatus string

const (
	WalletStatusActive   WalletStatus = "active"
	WalletStatusInactive WalletStatus = "inactive"
	WalletStatusBlocked  WalletStatus = "blocked"
	WalletStatusLocked   WalletStatus = "locked"
)
