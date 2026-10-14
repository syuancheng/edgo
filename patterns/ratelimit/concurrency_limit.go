package main

import (
	"fmt"
	"sync"
	"time"
)

var limiter chan struct{}

// initLimiter 向 limiter 中放入 k 个"名额"（token）。
// 注意：limiter 必须先 make 出带缓冲的 channel，否则向 nil channel
// 发送/接收会永久阻塞（select 里则表现为永远走 default 分支）。
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

func concurrency_limit() {
	// 修复：原代码没有初始化 limiter（nil channel），
	// 导致所有请求都直接走 default 被丢弃。
	// 正确做法：创建容量为 k 的 buffered channel 并放入 k 个名额。
	const maxConcurrency = 3
	limiter = make(chan struct{}, maxConcurrency)
	initLimiter(maxConcurrency)

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
