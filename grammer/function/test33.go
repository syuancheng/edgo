package main

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
