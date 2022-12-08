package _101_Symmetric_Tree

import (
	"fmt"
	"github.com/halfrost/LeetCode-Go/structures"
	"testing"
)

type question struct {
	para
	ans
}

type para struct {
	one []int
}

type ans struct {
	one bool
}

func Test_Problem101(t *testing.T) {

	tcs := []question{
		{
			para{[]int{3, 4, 4, 5, structures.NULL, structures.NULL, 5, 6, structures.NULL, structures.NULL, 6}},
			ans{true},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 101------------------------\n")

	for _, tc := range tcs {
		_, p := tc.ans, tc.para
		fmt.Printf("【input】:%v      ", p)
		rootOne := structures.Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", isSymmetric(rootOne))
	}
}
