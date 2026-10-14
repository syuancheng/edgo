package main

import (
	"sync"
	"time"
)

func test11_1() {
	var wg sync.WaitGroup
	wg.Add(2)

	a := make(chan int)

	// receive
	go func() {
		defer wg.Done()
		for {
			select {
			case x, okk := <-a:
				if !okk {
					return
				}
				name := "a"
				println(name, x, okk)
			default:
				// 修复编译：select 的 default 分支不能带接收语句（`default _, ok := <-a` 是非法语法）。
				// 原意是非阻塞探测；default 分支只能做普通语句，这里短暂让出 CPU 避免忙等。
				time.Sleep(time.Millisecond)
			}
		}
	}()

	//send
	go func() {
		defer wg.Done()
		defer close(a)

		for i := 0; i < 10; i++ {
			select {
			case a <- i:
			}
		}
	}()

	wg.Wait()
}
