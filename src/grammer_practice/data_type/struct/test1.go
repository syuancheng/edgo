package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
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

	println(d1 == d2) //不支持，因为map不支持==比较
}
