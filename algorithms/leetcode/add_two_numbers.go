package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(left, right *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	carry := 0

	for left != nil || right != nil || carry != 0 {
		if left != nil {
			carry += left.Val
			left = left.Next
		}
		if right != nil {
			carry += right.Val
			right = right.Next
		}

		tail.Next = &ListNode{Val: carry % 10}
		tail = tail.Next
		carry /= 10
	}
	return dummy.Next
}
