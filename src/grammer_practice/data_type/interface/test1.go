package main

import (
	"fmt"
)

type stringer interface {
	string() string
}

type tester interface {
	stringer
	test()
}

type data struct{}

func (*data) test() {}

func (data) string() string {
	return ""
}

func pp(s stringer) {
	println(s.string())
}

//匿名接口
type node struct {
	data interface {
		string() string
	}
}

//类型转换
type newData int

func (d newData) String() string {
	return fmt.Sprintf("newData: %d", d)
}

//技巧
type FuncString func() string

func (f FuncString) String() string {
	return f()
}

func main() {
	//0-basic
	var d data

	//var t tester = d data does not implement tester
	var t tester = &d //只有*data实现了接口的所有数据集
	t.test()
	println(t.string())
	fmt.Println(t == nil)
}
