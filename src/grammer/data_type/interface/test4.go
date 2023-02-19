package main

func main() {
	//匿名接口
	var t3 interface {
		string() string
	} = data{}

	n := node{
		data: t3,
	}
	println(n.data.string())
}
