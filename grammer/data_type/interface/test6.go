package main

import "fmt"

func main() {
	//技巧
	var f fmt.Stringer = FuncString(func() string {
		return "hello interface world"
	})
	fmt.Println(f)
}
