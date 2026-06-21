package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	mu   sync.RWMutex
	data map[string]string
}

func NewSafeMap() *SafeMap {
	return &SafeMap{data: make(map[string]string)}
}

func (m *SafeMap) Set(key, value string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = value
}

func (m *SafeMap) Get(key string) (string, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	value, ok := m.data[key]
	return value, ok
}

func main() {
	values := NewSafeMap()
	values.Set("language", "Go")
	value, ok := values.Get("language")
	fmt.Println(value, ok)
}
