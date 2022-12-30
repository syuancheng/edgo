package main

import (
	"fmt"
	"reflect"
)

//tag
type student struct {
	name string `昵称`
	sex  byte   `性别`
}
type Fruit struct {
	Name string `json:"name"`
	Color string `json:"color,omitempty"`
}

func main() {
	

	s := student{"Tom", 1}
	v := reflect.ValueOf(s)
	t := v.Type()

	for i, n := 0, t.NumField(); i < n; i++ {
		fmt.Printf("%s: %v\n", t.Field(i).Tag, v.Field(i))
	}

	var f = Fruit{
		Name: "apple",
		Color: "red",
	}

	var s = `{"Name":"banana","Weight":100}`

	
}
