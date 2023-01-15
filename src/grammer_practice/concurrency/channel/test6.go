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
		println("done")
	}()

	time.Sleep(time.Second * 5)

	c <- 1
	c <- 2
	c <- 3

	time.Sleep(time.Second * 5)
	close(c)
	<-done
}
