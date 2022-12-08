package bfs_tpl

import "github.com/halfrost/LeetCode-Go/structures"

type TreeNode = structures.TreeNode

// LevelOrder 单队列实现
func LevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		var list []int
		length := len(queue)
		for i := 0; i < length; i++ {
			level := queue[0]
			list = append(list, level.Val)
			if level.Left != nil {
				queue = append(queue, level.Left)
			}
			if level.Right != nil {
				queue = append(queue, level.Right)
			}
			queue = queue[1:]
		}
		result = append(result, list)
	}
	return result
}

// LevelOrder2 双队列实现
func LevelOrder2(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		var list []int
		var nextQueue []*TreeNode
		for _, node := range queue {
			list = append(list, node.Val)
			if node.Left != nil {
				nextQueue = append(nextQueue, node.Left)
			}
			if node.Right != nil {
				nextQueue = append(nextQueue, node.Right)
			}
		}
		result = append(result, list)
		queue = nextQueue
	}
	return result
}

// LevelOrder3 dummynode实现
func LevelOrder3(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}

	var list []int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			if len(list) == 0 {
				break
			}
			result = append(result, list)
			queue = append(queue, nil)
			list = []int{}
			continue
		}
		list = append(list, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}

	}
	return result
}
