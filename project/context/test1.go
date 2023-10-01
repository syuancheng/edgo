package main
//通过取消函数WithCancel()传播
import (
	"context"
	"fmt"
	"time"
)

func someHandler() {

	parentCtx := context.Background()

	childCtx, cancel := context.WithCancel(parentCtx)
	go doSth(childCtx, "child [1]")
	go doSth(childCtx, "child [2]")

	//模拟程序运行 - Sleep 3秒
	time.Sleep(3 * time.Second)
	cancel()

	//等待取消信号能被子Context读到
	time.Sleep(2 * time.Second)
}

func doSth(ctx context.Context, name string) {
	var i = 1
	for {
		time.Sleep(time.Second * 1)
		select {
		case <-ctx.Done():
			fmt.Printf("%s done!\n", name)
			return
		default:
			fmt.Printf("%s had worked %d seconds ago\n", name, i)
		}
		i++
	}
}

func main() {
	fmt.Println("start.")
	someHandler()
	fmt.Println("end.")
}
