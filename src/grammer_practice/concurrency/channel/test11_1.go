package main

import (
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	a := make(chan int)

	// receive
	go func() {
		defer wg.Done()
		for {
			select {
			case x, okk := <-a:
				if !okk {
					return
				}
				name := "a"
				println(name, x, okk)
			default _, ok := <-a;
				println()
				println(ok)
			}
		}
	}()

	//send
	go func() {
		defer wg.Done()
		defer close(a)

		for i := 0; i < 10; i++ {
			select {
			case a <- i:
			}
		}
	}()

	wg.Wait()
}
