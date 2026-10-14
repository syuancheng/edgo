package main

import "sync"

func test4() {
	var wg sync.WaitGroup

	// 修复（vet 警告）：Add 必须在启动 goroutine 之前调用，
	// 否则 Wait 可能先于 Add 执行，导致计数为 0 直接放行（竞态）。
	wg.Add(1)

	go func() {
		defer wg.Done()

		println("hi !")
	}()

	wg.Wait()
	println("exit !")
}
