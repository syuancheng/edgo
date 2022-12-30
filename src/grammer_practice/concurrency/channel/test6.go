package main

import "time"

func main() {
	done := make(chan struct{})
	c := make(chan int)

	go func() {
		defer close(done)

		for x := range c {
			println(x)
		}
	}()

	time.Sleep(time.Second * 5)

	c <- 1
	c <- 2
	c <- 3

	close(c)
	<-done
}
