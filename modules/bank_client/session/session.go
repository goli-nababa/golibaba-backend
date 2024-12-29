package session

import (
	"context"
	"fmt"
	"github.com/goli-nababa/golibaba-backend/modules/bank_service_client/client"

	"google.golang.org/grpc/metadata"
)

type Session struct {
	client    *client.Client
	sessionID string
	ctx       context.Context
	committed bool
}

func Begin(ctx context.Context, client *client.Client) (*Session, error) {
	sessionID, err := client.BeginTransaction(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}

	// Add session ID to context
	ctx = metadata.AppendToOutgoingContext(ctx, "session_id", sessionID)

	return &Session{
		client:    client,
		sessionID: sessionID,
		ctx:       ctx,
	}, nil
}

func (s *Session) Context() context.Context {
	return s.ctx
}

func (s *Session) Commit() error {
	if s.committed {
		return fmt.Errorf("session already committed")
	}

	success, err := s.client.CommitTransaction(s.ctx, s.sessionID)
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	if !success {
		return fmt.Errorf("commit transaction returned false")
	}

	s.committed = true
	return nil
}

func (s *Session) Rollback() error {
	if s.committed {
		return fmt.Errorf("cannot rollback already committed session")
	}

	success, err := s.client.RollbackTransaction(s.ctx, s.sessionID)
	if err != nil {
		return fmt.Errorf("failed to rollback transaction: %w", err)
	}

	if !success {
		return fmt.Errorf("rollback transaction returned false")
	}

	return nil
}

func (s *Session) MustCommit() {
	if err := s.Commit(); err != nil {
		panic(err)
	}
}

func (s *Session) MustRollback() {
	if err := s.Rollback(); err != nil {
		panic(err)
	}
}
