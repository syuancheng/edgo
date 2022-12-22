package main

import "fmt"

func main() {
	var a []int
	fmt.Println(a == nil) // a==nil
	fmt.Println(len(a))
	//fmt.Println(a[0])//panic: runtime error: index out of range [1] with length 0

	b := []int{}
	fmt.Println(b == nil) //b!=nil
	fmt.Println(len(b))
	//fmt.Println(b[0])//panic: runtime error: index out of range [1] with length 0

	c := []struct {
		Name string
	}{
		{
			Name: "a",
		},
		{
			Name: "b",
		},
		{
			Name: "c",
		},
	}

	for i := range c {
		//c[i].Name += "new"
		c = append(c, struct{ Name string }{Name: fmt.Sprintf("%d", i)})
	}

	for _, s := range c {
		fmt.Println(s)
	}
}
