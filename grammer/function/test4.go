package main

func main() {
	x, y := 1, 2
	//defer println(x, y)
	//defer func(a, b int) {
	//	println(a, b)
	//}(x, y)

	defer func() {
		println(x, y)
	}()

	x += 100
	y += 100
}
