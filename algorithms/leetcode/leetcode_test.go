package leetcode

import (
	"reflect"
	"testing"
)

func TestMaxArea(t *testing.T) {
	if got := MaxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}); got != 49 {
		t.Fatalf("MaxArea() = %d; want 49", got)
	}
}

func TestMedianSortedArrays(t *testing.T) {
	tests := []struct {
		name        string
		left, right []int
		want        float64
		wantOK      bool
	}{
		{name: "odd", left: []int{1, 3}, right: []int{2}, want: 2, wantOK: true},
		{name: "even", left: []int{1, 2}, right: []int{3, 4}, want: 2.5, wantOK: true},
		{name: "negative", left: []int{-5, -1}, right: []int{-3}, want: -3, wantOK: true},
		{name: "duplicates", left: []int{1, 1}, right: []int{1, 1}, want: 1, wantOK: true},
		{name: "one value", right: []int{2}, want: 2, wantOK: true},
		{name: "one empty", right: []int{2, 3}, want: 2.5, wantOK: true},
		{name: "both empty", wantOK: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := MedianSortedArrays(tt.left, tt.right)
			if got != tt.want || ok != tt.wantOK {
				t.Fatalf("MedianSortedArrays() = (%v, %t); want (%v, %t)", got, ok, tt.want, tt.wantOK)
			}
		})
	}
}

func TestAddTwoNumbers(t *testing.T) {
	left := list(2, 4, 3)
	right := list(5, 6, 4)
	if got := values(AddTwoNumbers(left, right)); !reflect.DeepEqual(got, []int{7, 0, 8}) {
		t.Fatalf("AddTwoNumbers() = %v; want [7 0 8]", got)
	}
}

func list(values ...int) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for _, value := range values {
		tail.Next = &ListNode{Val: value}
		tail = tail.Next
	}
	return dummy.Next
}

func values(node *ListNode) []int {
	var result []int
	for ; node != nil; node = node.Next {
		result = append(result, node.Val)
	}
	return result
}
