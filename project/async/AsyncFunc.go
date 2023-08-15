package lib

import (
	"context"
	"errors"

)

var (
	ErrAsyncFuncTimeout = errors.New("timeout waiting")
)

type funcResult struct {
	result interface{}
	err    error
}

type AsyncFuncWrapper struct {
	ctx           context.Context
	fn            func() (interface{}, error)
	resultChan    chan funcResult
	result        *funcResult
}

func NewAsyncFuncWrapper(ctx context.Context, fn func() (interface{}, error)) *AsyncFuncWrapper {
	asyncFunc := &AsyncFuncWrapper{
		ctx: ctx,
		fn: fn,
		resultChan: make(chan funcResult, 1),
	}
	asyncFunc.exec()

	return asyncFunc

}

func (w *AsyncFuncWrapper) exec() {
    go func() {
		result, err := w.fn()
		w.resultChan <- funcResult{
			result: result, 
			err: err,
		}
	}()
}

func (w *AsyncFuncWrapper) WaitAndGetResult() (interface{}, error) {
	if w.result != nil {
		return w.result, w.result.err
	}
	select{
	case result := <- w.resultChan:
		w.result = &result
		return result.result, result.err
	case <- w.ctx.Done():
		return nil, ErrAsyncFuncTimeout
	}
}