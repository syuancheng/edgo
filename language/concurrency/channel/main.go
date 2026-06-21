package main

import (
	"fmt"
	"sync"
)

func produce(out chan<- int) {
	defer close(out)
	for i := 1; i <= 5; i++ {
		out <- i
	}
}

func consume(in <-chan int) int {
	total := 0
	for value := range in {
		total += value
	}
	return total
}

func broadcastStart(workerCount int) {
	start := make(chan struct{})
	var wg sync.WaitGroup

	for id := 1; id <= workerCount; id++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			<-start
			fmt.Printf("worker %d started\n", id)
		}(id)
	}

	close(start) // 所有接收者都会立即收到关闭通知。
	wg.Wait()
}

func main() {
	values := make(chan int, 2)
	go produce(values)
	fmt.Println("sum:", consume(values))

	broadcastStart(3)
}
