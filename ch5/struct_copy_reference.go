package main

import (
	"encoding/json"
	"fmt"

	"github.com/gogo/protobuf/proto"
)

type Company struct {
	Recommend Rcmd
	Ego       Ego
	Search    *Search
	Ads       *Ads
}

type Rcmd struct {
	Name   string
	Report Report
	Score  Score
}

type Search struct {
	Name   *string
	Report *Report
	Score  *Score
}

type Ads struct {
	Name   string
	Report Report
	Score  Score
}

type Ego struct {
	Name *string
}

type Report struct {
	Content *string
}

type Score struct {
	Score int32
}

func main() {

	mainObj := Company{
		Recommend: Rcmd{
			Name: "rcmd",
			Report: Report{
				Content: proto.String("rcmd content"),
			},
			Score: Score{
				Score: 11,
			},
		},
		Search: &Search{
			Name: proto.String("search"),
			Report: &Report{
				Content: proto.String("search content"),
			},
			Score: &Score{
				Score: 22,
			},
		},
		Ads: &Ads{
			Name: "ads",
			Report: Report{
				Content: proto.String("ads content"),
			},
			Score: Score{
				Score: 33,
			},
		},
		Ego: Ego{
			Name: proto.String("ego"),
		},
	}

	o1 := mainObj
	o2 := mainObj

	// o1.Recommend.Report.Content = proto.String("rcmd content adding")

	// o1.Ego.Name = proto.String("ego add") //请注意区分以下两种方式的区别
	*o1.Ego.Name = "ego add"

	fmt.Printf("main pointer: %p\n", mainObj.Ego.Name)
	fmt.Printf("o1 pointer: %p\n", o1.Ego.Name)

	fmt.Println("main:")
	printObj(mainObj)

	fmt.Println("o1:")
	printObj(o1)

	fmt.Println("o2:")
	printObj(o2)

}

func printObj(obj Company) {
	jsonDataIndented, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling to indented JSON:", err)
		return
	}

	fmt.Println(string(jsonDataIndented))
}
