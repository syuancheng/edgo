package main

import (
	"fmt"
	"time"
)

var c int

func counter() int {
	c++
	return c
}

func doSomething() string {
	a := 100

	go func(x, y int) {
		time.Sleep(time.Second)
		println("go:", x, y)
	}(a, counter()) //100, 1

	a += 100
	println("main:", a, counter())

	return "doSomething done"
}

func main() {
	result := doSomething()
	fmt.Println(result)

	println("done")

	time.Sleep(time.Second * 2)
}
