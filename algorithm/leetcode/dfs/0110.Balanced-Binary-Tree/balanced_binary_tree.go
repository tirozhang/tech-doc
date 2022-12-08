package _110_Balanced_Binary_Tree

import (
	"github.com/halfrost/LeetCode-Go/structures"
	"math"
)

type TreeNode = structures.TreeNode

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	isLeftBalanced, leftHeight := divideConquer(root.Left)
	isRightBalanced, rightHeight := divideConquer(root.Right)

	return isLeftBalanced && isRightBalanced && math.Abs(float64(leftHeight-rightHeight)) < 2
}

func divideConquer(node *TreeNode) (bool, int) {
	if node == nil {
		return true, 0
	}
	isLeftBalanced, leftHeight := divideConquer(node.Left)
	isRightBalanced, rightHeight := divideConquer(node.Right)
	isNodeBalanced := isLeftBalanced && isRightBalanced && abs(leftHeight-rightHeight) < 2
	return isNodeBalanced, max(leftHeight, rightHeight) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
