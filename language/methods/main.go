package main

import "fmt"

type Counter int

func (c Counter) Value() int {
	return int(c)
}

func (c *Counter) Increment() {
	(*c)++
}

func main() {
	var count Counter = 2
	fmt.Println("before:", count.Value())
	count.Increment() // 可寻址变量会由编译器自动取地址。
	fmt.Println("after:", count.Value())
}
