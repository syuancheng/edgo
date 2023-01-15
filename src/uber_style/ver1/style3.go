package main

import "sync"

type SMap struct {
	mu sync.Mutex //不要使用嵌入字段形式

	data map[string]string
}

func NewSMap() *SMap {
	return &SMap{
		data: make(map[string]string),
	}
}

func (m *SMap) Get(k string) string {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.data[k]
}

func main() {

	mu1 := new(sync.Mutex)
	mu1.Lock()

	var mu sync.Mutex //recommend
	mu.Lock()

}
