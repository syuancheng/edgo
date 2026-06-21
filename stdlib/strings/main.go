package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "REQID:37d8,BND:daily_discover_main,QUES:BayesStrategy,RNK:ranker"
	fields := strings.Split(input, ",")

	selected := make([]string, 0, 2)
	for _, field := range fields {
		if strings.HasPrefix(field, "BND:") || strings.HasPrefix(field, "QUES:") {
			selected = append(selected, field)
		}
	}

	fmt.Println(selected)
}
