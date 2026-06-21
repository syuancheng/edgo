package calculator

import "errors"

var ErrDivideByZero = errors.New("division by zero")

func Add(left, right int) int {
	return left + right
}

func Divide(left, right int) (int, error) {
	if right == 0 {
		return 0, ErrDivideByZero
	}
	return left / right, nil
}
