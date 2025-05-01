package session

import (
	"context"
	"github.com/goli-nababa/golibaba-backend/modules/bank_service_client/client"
)

type Manager struct {
	client *client.Client
}

func NewManager(client *client.Client) *Manager {
	return &Manager{client: client}
}

// WithTransaction executes the provided function within a transaction session
func (m *Manager) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	session, err := Begin(ctx, m.client)
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			session.Rollback()
			panic(r)
		}
	}()

	if err := fn(session.Context()); err != nil {
		session.Rollback()
		return err
	}

	return session.Commit()
}

func WithTransactionE[T any](m *Manager, ctx context.Context, fn func(ctx context.Context) (T, error)) (T, error) {
	session, err := Begin(ctx, m.client)
	if err != nil {
		var zero T
		return zero, err
	}

	defer func() {
		if r := recover(); r != nil {
			session.Rollback()
			panic(r)
		}
	}()

	result, err := fn(session.Context())
	if err != nil {
		session.Rollback()
		var zero T
		return zero, err
	}

	if err := session.Commit(); err != nil {
		var zero T
		return zero, err
	}

	return result, nil
}

// SessionContext wraps a context with transaction session management
type SessionContext struct {
	context.Context
	session *Session
}

func NewSessionContext(ctx context.Context, client *client.Client) (*SessionContext, error) {
	session, err := Begin(ctx, client)
	if err != nil {
		return nil, err
	}

	return &SessionContext{
		Context: session.Context(),
		session: session,
	}, nil
}

func (c *SessionContext) Commit() error {
	return c.session.Commit()
}

func (c *SessionContext) Rollback() error {
	return c.session.Rollback()
}
