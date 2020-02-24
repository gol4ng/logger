package handler

import (
	"sync"

	"github.com/gol4ng/logger"
)

// Memory will store Entry to slice entries
// mainly develop for testing purpose in order to do some entries assertion
type Memory struct {
	mu      sync.Mutex
	entries []logger.Entry
}

func (s *Memory) lock() func() {
	s.mu.Lock()
	return func() { s.mu.Unlock() }
}

// Handle act as logger.HandlerInterface
func (s *Memory) Handle(entry logger.Entry) error {
	defer s.lock()()
	s.entries = append(s.entries, entry)
	return nil
}

// CleanEntries will reset the in memory entries list
func (s *Memory) CleanEntries() {
	defer s.lock()()
	s.entries = []logger.Entry{}
}

// GetEntries will return the in memory entries list
func (s *Memory) GetEntries() []logger.Entry {
	defer s.lock()()
	return s.entries
}

// GetAndCleanEntries will return and clean the in memory entries list
func (s *Memory) GetAndCleanEntries() []logger.Entry {
	defer s.lock()()
	entries := s.entries
	s.entries = []logger.Entry{}
	return entries
}

// NewMemory init a new Memory handler
func NewMemory() *Memory {
	return &Memory{}
}
