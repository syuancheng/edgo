package main

import (
	"fmt"
	"sync"
	"time"
)

var limiter chan struct{}

func initLimiter(k int) {
	for i := 0; i < k; i++ {
		limiter <- struct{}{}
	}
}

func handleRequest(id int) {
	select {
	case <-limiter:
		defer func() {
			limiter <- struct{}{}
		}()
		fmt.Printf("请求 %d 开始处理，当前剩余名额: %d\n", id, len(limiter))
		time.Sleep(500 * time.Millisecond) // 模拟处理耗时
		fmt.Printf("请求 %d 完成，归还名额\n", id)
	default:
		fmt.Printf("请求 %d 被丢弃，并发已满\n", id)
	}
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 6; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			handleRequest(id)
		}(i)
	}

	wg.Wait()

}
