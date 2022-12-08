package dfs_tpl

import (
	"github.com/halfrost/LeetCode-Go/structures"
	"testing"
)

func Test_DFS(t *testing.T) {
	t.Log(findNodes(structures.Ints2TreeNode([]int{1, 2, 3, structures.NULL, 5})))
	t.Log(findPaths(structures.Ints2TreeNode([]int{1, 2, 3, structures.NULL, 5})))
	t.Log(findPaths(structures.Ints2TreeNode([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})))
	t.Log(binaryTreePaths(structures.Ints2TreeNode([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})))

}
