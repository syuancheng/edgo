package main

import "fmt"

func testChange(a ...int) {
	fmt.Println(a)
}

func test2() {

	testChange(1, 2, 3)
	testChange(4)
}
