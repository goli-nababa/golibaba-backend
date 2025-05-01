package events

import (
	txDomain "bank_service/internal/services/transaction/domain"
	"context"
)

type Event interface {
	PublishTransactionCreated(ctx context.Context, tx *txDomain.Transaction) error
	PublishTransactionStatusChanged(ctx context.Context, tx *txDomain.Transaction, oldStatus txDomain.TransactionStatus) error
	PublishEvent(ctx context.Context, event interface{}) error
}

func (p *EventPublisher) PublishTransactionCreated(ctx context.Context, tx *txDomain.Transaction) error {
	return nil
}

func (p *EventPublisher) PublishTransactionStatusChanged(ctx context.Context, tx *txDomain.Transaction, oldStatus txDomain.TransactionStatus) error {
	return nil
}

type EventPublisher struct {
	// ... (same as before)
}

func NewEventPublisher(url, exchange string) (Event, error) {
	// ... (same as before)
	return &EventPublisher{}, nil
}

func (p *EventPublisher) PublishEvent(ctx context.Context, event interface{}) error {
	// ... (same as before)
	return nil
}

func (p *EventPublisher) Close() error {
	// ... (same as before)
	return nil
}
