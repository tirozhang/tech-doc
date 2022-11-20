package _094_Binary_Tree_Inorder_Traversal

import (
	"github.com/halfrost/LeetCode-Go/structures"
)

// TreeNode define
type TreeNode = structures.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 二叉树的中序遍历
func inorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return []int{}
	}
	res = append(res, inorderTraversal(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)
	return res
}

func inorderTraversalStack(root *TreeNode) (res []int) {
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return
}
