package main

import (
	"fmt"
)

func main() {
	x, y := 10, 20
	a := [...]*int{&x, &y}

	p := &a
	fmt.Printf("%T, %v\n", a, a)
	fmt.Printf("%T, %v\n", p, p)
}
