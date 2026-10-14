package main

import (
	"context"

	lib "github.com/syuancheng/edgo/project/async"
)


func asyncTestFn() (string, error) {

	return "test async function", nil
}


func main() {
	ctx := context.Background()

	rsAsyncFunc := lib.NewAsyncFuncWrapper(ctx, func() (interface{}, error) {
		return asyncTestFn, nil
	})
}