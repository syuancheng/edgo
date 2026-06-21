package main

import "fmt"

type profile struct {
	name   string
	scores []int
}

func main() {
	original := profile{name: "Gopher", scores: []int{80, 90}}
	copyValue := original

	copyValue.name = "Go"
	copyValue.scores[0] = 100 // 两个 slice 仍共享底层数组。

	fmt.Println("original:", original)
	fmt.Println("copy:", copyValue)

	deepCopy := original
	deepCopy.scores = append([]int(nil), original.scores...)
	deepCopy.scores[0] = 60
	fmt.Println("deep copy:", deepCopy)
	fmt.Println("original after deep copy:", original)
}
