package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	text := "Go语言"
	fmt.Println("bytes:", len(text))
	fmt.Println("runes:", utf8.RuneCountInString(text))
	fmt.Println("upper:", strings.ToUpper(text))

	for byteIndex, r := range text {
		fmt.Printf("index=%d rune=%q\n", byteIndex, r)
	}
}
