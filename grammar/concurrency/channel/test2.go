package main

type N int

func test2() {
	c := make(chan *N, 1)
	defer close(c)

	c <- nil

	var n N = 34

	c <- &n

	println(<-c)
	//println(<-c)

}
