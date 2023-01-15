package main

import (
	"sync"
	"time"
)

func main() {
	//fatal error: concurrent map read and map write if don't use RW lock
	var lock sync.RWMutex
	m4 := make(map[string]int)
	go func() {
		for {
			lock.Lock()
			m4["a"] = 1
			lock.Unlock()
			time.Sleep(time.Microsecond)
		}
	}()

	go func() {
		for {
			lock.RLock()
			_ = m4["b"]
			lock.RUnlock()
			time.Sleep(time.Microsecond)
		}
	}()

	select {}
}
