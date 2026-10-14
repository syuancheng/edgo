package main

import (
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

// errgroupDemo 演示 errgroup：一组并发子任务，任何一个失败，Wait 就返回该错误。
//
// 工作中最常见的场景：并行调用下游（RPC/DB/缓存），任何一路失败就整体失败。
// errgroup 内部基于 sync.WaitGroup 封装，省去手工计数，
// 还可以用 g.SetLimit(n) 控制最大并发数。
func errgroupDemo() {
	var g errgroup.Group

	// g.SetLimit(2) // 如需限制并发数（信号量），打开这行

	for i := 1; i <= 3; i++ {
		// 注意：Go 1.22+ 的 for 循环变量每次迭代都是新变量，
		// 闭包可以直接捕获 i；旧版本需要 i := i 拷贝一份
		g.Go(func() error {
			time.Sleep(time.Duration(i*50) * time.Millisecond) // 模拟下游调用耗时
			if i == 2 {
				return fmt.Errorf("下游 %d 调用失败", i)
			}
			fmt.Println("task", i, "done")
			return nil
		})
	}

	// Wait 等全部结束；若其中有错误，返回最先触发的那一个
	if err := g.Wait(); err != nil {
		fmt.Println("errgroup 返回错误:", err)
	}
}
