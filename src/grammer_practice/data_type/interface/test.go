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

	var d data

	//var t tester = d data does not implement tester

	var t tester = &d //只有*data实现了接口的所有数据集
	t.test()
	println(t.string())
	fmt.Println(t == nil)

	// interface variable
	var t1 tester
	fmt.Println(t1 == nil)

	var t2 tester = &d
	pp(t2) //超集转为子集

	var s stringer = t2 //超集转为子集
	println(s.string())

	//匿名接口
	var t3 interface {
		string() string
	} = data{}

	n := node{
		data: t3,
	}
	println(n.data.string())

	//类型转换 ok-idiom
	var m newData = 15
	var x interface{} = m

	if r, ok := x.(fmt.Stringer); ok {
		fmt.Println(r)
	}

	if n1, ok := x.(newData); ok {
		fmt.Println(n1)
	}

	//类型转换 switch
	var y interface{} = func(x int) string {
		return fmt.Sprintf("d: %d", x)
	}

	switch v := y.(type) { //局部变量v是类型转换后的结果
	case nil:
		println("nil")
	case *int:
		println(*v)
	case func(int) string:
		println(v(100))
	case fmt.Stringer:
		println(v)
	default:
		println("unknown")
	}

	//技巧
	var f fmt.Stringer = FuncString(func() string {
		return "hello interface world"
	})
	fmt.Println(f)
}
