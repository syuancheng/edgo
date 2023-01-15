package main

import (
	"context"
	"github.com/syuancheng/edgo/src/project_study/async"

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