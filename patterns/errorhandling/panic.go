package main

import (
	"log"
	"time"
)

// panicDemo 演示 goroutine 中 panic 的恢复：
// recover 必须在引发 panic 的同一个 goroutine 中 defer 才有效。
// （原文件入口是 main，并入本包后改名）
func panicDemo() {
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
