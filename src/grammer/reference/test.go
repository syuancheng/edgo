package main

import "fmt"

type Sex struct {
	gender *bool
}

type User struct {
	name   string
	age    int
	height *float32
	family []*string
	sex    Sex
}

func main() {
	var h float32 = 1.85
	var g bool = false

	f := make([]*string, 1, 3)

	u1 := User{"krystal", 28, &h, f, Sex{&g}}
	u2 := u1                                            //copy u1 to u2, u2 and u1 have different address
	fmt.Printf("u1 pointer: %p, target: %v\n", &u1, u1) //u1 pointer: 0xc0000b6000, target: {krystal 28 0xc0000ae004}
	fmt.Printf("u2 pointer: %p, target: %v\n", &u2, u2) //u2 pointer: 0xc0000b6020, target: {krystal 28 0xc0000ae004}

	u2.name = "jessica"
	u2.age++
	fmt.Printf("u1 pointer: %p, target: %v\n", &u1, u1) //u1 pointer: 0xc0000b6000, target: {krystal 28 0xc0000ae004}
	fmt.Printf("u2 pointer: %p, target: %v\n", &u2, u2) //u2 pointer: 0xc0000b6020, target: {jessica 29 0xc0000ae004}

	*u2.height = 1.67
	fmt.Printf("u1 pointer: %p, target: %v, height: %f\n", &u1, u1, *u1.height)
	fmt.Printf("u2 pointer: %p, target: %v, height: %f\n", &u2, u2, *u2.height)
	/*
		u1 pointer: 0xc00005e020, target: {krystal 28 0xc000016074}, height: 1.670000
		u2 pointer: 0xc00005e040, target: {jessica 29 0xc000016074}, height: 1.670000
	*/

	*u2.sex.gender = true
	//u1 gender will be changed
	fmt.Printf("u1 pointer: %p, target: %v, sex.gender: %v\n", &u1, u1, *u1.sex.gender) 
	fmt.Printf("u2 pointer: %p, target: %v, sex.gender: %v\n", &u2, u2, *u2.sex.gender)

	//u2.family = append(u2.family, "qian")

	t := "qian"
	fmt.Printf("u2.family pointer: %p, value: %v\n", &u2.family, u2.family)
	u2.family = append(u2.family, &t)
	fmt.Printf("u2.family pointer: %p, value: %v\n", &u2.family, u2.family)
	//u2.family[0] = &t
	fmt.Printf("u1 pointer: %p, target: %v\n", &u1, u1)
	fmt.Printf("u2 pointer: %p, target: %v\n", &u2, u2)

	s1 := make([]string, 1, 2)
	s2 := s1
	fmt.Printf("s1 pointer: %p, target: %v\n", &s1, s1)
	fmt.Printf("s2 pointer: %p, target: %v\n", &s2, s2)

	s1 = append(s1, "a")
	fmt.Printf("new s1 pointer: %p, target: %v\n", &s1, s1)
	fmt.Printf("new s2 pointer: %p, target: %v\n", &s2, s2)

	s1[0] = "b"
	fmt.Printf("new s1 pointer: %p, target: %v\n", &s1, s1)
	fmt.Printf("new s2 pointer: %p, target: %v\n", &s2, s2)
}
