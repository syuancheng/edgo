package main

import (
	"time"
)

func ex_5_5_3() {
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
