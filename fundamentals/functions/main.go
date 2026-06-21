package main

import "fmt"

func sum(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

func multiplier(factor int) func(int) int {
	return func(value int) int {
		return factor * value
	}
}

func main() {
	values := []int{1, 2, 3}
	fmt.Println("sum:", sum(values...))

	double := multiplier(2)
	fmt.Println("double:", double(21))

	x := 1
	defer func() { fmt.Println("deferred closure:", x) }()
	defer fmt.Println("deferred argument:", x)
	x = 2
}
