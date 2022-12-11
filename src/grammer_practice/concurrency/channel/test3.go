package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a, b chan int = make(chan int, 3), make(chan int)
	var c chan bool //nil chan

	println(a == b)
	println(c == nil)

	fmt.Printf("%p, %d\n", a, unsafe.Sizeof(a))
}
