package main

import "fmt"

func main() {
	var zero [2]int
	partial := [4]int{2, 3: 10}
	inferred := [...]int{1, 2}

	fmt.Println("zero:", zero)
	fmt.Println("partial:", partial)
	fmt.Printf("inferred type: %T, value: %v\n", inferred, inferred)
}
