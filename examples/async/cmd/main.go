package main

import (
	"context"
	"fmt"

	"github.com/syuancheng/edgo/examples/async"
)

func main() {
	task := async.New(context.Background(), func(context.Context) (string, error) {
		return "async result", nil
	})

	result, err := task.Wait(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
