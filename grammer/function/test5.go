package main

import (
	"encoding/json"
	"time"
)

type Persion struct {
	Name      string
	Age       int
	Education *Education
}

type Education struct {
	School string
	Rank   int
}

func main() {
	p := &Persion{
		Name: "Krystal",
		Age:  29,
		Education: &Education{
			School: "CJGU",
			Rank:   5,
		},
	}

	go func(p1 *Persion) {
		time.Sleep(40 * time.Millisecond)

		println("go: ", &p1)
		data, err := json.Marshal(p1)
		if err != nil {
			println(err)
		}
		println("data: ", string(data))
	}(p)

	p.Age = 100
	p.Name = "jessica"
	p.Education.School = "SUEO"

	time.Sleep(100 * time.Millisecond)

	println("main: ", &p, p)
}
