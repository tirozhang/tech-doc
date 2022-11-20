package _002_Add_Two_Numbers

import (
	"github.com/halfrost/LeetCode-Go/structures"
)

// ListNode define
type ListNode = structures.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{Val: 0}
	head := root
	tmpVal := 0
	for l1 != nil || l2 != nil || tmpVal != 0 {

		tmp := tmpVal
		if l1 != nil {
			tmp += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			tmp += l2.Val
			l2 = l2.Next
		}
		tmpVal = tmp / 10
		head.Next = &ListNode{
			Val: tmp % 10,
		}
		head = head.Next
	}
	return root.Next
}
