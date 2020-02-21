package testing

import (
	"sync"

	"github.com/gol4ng/logger"
)

type inMemoryEntriesStore struct {
	mu      sync.Mutex
	entries []logger.Entry
}

func (s *inMemoryEntriesStore) lock() func() {
	s.mu.Lock()
	return func() { s.mu.Unlock() }
}

// Handle act as logger.HandlerInterface
func (s *inMemoryEntriesStore) Handle(entry logger.Entry) error {
	defer s.lock()()
	s.entries = append(s.entries, entry)
	return nil
}

// CleanEntries will reset the in memory entries list
func (s *inMemoryEntriesStore) CleanEntries() {
	defer s.lock()()
	s.entries = []logger.Entry{}
}

// GetEntries will return the in memory entries list
func (s *inMemoryEntriesStore) GetEntries() []logger.Entry {
	defer s.lock()()
	return s.entries
}

// GetAndCleanEntries will return and clean the in memory entries list
func (s *inMemoryEntriesStore) GetAndCleanEntries() []logger.Entry {
	defer s.lock()()
	entries := s.entries
	s.entries = []logger.Entry{}
	return entries
}

// NewLogger will return a new inMemory Logger in order to do some assertion during test
func NewLogger() (*logger.Logger, *inMemoryEntriesStore) {
	store := &inMemoryEntriesStore{}
	return logger.NewLogger(store.Handle), store
}
