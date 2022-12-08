package dfs_tpl

import (
	"fmt"
	"github.com/halfrost/LeetCode-Go/structures"
	"strconv"
)

type TreeNode = structures.TreeNode

func findNodes(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	result = append(result, root.Val)
	result = append(result, findNodes(root.Left)...)
	result = append(result, findNodes(root.Right)...)
	return result
}

// 分治法
func findPaths(root *TreeNode) []string {
	result := make([]string, 0)
	if root == nil {
		return result
	}
	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}

	tmpLeft := findPaths(root.Left)
	for i := 0; i < len(tmpLeft); i++ {
		result = append(result, fmt.Sprintf("%d->%s", root.Val, tmpLeft[i]))
	}
	tmpRight := findPaths(root.Right)
	for i := 0; i < len(tmpRight); i++ {
		result = append(result, fmt.Sprintf("%d->%s", root.Val, tmpRight[i]))
	}
	return result
}

// 遍历法
func binaryTreePaths(root *TreeNode) []string {
	var res []string
	if root == nil {
		return res
	}
	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}

	path := []*TreeNode{root}
	findPaths2(root, path, &res)
	return res
}

func findPaths2(node *TreeNode, path []*TreeNode, result *[]string) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		pathString := strconv.Itoa(path[0].Val)
		for _, v := range path[1:] {
			pathString = fmt.Sprintf("%s->%d", pathString, v.Val)
		}

		*result = append(*result, pathString)
	}
	path = append(path, node.Left)
	findPaths2(node.Left, path, result)
	path = path[0 : len(path)-1]

	path = append(path, node.Right)
	findPaths2(node.Right, path, result)
	//path = path[0 : len(path)-1]
}
