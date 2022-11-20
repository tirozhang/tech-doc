package _167_Two_Sum_SortArray

import (
	"fmt"
	"testing"
)

type question struct {
	para
	ans
}

type para struct {
	nums   []int
	target int
}

type ans struct {
	one []int
}

func Test_Problem167(t *testing.T) {

	tcs := []question{
		{
			para{[]int{2, 7, 11, 15}, 9},
			ans{[]int{1, 2}},
		},
		{
			para{[]int{1, 2, 7, 11, 15}, 26},
			ans{[]int{4, 5}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 001------------------------\n")

	for _, tc := range tcs {
		_, p := tc.ans, tc.para
		fmt.Printf("【input】:%v       【output】:%v\n\n", p, twoSumSortArray(p.nums, p.target))
	}
	fmt.Println()
}
