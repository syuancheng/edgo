package main

import "fmt"

func main() {
	type point struct {
		X int
		Y int
	}

	left := point{X: 1, Y: 2}
	right := point{X: 1, Y: 2}
	fmt.Println(left == right)

	// 包含 slice、map 或 function 字段的结构体不可比较。
}
