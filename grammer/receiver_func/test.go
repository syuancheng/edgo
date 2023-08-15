package main

import "fmt"

type N int

func (n N) value() {
	n++
	fmt.Printf("v: %p, %v\n", &n, n)
}

func (n *N) pointer() {
	(*n)++
	fmt.Printf("p: %p, %v\n", n, *n)
}

type user struct {
	name string
}
type manager struct {
	name string
}

func (u user) toString() string {
	return u.name
}

func (m manager) toString() string {
	return m.name
}

func main() {
	var a N = 23
	p := &a
	fmt.Printf("a: %p, %v\n", &a, a)
	a.value() //copy
	a.pointer()
	fmt.Printf("a: %p, %v\n", &a, a)
	//a, p会自动转换
	p.value()
	p.pointer()

}
