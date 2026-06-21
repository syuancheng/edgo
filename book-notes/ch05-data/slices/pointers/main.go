package main

func main() {
	s := []int{0, 1, 2, 3, 4}

	p := &s

	p0 := &s[0]
	p1 := &s[1]

	println(p, p0, p1)

	(*p)[0] += 100 // *[]int 不支持直接索引，要先解引用。
	println((*p)[0])

	// *p += 100
}
