package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2

	m3 := map[int]string{
		1: "a",
		2: "b",
	}

	m2 := map[int]struct {
		x int
	}{
		1: {x: 100},
		2: {x: 200},
	}
	fmt.Println(m, m2, m3)

	m4 := make(map[string]int, 5)
	fmt.Println(m4)
	m4["a"] = 1
	fmt.Println(m4)

	var m5 map[int]string

	fmt.Println(m5 == nil) //m5 is nil
	//m5[1] = "" //panic
	delete(m5, 1) //
	val, ok := m5[1]
	fmt.Println(val, ok) //"" false

}
