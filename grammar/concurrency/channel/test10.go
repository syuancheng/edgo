package main

func test10() {
	c := make(chan int, 2)

	var send chan<- int = c // send-only：只能发送
	var recv <-chan int = c // receive-only：只能接收

	// 单向通道的限制演示（以下操作均无法编译，故以注释保留原写法）：
	// <-send       // cannot receive from send-only channel
	// recv <- 1    // cannot send to receive-only channel
	// close(recv)  // cannot close receive-only channel
	// b := (chan int)(send) // 单向通道不能转回双向通道

	// 合法操作示例：
	send <- 1
	send <- 2
	println("recv:", <-recv, <-recv)
	close(c)
}
