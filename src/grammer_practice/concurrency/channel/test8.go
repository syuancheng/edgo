package main

import "fmt"

func main() {
	c := make(chan int, 3)

	fmt.Println(cap(c))

	c <- 10
	c <- 20
	c <- 30

	close(c)

	for i := 0; i < cap(c)+1; i++ {
		x, ok := <-c
		println(i, ":", ok, x)
	}
}
