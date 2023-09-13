package main

import "fmt"

type Person struct {
	name   string
	school string
	score  []int32
}

func main() {
	p1 := Person{
		name:   "test",
		school: "ff",
		score: []int32{
			1, 2, 3,
		},
	}

	p2 := p1

	p2.name = "new_test"

	p2.score[0] = 90

	fmt.Printf("p1: v%\n", p1)
	fmt.Printf("p2: v%\n", p2)
}
