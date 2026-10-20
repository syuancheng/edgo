package main

import (
	"fmt"
	"time"
)

// 交替打印奇数和偶数
func main() {
	ch := make(chan struct{})

	//偶数
	go func() {
		for i := 0; i < 10; i++ {
			ch <- struct{}{}
			if i%2 == 0 {
				fmt.Printf("#2:%d\n", i)
			}
		}
	}()
	//奇数
	go func() {
		for i := 0; i < 10; i++ {
			<-ch
			if i%2 == 1 {
				fmt.Printf("#1:%d\n", i)
			}
		}
	}()

	select {
	case <-time.After(time.Second * 10):
	}
}
