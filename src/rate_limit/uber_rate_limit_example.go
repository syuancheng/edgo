package main

import (
	"fmt"
	"time"

	"go.uber.org/ratelimit"
)

func main() {
	// 创建一个限流器，100 RPS（每秒100个请求）
	rl := ratelimit.New(2)

	for i := 0; i <= 10; i++ {
		rl.Take() // 没有令牌就阻塞等待，不会丢弃请求
		fmt.Printf("req %d processed at %s\n", i, time.Now().Format("15:04:05.000"))
	}
}
