package main

import (
	"context"
	"fmt"
	"time"
)

//通过超时控制函数传播

func someHandler2() {

	parentCtx := context.Background()

	childCtx, cancel := context.WithTimeout(parentCtx, 5*time.Second)

	go doSth2(childCtx, "child [1]")
	go doSth2(childCtx, "child [2]")

	//模拟程序运行 - Sleep 10秒
	time.Sleep(10 * time.Second)
	cancel()
	//等待取消信号能被子Context读到
	time.Sleep(2 * time.Second)

}

func doSth2(ctx context.Context, name string) {
	var i = 1
	for {
		//模拟具体work - Sleep 1秒
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Printf("%s done!\n", name)
			return
		default:
			fmt.Printf("%s had worked %d seconds \n", name, i)
		}
		i++
	}
}

func main() {
	fmt.Println("start.")
	someHandler2()
	fmt.Println("end.")

}
