package main

/**
*  https://leetcode.com/problems/add-two-numbers/description/
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type listNode struct {
	Val  int
	Next *listNode
}

func addTwoNumbers(l1 *listNode, l2 *listNode) *listNode {
	carry, dummy := 0, &listNode{}
	for node := dummy; l1 != nil || l2 != nil || carry > 0; node=node.Next {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}

		//calculate the value of the current location
		node.Next = &listNode{carry%10, nil}

		//if the sum>10, then should pass 1 to next round calculation
		carry /= 10
	}
	return dummy.Next
}

func main() {

	dummy := new(listNode)

	println(dummy.Val)

	carry := 7

	carry /= 10 //0

	println(carry)

}
