package main

import (
	"fmt"
)

func ex_5_5_1() {
	type data struct {
		x int
		y map[int]int
	}

	d1 := data{
		x: 1,
	}

	d2 := data{
		x: 2,
	}

	// 修复编译：原文件 d1/d2 声明后未使用，补充打印
	fmt.Printf("d1: %+v, d2: %+v\n", d1, d2)

	//println(d1 == d2) //不支持，因为map不支持==比较
}
