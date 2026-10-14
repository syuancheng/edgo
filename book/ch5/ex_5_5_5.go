package main

import (
	"fmt"
	"reflect"
)

// tag
type student struct {
	name string `昵称`
	sex  byte   `性别`
}
type Fruit struct {
	Name  string `json:"name"`
	Color string `json:"color,omitempty"`
}

func ex_5_5_5() {

	s := student{"Tom", 1}
	v := reflect.ValueOf(s)
	t := v.Type()

	for i, n := 0, t.NumField(); i < n; i++ {
		fmt.Printf("%s: %v\n", t.Field(i).Tag, v.Field(i))
	}

	var f = Fruit{
		Name:  "apple",
		Color: "red",
	}

	// 修复编译：原文件此变量也叫 s，与上面的 student 变量重名，改名 jsonStr
	var jsonStr = `{"Name":"banana","Weight":100}`

	// 修复编译：f/jsonStr 声明后未使用，补充打印
	fmt.Printf("fruit: %+v\n", f)
	fmt.Println("raw json:", jsonStr)
}
