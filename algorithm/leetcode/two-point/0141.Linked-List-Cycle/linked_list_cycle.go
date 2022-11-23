package _141_Linked_List_Cycle

import (
	"github.com/halfrost/LeetCode-Go/structures"
)

func hasCycle(head *structures.ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next

	for fast.Next != nil && fast != slow {
		slow = slow.Next
		fast = fast.Next
		if fast.Next != nil {
			fast = fast.Next
		}
	}
	return fast.Next != nil
}
