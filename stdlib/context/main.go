package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, completed chan<- int) {
	ticker := time.NewTicker(5 * time.Millisecond)
	defer ticker.Stop()

	for step := 1; ; step++ {
		select {
		case <-ctx.Done():
			fmt.Println("worker stopped:", ctx.Err())
			return
		case <-ticker.C:
			select {
			case completed <- step:
			case <-ctx.Done():
				return
			}
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 25*time.Millisecond)
	defer cancel()

	completed := make(chan int)
	go worker(ctx, completed)

	for {
		select {
		case step := <-completed:
			fmt.Println("completed step", step)
		case <-ctx.Done():
			fmt.Println("caller stopped:", ctx.Err())
			return
		}
	}
}
