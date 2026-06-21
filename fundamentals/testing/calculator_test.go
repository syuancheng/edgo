package calculator

import (
	"errors"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name        string
		left, right int
		want        int
	}{
		{name: "positive", left: 2, right: 3, want: 5},
		{name: "negative", left: -2, right: -3, want: -5},
		{name: "zero", left: 4, right: 0, want: 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.left, tt.right); got != tt.want {
				t.Fatalf("Add(%d, %d) = %d; want %d", tt.left, tt.right, got, tt.want)
			}
		})
	}
}

func TestDivideByZero(t *testing.T) {
	_, err := Divide(1, 0)
	if !errors.Is(err, ErrDivideByZero) {
		t.Fatalf("Divide(1, 0) error = %v; want %v", err, ErrDivideByZero)
	}
}
