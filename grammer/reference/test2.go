package main

import "fmt"

type aircraft struct {
	name      string
	category  string
	newly     bool
	basicInfo basicInfo
	useInfo   *useInfo
}

type basicInfo struct {
	height   int32
	capacity int32
	length   int32
	factor   string
}

type useInfo struct {
	duration int64
	repaired bool
}

func main() {
	var c = &aircraft{
		name:     "c919",
		category: "media",
		newly:    true,
		basicInfo: basicInfo{
			height:   10,
			capacity: 180,
			length:   42,
			factor:   "COMAIC",
		},
		useInfo: &useInfo{
			duration: 2400,
			repaired: true,
		},
	}

	a := c

	fmt.Printf("c|pointer: %p, target: %v\n", &c, c)
	fmt.Printf("a|pointer: %p, target: %v\n", &a, a)

	a.basicInfo.factor = "AIRBUS"
	fmt.Printf("c-info|pointer: %p, target: %v\n", &c.basicInfo, c.basicInfo)
	fmt.Printf("a-info|pointer: %p, target: %v\n", &a.basicInfo, a.basicInfo)

	a.useInfo.duration = 5000
	fmt.Printf("c-use|pointer: %p, target: %v\n", &c.useInfo, c.useInfo)
	fmt.Printf("a-use|pointer: %p, target: %v\n", &a.useInfo, a.useInfo)
}
