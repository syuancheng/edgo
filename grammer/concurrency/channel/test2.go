package main

type N int

func main() {
	c := make(chan *N, 1)
	defer close(c)

	c <- nil

	var n N = 34

	c <- &n

	println(<-c)
	//println(<-c)

}
