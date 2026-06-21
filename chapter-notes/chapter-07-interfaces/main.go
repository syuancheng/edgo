package main

import "fmt"

type Stringer interface {
	String() string
}

type Tester interface {
	Stringer
	Test()
}

type Data struct {
	Value string
}

func (d Data) String() string {
	return d.Value
}

func (d *Data) Test() {
	d.Value += " tested"
}

type StringFunc func() string

func (f StringFunc) String() string {
	return f()
}

func printDynamicType(value any) {
	switch value := value.(type) {
	case nil:
		fmt.Println("nil")
	case fmt.Stringer:
		fmt.Println("stringer:", value.String())
	case int:
		fmt.Println("int:", value)
	default:
		fmt.Printf("unknown: %T\n", value)
	}
}

func main() {
	data := &Data{Value: "Go"}
	var tester Tester = data // 只有 *Data 的方法集包含 Test。
	tester.Test()
	fmt.Println(tester.String())

	var empty Stringer
	fmt.Println("nil interface:", empty == nil)

	var typedNil *Data
	empty = typedNil
	fmt.Println("interface containing typed nil:", empty == nil)

	printDynamicType(StringFunc(func() string { return "function adapter" }))
}
