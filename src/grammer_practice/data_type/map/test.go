package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2

	m2 := map[int]struct {
		x int
	}{
		1: {x: 100},
		2: {x: 200},
	}
	fmt.Println(m, m2)

	//use ok-idiom 判断key是否存在
	if v, ok := m["d"]; ok {
		println(v) //返回零值
	}
	delete(m, "d") //即使key不存在也不会出错
}
