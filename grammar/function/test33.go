package main

// 本文件是 test3.go 的练习副本（原文件独立时可运行），
// 同包后函数名冲突，统一加 33 后缀。

func testFn33(f func()) {
	f()
}

func testReturnFn33() func(int, int) int {
	return func(a, b int) int {
		return a * b
	}
}

func testStruct33() {
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

func test33() {

	func(s string) {
		println(s)
	}("hell0 syuan!")

	add := func(x, y int) int {
		return x + y
	}
	println(add(1, 5))

	testFn33(func() {
		println("hello testFn")
	})

	fn := testReturnFn33()
	println(fn(5, 8))

	testStruct33()
}
