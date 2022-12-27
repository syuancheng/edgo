package main

import (
	"fmt"
	"reflect"
)

func main() {
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
