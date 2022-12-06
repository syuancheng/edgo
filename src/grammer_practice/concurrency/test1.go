package main

import "time"

var c int

func counter() int {
	c++
	return c
}

func main() {
	a := 100

	go func(x, y int) {
		time.Sleep(time.Second)
		println("go:", x, y)
	}(a, counter()) //100, 1

	a += 100
	println("main:", a, counter()) // 200, 2

	time.Sleep(time.Second * 3)
}
