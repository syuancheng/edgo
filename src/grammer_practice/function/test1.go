package main

import "fmt"

func test(n int) {
	fmt.Printf("pointer: %p, target: %v\n", &n, n)
}

func testPtr(n *int) {
	fmt.Printf("pointer: %p, target: %v\n", &n, n)
}

func testChange(a ...int) {
	fmt.Println(a)
}

func testReturn(n *int) *int {
	fmt.Printf("pointer: %p, target: %v\n", &n, n)
	*n++
	return n
}

func testFn(f func()) {
	f()
}

func testReturnFn() func(int, int) int {
	return func(a, b int) int {
		return a * b
	}
}

func testStruct() {
	type calc struct {
		mul func(x, y int) int
	}

	x := calc{
		mul: func(x, y int) int {
			return x * y
		},
	}

	println(x.mul(6, 7))
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

	testChange(1, 2, 3)
	testChange(4)

	c := 100
	pc := &c
	fmt.Printf("pointer: %p, target: %v\n", &pc, pc)
	cr := testReturn(pc)
	fmt.Printf("pointer: %p, target: %v\n", &cr, cr)

	func(s string) {
		println(s)
	}("hell0 syuan!")

	add := func(x, y int) int {
		return x + y
	}
	println(add(1, 5))

	testFn(func() {
		println("hello testFn")
	})

	fn := testReturnFn()
	println(fn(5, 8))

	testStruct()
}
