package main

import "fmt"

func main() {
	type user struct {
		name string
		age  int
	}

	value := &user{name: "Zhao", age: 12}
	value.age++ // 选择器会自动解引用一级结构体指针。

	pointerToPointer := &value
	(**pointerToPointer).name = "Cheng"
	fmt.Println(*value)
}
