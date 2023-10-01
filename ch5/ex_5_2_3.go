package main

import (
	"fmt"
)

func testCopyArr(x [2]int) {
	fmt.Printf("x: %p, %v\n", &x, x)
}

func main() {
	x, y := 10, 20
	a := [...]*int{&x, &y}

	p := &a
	fmt.Printf("A len=%d\n", len(a))
	fmt.Printf("%T, %v\n", a, a)
	fmt.Printf("%T, %v\n", p, p)

	b := [...]int{1, 2}
	println(&b, &b[0], &b[1])
	//0xc000092ee0 0xc000092ee0 0xc000092ee8  &b get the pointer of the first element
	pb := &b
	pb[1] += 10
	println(pb[1])

	fmt.Printf("x: %p, %v\n", &b, b)
	testCopyArr(b)
	/*
		a  b
		c  b
	*/

	var cb [2]int
	fmt.Printf("x: %p, %v\n", &b, b)
	cb = b
	cb[0] += 10
	fmt.Printf("cb: %p, %v\n", &cb, cb)
	fmt.Printf("b: %p, %v\n", &b, b)
	/** cb will copy b
	a   b
	c   b
	*/

	fmt.Println("====Slice=====")
	s1 := make([]int, 1, 3)
	s3 := []int{10, 20, 5: 30}
	fmt.Printf("s1: %d, cap: %d, slice: %v\n", len(s1), cap(s1), s1)
	//s1[0] = 1
	//s1[2] = 2  //panic
	fmt.Printf("s1: %d, cap: %d, slice: %v\n", len(s1), cap(s1), s1)
	s1 = append(s1, 2, 3, 4) //这里会扩容
	fmt.Printf("s1: %d, cap: %d, slice: %v\n", len(s1), cap(s1), s1)
	fmt.Printf("s3: %d, cap: %d, slice: %v\n", len(s3), cap(s3), s3)

	//两种创建方式的区别
	var s2 []int
	s4 := []int{}
	fmt.Printf("s2: %d, cap: %d, slice: %v\n", len(s2), cap(s2), s2)
	fmt.Printf("s4: %d, cap: %d, slice: %v\n", len(s4), cap(s4), s4)
	fmt.Println(s2 == nil) //true
	fmt.Println(s4 == nil) //false
	//fmt.Println(s4[0])//s2 and s4 don't read, only append

	s5 := []int{0, 1, 2, 3, 4}
	ps := &s5
	p0 := &s5[0]
	p1 := &s5[1]
	println(ps, p0, p1)
	*p0 += 5 //right operation
	fmt.Println(s5)

	//reslice
	d := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s6 := d[3:7]
	s7 := s6[1:3]
	for i := range s7 {
		s7[i] += 100
	}
	fmt.Println(d)
	fmt.Println(s6)
	fmt.Println(s7)

	s8 := make([]int, 0, 5)
	s9 := append(s8, 10)
	s10 := append(s9, 20, 30)
	fmt.Println(s8)
	fmt.Println(s9)
	fmt.Println(s10)

}
