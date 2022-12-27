package main

import "fmt"

func main() {

	// 声明一个slice，此时output是空。不会分配内存
	var output []int

	for _, i := range output {
		fmt.Println("inside", i)
	}

	//不能对空切片进行读操作，会引起数组越界，导致panic
	//s := output[0]
	// panic: runtime error: index out of range [0] with length 0

	//但是往一个声明了的切片中append元素是被允许的， 不会panic
	output = append(output, 1)
}
