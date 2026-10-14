package main

import (
	"fmt"
)

type Education struct {
	school string
	grade  int32
	score  float32
}

type Student struct {
	name      string
	age       int
	height    *float32
	family    []*string
	education *Education
}

func main() {
	var h1 float32 = 1.85
	dad1 := "player"
	s1 := Student{
		name:   "krystal",
		age:    28,
		height: &h1,
		family: []*string{
			&dad1,
		},
		education: &Education{
			school: "whu",
			grade:  3,
			score:  3.7,
		},
	}

	var h2 float32 = 1.76
	dad2 := "musitian"
	s2 := Student{
		name:   "jessica",
		age:    30,
		height: &h2,
		family: []*string{
			&dad2,
		},
		education: &Education{
			school: "korean",
			grade:  5,
			score:  3.8,
		},
	}

	var h3 float32 = 1.85
	dad3 := "player"
	s3 := &Student{
		name:   "krystal",
		age:    28,
		height: &h3,
		family: []*string{
			&dad3,
		},
		education: &Education{
			school: "test",
			grade:  3,
			score:  3.7,
		},
	}

	var h4 float32 = 1.76
	dad4 := "musitian"
	s4 := &Student{
		name:   "jessica",
		age:    30,
		height: &h4,
		family: []*string{
			&dad4,
		},
		education: &Education{
			school: "japan",
			grade:  5,
			score:  3.8,
		},
	}

	/**
	s1 and s2 have different address, inside s1 and s2, the elements also have different address
	*/
	fmt.Println("define")
	fmt.Printf("s1 pointer: %p, target: %v\n", &s1, s1)
	fmt.Printf("s2 pointer: %p, target: %v\n", &s2, s2)
	fmt.Printf("s3 pointer: %p, target: %v\n", &s3, s3)
	fmt.Printf("s4 pointer: %p, target: %v\n", &s4, s4)

	//copy s1 to tmp, inside struct, the elements have the same address
	tmp1 := s1
	fmt.Println("tmp1 = s1")
	fmt.Printf("s1 pointer: %p, target: %v\n", &s1, s1)
	fmt.Printf("tmp1 pointer: %p, target: %v\n", &tmp1, tmp1)

	tmp2 := s3
	fmt.Println("tmp2 := s3")
	fmt.Printf("s3 pointer: %p, target: %v\n", &s3, s3)
	fmt.Printf("tmp2 pointer: %p, target: %v\n", &tmp2, tmp2)

	// u2.name = "jessica"
	// u2.age++
	// fmt.Printf("u1 pointer: %p, target: %v\n", &u1, u1) //u1 pointer: 0xc0000b6000, target: {krystal 28 0xc0000ae004}
	// fmt.Printf("u2 pointer: %p, target: %v\n", &u2, u2) //u2 pointer: 0xc0000b6020, target: {jessica 29 0xc0000ae004}

	// *u2.height = 1.67
	// fmt.Printf("u1 pointer: %p, target: %v, height: %f\n", &u1, u1, *u1.height)
	// fmt.Printf("u2 pointer: %p, target: %v, height: %f\n", &u2, u2, *u2.height)
	// /*
	// 	u1 pointer: 0xc00005e020, target: {krystal 28 0xc000016074}, height: 1.670000
	// 	u2 pointer: 0xc00005e040, target: {jessica 29 0xc000016074}, height: 1.670000
	// */

	// *u2.sex.gender = true
	// //u1 gender will be changed
	// fmt.Printf("u1 pointer: %p, target: %v, sex.gender: %v\n", &u1, u1, *u1.sex.gender)
	// fmt.Printf("u2 pointer: %p, target: %v, sex.gender: %v\n", &u2, u2, *u2.sex.gender)

	// //u2.family = append(u2.family, "qian")

	// t := "qian"
	// fmt.Printf("u2.family pointer: %p, value: %v\n", &u2.family, u2.family)
	// u2.family = append(u2.family, &t)
	// fmt.Printf("u2.family pointer: %p, value: %v\n", &u2.family, u2.family)
	// //u2.family[0] = &t
	// fmt.Printf("u1 pointer: %p, target: %v\n", &u1, u1)
	// fmt.Printf("u2 pointer: %p, target: %v\n", &u2, u2)

	// s1 := make([]string, 1, 2)
	// s2 := s1
	// fmt.Printf("s1 pointer: %p, target: %v\n", &s1, s1)
	// fmt.Printf("s2 pointer: %p, target: %v\n", &s2, s2)

	// s1 = append(s1, "a")
	// fmt.Printf("new s1 pointer: %p, target: %v\n", &s1, s1)
	// fmt.Printf("new s2 pointer: %p, target: %v\n", &s2, s2)

	// s1[0] = "b"
	// fmt.Printf("new s1 pointer: %p, target: %v\n", &s1, s1)
	// fmt.Printf("new s2 pointer: %p, target: %v\n", &s2, s2)
}
