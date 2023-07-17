package main

import "fmt"

func test(n int) {
	fmt.Printf("pointer: %p, target: %v\n", &n, n)
}

func testPtr(n *int) {
	fmt.Printf("pointer: %p, target: %v\n", &n, n)
}

func testReturn(n *int) *int {
	fmt.Printf("pointer: %p, target: %v\n", &n, n)
	*n++
	return n
}

func testReturnInner() int {
	res := 56
	fmt.Printf("pointer: %p, target: %v\n", &res, res)

	return res
}

func main() {
	a := 100
	fmt.Printf("pointer: %p, target: %v\n", &a, a)
	test(a)
	/*
		a, b
		c, b
	*/

	b := 1
	p := &b
	fmt.Printf("pointer: %p, target: %v\n", &p, p)
	testPtr(p)
	/*
		a, b
		c, b
	*/

	c := 100
	pc := &c
	fmt.Printf("pointer: %p, target: %v\n", &pc, pc)
	cr := testReturn(pc)
	fmt.Printf("pointer: %p, target: %v\n", &cr, cr)
	/**
	    a, b
		c, b
		d, b
	*/
	res := testReturnInner()
	fmt.Printf("pointer: %p, target: %v\n", &res, res)
}
