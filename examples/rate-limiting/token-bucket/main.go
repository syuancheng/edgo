package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	// 每秒生成 2 个令牌,桶容量 5(允许瞬时突发 5 个)
	limiter := rate.NewLimiter(2, 5)

	ctx := context.Background()

	for i := 1; i <= 10; i++ {
		// Wait 会阻塞直到拿到一个令牌(或 ctx 取消)
		if err := limiter.Wait(ctx); err != nil {
			fmt.Printf("req %d: wait err: %v\n", i, err)
			continue
		}
		fmt.Printf("req %d processed at %s\n", i, time.Now().Format("15:04:05.000"))
	}
}
