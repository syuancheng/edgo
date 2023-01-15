package main

import "fmt"

func main() {

	s1 := make([]int, 0, 10)

	fmt.Printf("s: %v\n", s1)

	s2 := make([]int, 2, 10)

	fmt.Printf("s: %v\n", s2)

	s3 := make([]int, 2, 2)

	fmt.Printf("s3|poiner: %p, v: %v\n", &s3, s3) // 0xc000096090, v: [0 0]
	s3 = append(s3, 3, 4)
	fmt.Printf("s3|poiner: %p, v: %v\n", &s3, s3) // 0xc000096090, v: [0 0 3]

	s4 := make([]int, 1, 2)
	fmt.Printf("s4|poiner: %p, v: %v\n", &s4, s4) //s4|poiner: 0xc0000160e0, v: [0]
	s4 = append(s4, 3)
	fmt.Printf("s4|poiner: %p, v: %v\n", &s4, s4) //s4|poiner: 0xc0000160e0, v: [0 3]
}
