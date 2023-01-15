package main

func main() {
	c := make(chan int, 2)

	var send chan<- int = c
	var recv <-chan int = c

	<-send //cannot receive from send-only channel
	recv<-

	close(recv)//cannot close receive-only channel recv (variable of type <-chan int)

	b := (chan int)(send) //cannot convert send (variable of type chan<- int) to chan int
}