package main

import (
	"context"
	"fmt"
	"time"

	"github.com/syuancheng/edgo/patterns/async"
)

func asyncTestFn() (string, error) {
	// 模拟耗时操作
	time.Sleep(200 * time.Millisecond)
	return "test async function", nil
}

func main() {
	ctx := context.Background()

	rsAsyncFunc := async.NewAsyncFuncWrapper(ctx, func() (interface{}, error) {
		// 修复：原代码写成 return asyncTestFn, nil（返回了函数本身而不是调用结果）
		return asyncTestFn()
	})

	// 修复：原代码拿到 rsAsyncFunc 后从未使用，导致编译失败；
	// 这里补上真正的等待与结果消费。
	result, err := rsAsyncFunc.WaitAndGetResult()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Printf("result: %v\n", result)

	// 再演示一个超时场景：1ms 超时，函数需要 200ms，必然走 ErrAsyncFuncTimeout
	timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()

	timeoutWrapper := async.NewAsyncFuncWrapper(timeoutCtx, func() (interface{}, error) {
		return asyncTestFn()
	})
	result2, err2 := timeoutWrapper.WaitAndGetResult()
	fmt.Printf("timeout demo: result=%v, err=%v\n", result2, err2)
}
