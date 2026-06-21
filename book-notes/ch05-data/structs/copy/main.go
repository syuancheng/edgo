package main

import "fmt"

type Student struct {
	Name   string
	Scores []int
}

func main() {
	original := Student{Name: "Gopher", Scores: []int{80, 90}}
	copyValue := original

	copyValue.Name = "Go"
	copyValue.Scores[0] = 100

	fmt.Println("original:", original)
	fmt.Println("copy:", copyValue)
}
