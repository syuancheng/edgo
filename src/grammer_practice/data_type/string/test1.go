package main

import (
	"fmt"
	"strings"
)

func main() {
	var s1 = ""
	us := strings.ToUpper(s1)
	fmt.Println(us)

	var s2 string
	fmt.Println(s2 == "")
	us2 := strings.ToUpper(s2)
	fmt.Println(us2)

	s3 := "aB"
	us3 := strings.ToUpper(s3)
	fmt.Println(us3)
}
