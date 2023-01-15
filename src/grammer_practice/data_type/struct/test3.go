package main

import (
	"time"
)

func main() {
	// struct{}的使用
	exit := make(chan struct{})

	go func() {
		println("hello world")
		time.Sleep(time.Microsecond)
		exit <- struct{}{}
	}()

	<-exit
	println("end.")
}
