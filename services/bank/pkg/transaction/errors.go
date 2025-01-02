package transaction

import "errors"

var (
	ErrSessionNotFound  = errors.New("transaction session not found")
	ErrSessionExpired   = errors.New("transaction session expired")
	ErrAlreadyCommitted = errors.New("transaction already committed")
	ErrInvalidSessionID = errors.New("invalid session ID")
)
