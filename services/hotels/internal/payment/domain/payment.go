package domain

import (
	bookingDomain "hotels-service/internal/booking/domain"
	"time"

	"github.com/google/uuid"
)

type PaymentMetodType uint8

const (
	PaymentMetodTypeUnknown PaymentMetodType = iota
	PaymentMetodTypeCash
	PaymentMetodTypeCard
)

type StatusType uint8

const (
	StatusTypeUnknown StatusType = iota
	StatusTypeSuccessfully
	StatusTypeUnsuccessfully
)

type PaymentID = uuid.UUID
type Payment struct {
	ID             PaymentID
	BookingID      bookingDomain.BookingID
	amount         float64
	payment_method PaymentMetodType
	payment_date   time.Time
	status         StatusType
}
