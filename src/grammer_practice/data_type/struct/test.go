package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	type data struct {
		x int
		y map[int]int
	}

	//d1 := data{
	//	x: 1,
	//}
	//
	//d2 := data{
	//	x: 2,
	//}

	//println(d1 == d2)  不支持，因为map不支持==比较

	type user struct {
		name string
		age  int
	}

	p := &user{
		name: "zhao",
		age:  12,
	}

	p.age++
	p.name += " kuangyi"

	//p1 := &p
	//*p1.name = "zhao" 不支持

	// struct{}的使用
	exit := make(chan struct{})

	go func() {
		println("hello world")
		time.Sleep(time.Microsecond)
		exit <- struct{}{}
	}()

	<-exit
	println("end.")

	type attr struct {
		perm int
	}
	type mymap map[string]int
	type file struct {
		attr
		mymap
		name string
	}

	f := file{
		name: "test.dat",
		attr: attr{
			perm: 94949,
		},
		mymap: map[string]int{},
	}
	f.attr.perm = 14433
	f.perm = 4567
	f.mymap["siyuan"] = 23
	println(f.mymap)

	//tag
	type student struct {
		name string `昵称`
		sex  byte   `性别`
	}

	s := student{"Tom", 1}
	v := reflect.ValueOf(s)
	t := v.Type()

	for i, n := 0, t.NumField(); i < n; i++ {
		fmt.Printf("%s: %v\n", t.Field(i).Tag, v.Field(i))
	}

}
