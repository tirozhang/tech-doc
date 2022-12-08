package _133_Clone_Graph

import (
	"github.com/halfrost/LeetCode-Go/structures"
)

// NestedInteger define
type NestedInteger = structures.NestedInteger

func cloneGraph(node *NestedInteger) *NestedInteger {
	if node == nil {
		return nil
	}

	// 第一步:找到所有的点并复制 入queue即标记
	mapping := map[*NestedInteger]*NestedInteger{node: {Num: node.Num}}
	queue := []*NestedInteger{node}

	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]
		for _, neighbor := range currentNode.Ns {
			if _, ok := mapping[neighbor]; ok {
				continue
			}
			// 入队后立刻标志 防止重复入队
			newNode := &NestedInteger{Num: neighbor.Num}
			mapping[neighbor] = newNode
			queue = append(queue, neighbor)
		}
	}
	// 第二步:找到所有的边并复制
	for currentNode, newNode := range mapping {
		for _, neighbor := range currentNode.Ns {
			newNode.Ns = append(newNode.Ns, mapping[neighbor])
		}
	}
	return mapping[node]
}
