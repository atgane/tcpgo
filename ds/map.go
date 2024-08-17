package ds

import "sync"

type Map[F comparable, T any] struct {
	mm map[F]T
	mu sync.RWMutex
}

func NewMap[F comparable, T any](initSize int) *Map[F, T] {
	m := &Map[F, T]{
		mm: make(map[F]T, initSize),
	}
	return m
}

func (m *Map[F, T]) Load(key F) (value T, ok bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	value, ok = m.mm[key]
	return value, ok
}

func (m *Map[F, T]) Store(key F, value T) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.mm[key] = value
}

func (m *Map[F, T]) Delete(key F) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.mm, key)
}

// ⚠️
func (m *Map[F, T]) Range(f func(key F, value T) bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for key, value := range m.mm {
		if !f(key, value) {
			break
		}
	}
}
