package main

import (
	"context"
	"fmt"
)

type TabProcessor interface {
	ResultCh(ctx context.Context) string
	Result() string
	Type() string
}

type AsyncTabProc struct {
	TabProcessor
}

func (at *AsyncTabProc) ResultCh(ctx context.Context) string {

	return "asynctabproc resultch"
}

type PersonalizedTabProc struct {
	*AsyncTabProc
	name string
	age  int
}

func (pt *PersonalizedTabProc) Result() string {
	return "result_personal"
}

func (pt *PersonalizedTabProc) Type() string {
	return "result_personal"
}

func main() {
	var ptp = &PersonalizedTabProc{
		name: "category tab",
		age:  12,
	}

	res := ptp.Result()
	rch := ptp.ResultCh(context.Background())
	tp := ptp.Type()

	fmt.Println(res)
	fmt.Println(rch)
	fmt.Println(tp)

	var asy = AsyncTabProc{
		TabProcessor: ptp,
	}

	asy.Result()
	asy.ResultCh(context.Background())
	asy.Type()

	fmt.Println("test")
}
