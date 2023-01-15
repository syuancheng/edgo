package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	// close(c)

	go func() {
		x, ok := <-c
		if !ok {
			fmt.Println("not ok")
		} else {
			fmt.Println(x)
		}
	}()

	time.Sleep(time.Second * 5)

	close(c)

	time.Sleep(time.Second * 2)

}
