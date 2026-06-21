package main

import (
	"context"
	"fmt"
)

type Processor interface {
	Result(context.Context) (string, error)
	Type() string
}

type BaseProcessor struct {
	name string
}

func (p BaseProcessor) Type() string {
	return p.name
}

type PersonalizedProcessor struct {
	BaseProcessor
}

func (p PersonalizedProcessor) Result(ctx context.Context) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
		return "personalized result", nil
	}
}

func main() {
	var processor Processor = PersonalizedProcessor{
		BaseProcessor: BaseProcessor{name: "personalized"},
	}
	result, err := processor.Result(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(processor.Type(), result)
}
