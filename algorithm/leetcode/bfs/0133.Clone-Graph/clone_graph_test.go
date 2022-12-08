package _133_Clone_Graph

import (
	"fmt"
	"testing"
)

func Test_CloneGraph(t *testing.T) {
	node1 := &NestedInteger{
		Num: 1,
	}
	node2 := &NestedInteger{
		Num: 2,
	}
	node3 := &NestedInteger{
		Num: 3,
	}

	node1.Ns = []*NestedInteger{node2, node3}
	node2.Ns = []*NestedInteger{node1, node3}
	node3.Ns = []*NestedInteger{node1, node2}

	newNode := cloneGraph(node1)
	for _, node := range newNode.Ns {
		fmt.Println(node.Num)
		//	//for _, nodens := range node.Ns {
		//	//	fmt.Println(nodens.Num)
		//	//}
	}
}
