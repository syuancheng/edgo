package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "REQID:37d8bbae008dce5f3c085fafd2713b00,BND:daily_discover_main,QUES:BayesStrategy_instance_1,RNK:rr_ranker_1:{SCORE:0.000000},SECKEY:dd_main_sec"
	fmt.Println(s1)

	arrary := strings.Split(s1, ",")
	//newArray := make([]string, 0, 2)
	for i, elem := range arrary {
		if strings.HasPrefix(elem, "BND") || strings.HasPrefix(elem, "QUES") {
			fmt.Println(elem)
		}
	}

	fmt.Println(arrary)
}
