package main

import "fmt"

func ex_5_3_1() {
	var a []int
	fmt.Println(a == nil) //true
	fmt.Println(len(a))   //0
	for i := range a {
		// 修复：Println 不做格式化，含 %d 应使用 Printf
		fmt.Printf("test:%d\n", i)
	}
	//fmt.Println(a[0])//panic: runtime error: index out of range [1] with length 0

	b := []int{}
	fmt.Println(b == nil) //false
	fmt.Println(len(b))   //0
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

	d := make([]int, 2)
	fmt.Println(len(d))
	fmt.Println(cap(d))
	fmt.Println(d)

	d = append(d, 2)
	fmt.Println(d)
}
