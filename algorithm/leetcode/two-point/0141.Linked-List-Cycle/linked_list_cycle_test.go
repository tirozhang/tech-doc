package _141_Linked_List_Cycle

import (
	"github.com/halfrost/LeetCode-Go/structures"
	"testing"
)

func Test_Problem141(t *testing.T) {
	input := []int{3, 2, 0, -4}
	head := structures.Ints2List(input)
	t.Log(hasCycle(head))
	input = []int{3, 2, 0, 1}
	head = structures.Ints2List(input)
	t.Log(hasCycle(head))

}
