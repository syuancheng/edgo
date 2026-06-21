package main

func main() {
	// struct{}的使用
	exit := make(chan struct{})

	go func() {
		println("hello world")
		close(exit)
	}()

	<-exit
	println("end.")
}
