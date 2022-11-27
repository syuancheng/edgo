package main

import "fmt"

func main() {
	var x int
	fmt.Printf("x = %d\n", x)

	var y = false
	fmt.Printf("y = %s\n", y)

	var a, b int
	fmt.Printf("a = %d\n", a)
	fmt.Printf("b = %d\n", b)

	var t, s = 100, "abc"
	fmt.Printf("t = %d\n", t)
	fmt.Printf("s = %s\n", s)

	var (
		g, h  int
		yy, z = 100, "abc"
	)
	fmt.Printf("g = %d\n", g)
	fmt.Printf("h = %d\n", h)
	fmt.Printf("yy = %d\n", yy)
	fmt.Printf("z = %s\n", z)
}
