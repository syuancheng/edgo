package main

import "fmt"

func testChange(a ...int) {
	fmt.Println(a)
}

func main() {

	testChange(1, 2, 3)
	testChange(4)
}
