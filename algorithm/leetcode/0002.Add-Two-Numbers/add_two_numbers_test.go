package _002_Add_Two_Numbers

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
	l1 []int
	l2 []int
}

type ans struct {
	one []int
}

func Test_Problem001(t *testing.T) {
	tcs := []question{
		{
			para{[]int{3, 2, 4}, []int{3, 2, 4}},
			ans{[]int{6, 4, 8}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 001------------------------\n")

	for _, tc := range tcs {
		_, p := tc.ans, tc.para
		fmt.Printf("【input】:%v       【output】:%v\n\n", p, addTwoNumbers(structures.Ints2List(p.l1), structures.Ints2List(p.l2)))
	}
	fmt.Println()
}
