package main

import (
	"fmt"
)

func ex_5_2_2() {
	x, y := 10, 20
	a := [...]*int{&x, &y}

	p := &a
	fmt.Printf("%T, %v\n", a, a)
	fmt.Printf("%T, %v\n", p, p)
}
