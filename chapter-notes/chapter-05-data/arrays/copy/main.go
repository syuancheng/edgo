package main

import "fmt"

func mutateCopy(values [2]int) {
	values[0] = 100
	fmt.Println("inside function:", values)
}

func main() {
	original := [2]int{1, 2}
	copyValue := original
	copyValue[1] = 20

	mutateCopy(original)
	fmt.Println("original:", original)
	fmt.Println("copy:", copyValue)
}
