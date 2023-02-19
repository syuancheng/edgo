package main

import "fmt"

func main() {
	//当append向切片中追加元素时， 在空间足够的情况下， 新元素将从x[len(x)]位置开始存放， append会生成一个新的切片
	//但不会修改原切片x
	x := make([]int, 0, 10)
	x = append(x, 1, 2, 3)
	y := append(x, 4)
	z := append(x, 5)
	fmt.Println(x) //1,2,3
	fmt.Println(y) //1,2,3,5
	fmt.Println(z) //1,2,3,5
}
