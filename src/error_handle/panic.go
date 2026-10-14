package main

import (
	"log"
	"time"
)

func main() {
	log.Println("Main logic start...")

	go func() {
		// 必须在引发 panic 的同一个 goroutine 中 defer
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recovered in goroutine: %v", err)
			}
		}()
		panic("A goroutine crashed!")
	}()

	time.Sleep(time.Second) // 保证 goroutine 执行完毕
	log.Printf("Main goroutine is still alive.")
}

// 输出：// Main logic start...// Recovered in goroutine: A goroutine crashed!
