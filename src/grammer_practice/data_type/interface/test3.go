package main

func main() {
	var d data

	var t2 tester = &d
	pp(t2) //超集转为子集

	var s stringer = t2 //超集转为子集
	println(s.string())
}
