package _101_Symmetric_Tree

import (
	"github.com/halfrost/LeetCode-Go/structures"
)

type TreeNode = structures.TreeNode

// 解法一
func isSymmetric(root *structures.TreeNode) bool {
	if root == nil {
		return true
	}

	return isMirror(root.Left, root.Right)
}

func isMirror(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return true
	}

	return left.Val == right.Val && isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
}

//解法二

func isSymmetric1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSameTree(invertTree(root.Left), root.Right)
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	} else {
		return false
	}
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	invertTree(root.Left)
	invertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}
