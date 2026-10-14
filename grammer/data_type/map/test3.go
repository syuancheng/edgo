package main

import "fmt"

func main() {
	m3 := map[int]string{
		1: "krystal",
		2: "jessica",
		3: "luna",
		4: "suli",
		5: "lisa",
	}
	//乱序
	for i := 0; i < 4; i++ {
		for k, v := range m3 {
			print(k, ":", v, " ")
		}
		println()
	}

	//for
	for k, v := range m3 {
		if k == 5 {
			delete(m3, k)
			m3[6] = "amber"
			m3[k] = "test"
		}
		println(k, ":", v, " ")

	}
	println(m3)
	fmt.Println(m3)
}
