package main

func testBasic(x int) func() {
	println(&x)

	return func() {
		println(&x, x) //返回的匿名函数中直接用到了外部环境的变量
	}
}

// makeClosuresWrong 演示闭包捕获循环变量的经典陷阱：
// 所有闭包共享同一个变量 i，最终都打印最后一次的值。
// （原文件里叫 test1，与 test1.go 撞名，改名）
func makeClosuresWrong() []func() {
	var s []func()

	for i := 0; i < 2; i++ {
		s = append(s, func() {
			println(&i, i)
		})
	}

	return s
}

// makeClosuresFixed 演示修复方式：每次循环拷贝一份 i，
// 每个闭包捕获自己那份变量。
// （原文件里叫 test2，与 test2.go 撞名，改名）
func makeClosuresFixed() []func() {
	var s []func()

	for i := 0; i < 2; i++ {
		x := i
		s = append(s, func() {
			println(&x, x)
		})
	}

	return s
}

func test4() {
	//6.闭包
	f := testBasic(0x100)
	f()

	//闭包1
	for _, f := range makeClosuresWrong() {
		f()
	}

	//闭包2
	for _, f := range makeClosuresFixed() {
		f()
	}
}
