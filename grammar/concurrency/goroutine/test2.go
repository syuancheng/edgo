package main

import "time"

func test2() {
	exit := make(chan struct{}) // 创建无缓冲 channel；空结构体用于传递完成通知，无需携带数据。

	go func() {
		time.Sleep(time.Second * 3)
		println("goroutine done.")

		close(exit)
	}()

	println("main ...")
	<-exit // 没有可接收的值且 channel 未关闭时阻塞；子 goroutine 调用 close(exit) 后继续执行。
	println("main exit")
}
