package main

import "fmt"

type Number interface {
	~int | ~int64 | ~float64
}

func Sum[T Number](values []T) T {
	var total T
	for _, value := range values {
		total += value
	}
	return total
}

func Map[T, R any](values []T, transform func(T) R) []R {
	result := make([]R, len(values))
	for i, value := range values {
		result[i] = transform(value)
	}
	return result
}

func main() {
	fmt.Println(Sum([]int{1, 2, 3}))
	fmt.Println(Map([]int{1, 2, 3}, func(value int) string {
		return fmt.Sprintf("item-%d", value)
	}))
}
