package main

import (
	"fmt"
	"sync"
	"time"
)

func tryHandle(id int, semaphore chan struct{}) bool {
	select {
	case semaphore <- struct{}{}:
		defer func() { <-semaphore }()
		fmt.Printf("request %d handled\n", id)
		time.Sleep(10 * time.Millisecond) // 模拟占用资源，不用于 goroutine 同步。
		return true
	default:
		fmt.Printf("request %d rejected: at capacity\n", id)
		return false
	}
}

func main() {
	semaphore := make(chan struct{}, 3)
	var wg sync.WaitGroup
	for id := 1; id <= 6; id++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			tryHandle(id, semaphore)
		}(id)
	}
	wg.Wait()
}
