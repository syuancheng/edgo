package main

import "fmt"

func main() {
	done := make(chan struct{})
	go func() {
		defer close(done)
		defer func() {
			if recovered := recover(); recovered != nil {
				fmt.Println("recovered:", recovered)
			}
		}()
		panic("goroutine failed")
	}()

	<-done
	fmt.Println("main continues")
}
