package _110_Balanced_Binary_Tree

import (
	"github.com/halfrost/LeetCode-Go/structures"
	"testing"
)

func Test_Problem(t *testing.T) {
	t.Log(isBalanced(structures.Ints2TreeNode([]int{1, 2, 3, structures.NULL, 5})))
	t.Log(isBalanced(structures.Ints2TreeNode([]int{1, 2, structures.NULL, 4, 5})))
}
