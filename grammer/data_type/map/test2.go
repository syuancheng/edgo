package main

func main() {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2

	//use ok-idiom 判断key是否存在
	v, ok := m["d"]
	println(v, ok) //0 false

	v, ok = m["a"]
	println(v, ok) //1 true

	delete(m, "d") //即使key不存在也不会出错
}
