package main

import (
	"fmt"
	"sync"
)

// workerPoolDemo 演示固定容量的 worker pool：
// 固定数量的 worker 从任务通道取活干，是控制并发度和内存占用的标准姿势。
//
// 为什么不用"每个任务一个 goroutine"？
// 任务量不可控时会创建海量 goroutine，打爆下游连接数/内存。
// worker pool 把并发度钉死在 workers，任务在 channel 里排队。
func workerPoolDemo() {
	const (
		workers = 3
		nTasks  = 9
	)

	tasks := make(chan int, nTasks)
	results := make(chan string, nTasks)

	// 启动 worker：for-range 通道，通道关闭后自动退出
	var wg sync.WaitGroup
	for w := 1; w <= workers; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for t := range tasks {
				results <- fmt.Sprintf("worker %d 处理了任务 %d", id, t)
			}
		}(w)
	}

	// 投递任务，然后关闭表示"没有更多任务了"
	for t := 1; t <= nTasks; t++ {
		tasks <- t
	}
	close(tasks)

	// 等 worker 全部干完，再关 results，消费端才能安全结束
	wg.Wait()
	close(results)

	for r := range results {
		fmt.Println(r)
	}
}
