package main

import "fmt"

func main() {

	s := make([]int, 0, 100)

	s1 := s[:2:4]

	s3 := s[:3:4]

	fmt.Printf("s3: %p, %v\n", &s3, s3)//

	s2 := append(s1, 1, 2, 3, 4, 5)//试着看看扩容与不扩容有什么不同

	fmt.Printf("s1: %p, %v\n", &s1, s1) // s1: 0x1400011a018, [0 0]
	fmt.Printf("s2: %p, %v\n", &s2, s2) // s2: 0x1400011a030, [0 0 1 2 3 4 5]
	fmt.Printf("s3: %p, %v\n", &s3, s3) //

	fmt.Printf("s : %v\n", s[:10])                           //s : [0 0 0 0 0 0 0 0 0 0]
	fmt.Printf("s1 cap: %d, s2 cap: %d\n", cap(s1), cap(s2)) //s1 cap: 4, s2 cap: 8

}
