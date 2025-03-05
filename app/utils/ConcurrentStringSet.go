package utils

import "sync"

type ConcurrentStringSet struct {
	mu    sync.RWMutex
	store map[string]struct{}
}

func NewConcurrentStringSet() *ConcurrentStringSet {
	return &ConcurrentStringSet{
		store: make(map[string]struct{}),
	}
}

// Add adds a string to the set if it's not already present
func (s *ConcurrentStringSet) Add(item string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.store[item] = struct{}{}
}

// Exists checks if a string exists in the set
func (s *ConcurrentStringSet) Exists(item string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, exists := s.store[item]
	return exists
}

// Remove removes a string from the set
func (s *ConcurrentStringSet) Remove(item string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.store, item)
}

// List returns all the elements in the set
func (s *ConcurrentStringSet) List() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var result []string
	for key := range s.store {
		result = append(result, key)
	}
	return result
}
