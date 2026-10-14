package main

import (
	"fmt"
	"sync"
)

// pipeline 三个阶段，全部用单向通道连接，可自由组合：
//
//	gen -> square -> merge（fan-in）
//	    \-> square /
//
// 每个阶段都是独立的 goroutine，靠通道关闭传递"结束"信号。

// gen 数据源阶段：把一组数字发送到输出通道
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

// square 处理阶段：读取输入，计算平方后写出
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}

// merge 汇聚阶段（fan-in）：把多个通道的结果合并到一个通道
func merge(chs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	for _, ch := range chs {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for n := range c {
				out <- n
			}
		}(ch)
	}

	// 单独一个 goroutine 负责收尾：全部输入读完后关闭输出
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// pipelineDemo 演示 fan-out / fan-in 流水线。
func pipelineDemo() {
	in := gen(1, 2, 3, 4)

	// fan-out：同一份数据分发给多个 square 阶段并发消费（工作=2 时相当于两路并行处理）
	s1, s2 := square(in), square(in)

	// fan-in：合并两路结果。每个数字只会被其中一路处理，输出仍是 4 个平方数
	for n := range merge(s1, s2) {
		fmt.Println(n)
	}
}
