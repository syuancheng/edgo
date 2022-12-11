package main

func main() {
	c := make(chan int, 3)

	c <- 1
	c <- 2

	println(<-c)
	println(<-c)
}
