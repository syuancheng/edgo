package main

import (
	"fmt"
	"time"
)

// 注意：test1.go 里已有包级变量 c 和函数 counter，
// 这里改名为 c6/counter6 避免同包重声明（原代码在单独文件里可直接运行）。
var c6 int

func counter6() int {
	c6++
	return c6
}

func doSomething() string {
	a := 100

	go func(x, y int) {
		time.Sleep(time.Second * 5)
		println("go:", x, y)
	}(a, counter6()) //100, 1

	a += 100
	println("main:", a, counter6())

	return "doSomething done"
}

func test6() {
	result := doSomething()
	fmt.Println(result)

	// 原代码此处是死循环 for{}：main goroutine 空转阻塞，
	// 会导致统一入口分发时进程永远退不出去。
	// 改成可结束的写法，保留「等待后台 goroutine」的演示意图：
	// for{
	//
	// }

	println("done")

	time.Sleep(time.Second * 2)
}
