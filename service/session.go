package service

import (
	"sync"
)

// SessionStore is a struct to represent the in-memory session store
type SessionStore struct {
	sessions map[string]string
	mu       sync.RWMutex
}

// NewSessionStore creates a new instance of SessionStore
func NewSessionStore() *SessionStore {
	return &SessionStore{
		sessions: make(map[string]string),
	}
}

// Set stores a value in the session store
func (s *SessionStore) Set(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.sessions[key] = value
}

// Get retrieves a value from the session store
func (s *SessionStore) Get(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, ok := s.sessions[key]
	return val, ok
}

// Delete removes a session from the session store
func (s *SessionStore) Del(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.sessions, key)
}
