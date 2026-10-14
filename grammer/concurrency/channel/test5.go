package main

import "time"

func main() {
	done := make(chan struct{})
	c := make(chan int)

	go func() {
		defer close(done)
		for {
			x, ok := <-c
			if !ok {
				return
			}
			println(x)
		}
	}()

	time.Sleep(time.Second * 3)

	c <- 1
	c <- 2
	c <- 3

	close(c)
	<-done
}
/**
result：
1
2
3
*/