package main

import "sync"

func main() {
	var wg sync.WaitGroup

	go func() {
		wg.Add(1)

		println("hi !")
	}()

	wg.Wait()
	println("exit !")
}
