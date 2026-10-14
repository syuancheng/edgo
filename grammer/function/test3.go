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

func testNew(x int) func() {
	println(&x)

	return func() {
		println(&x, x)
	}
}

func main() {
	//1. 直接执行
	func(s string) {
		println(s)
	}("hell0 syuan!")

	//2.赋值给变量
	add := func(x, y int) int {
		return x + y
	}
	println(add(1, 5))

	//3. 作为参数，testFn 是一个参数为函数的函数
	testFn(func() {
		println("hello testFn")
	})

	//4. 作为返回值
	fn := testReturnFn()
	println(fn(5, 8))

	//5.普通函数和匿名函数都可以作为结构体字段，或经通道传递
	testStruct()

	//6.闭包
	f := testNew(0x100)
	f()
}
