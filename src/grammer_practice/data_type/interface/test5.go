package main

import "fmt"

func main() {
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
}
