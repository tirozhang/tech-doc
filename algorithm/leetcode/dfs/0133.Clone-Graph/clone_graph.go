package _133_Clone_Graph

import (
	"github.com/halfrost/LeetCode-Go/structures"
)

// NestedInteger define
type NestedInteger = structures.NestedInteger

// dfs
var visited = map[*NestedInteger]*NestedInteger{}

func cloneGraph(node *NestedInteger) *NestedInteger {
	if node == nil {
		return nil
	}
	if visited[node] != nil {
		return visited[node]
	}
	head := &NestedInteger{Num: node.Num, Ns: []*NestedInteger{}}
	visited[node] = head
	for _, n := range node.Ns {
		head.Ns = append(head.Ns, cloneGraph(n))
	}
	return head
}
