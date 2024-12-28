package transaction

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"sync"
	"time"
)

type Session struct {
	ID        string
	Tx        *gorm.DB
	CreatedAt time.Time
	TTL       time.Duration
	Mu        sync.Mutex
}

type SessionStore struct {
	mu       sync.RWMutex
	sessions map[string]*Session
}

func NewSessionStore() *SessionStore {
	store := &SessionStore{
		sessions: make(map[string]*Session),
	}
	go store.startCleanupTicker()
	return store
}

func (s *SessionStore) startCleanupTicker() {
	ticker := time.NewTicker(5 * time.Minute)
	for range ticker.C {
		s.cleanup()
	}
}

func (s *SessionStore) cleanup() {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now()
	for id, session := range s.sessions {
		if now.Sub(session.CreatedAt) > session.TTL {
			session.Mu.Lock()
			session.Tx.Rollback()
			session.Mu.Unlock()
			delete(s.sessions, id)
		}
	}
}

func (s *SessionStore) Create(db *gorm.DB) (*Session, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	tx := db.Begin()
	if tx.Error != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", tx.Error)
	}

	session := &Session{
		ID:        uuid.New().String(),
		Tx:        tx,
		CreatedAt: time.Now(),
		TTL:       30 * time.Minute,
	}
	s.sessions[session.ID] = session

	return session, nil
}

func (s *SessionStore) Get(id string) (*Session, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	session, exists := s.sessions[id]
	if !exists {
		return nil, false
	}

	if time.Since(session.CreatedAt) > session.TTL {
		s.mu.Lock()
		session.Tx.Rollback()
		delete(s.sessions, id)
		s.mu.Unlock()
		return nil, false
	}

	return session, true
}

func (s *SessionStore) Remove(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.sessions, id)
}
