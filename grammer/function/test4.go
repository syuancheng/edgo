package main

func testBasic(x int) func() {
	println(&x)

	return func() {
		println(&x, x) //返回的匿名函数中直接用到了外部环境的变量
	}
}

func test1() []func() {
	var s []func()

	for i := 0; i < 2; i++ {
		s = append(s, func() {
			println(&i, i)
		})
	}

	return s
}

func test2() []func() {
	var s []func()

	for i := 0; i < 2; i++ {
		x := i
		s = append(s, func() {
			println(&x, x)
		})
	}

	return s
}

func main() {
	//6.闭包
	f := testBasic(0x100)
	f()

	//闭包1
	for _, f := range test1() {
		f()
	}

	//闭包2
	for _, f := range test2() {
		f()
	}
}
