package _001_Two_Sum

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

func Test_Problem001(t *testing.T) {

	tcs := []question{
		{
			para{[]int{3, 2, 4}, 6},
			ans{[]int{1, 2}},
		},

		{
			para{[]int{3, 2, 4}, 5},
			ans{[]int{0, 1}},
		},

		{
			para{[]int{0, 8, 7, 3, 3, 4, 2}, 11},
			ans{[]int{1, 3}},
		},

		{
			para{[]int{0, 1}, 1},
			ans{[]int{0, 1}},
		},

		{
			para{[]int{0, 3}, 5},
			ans{[]int{}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 001------------------------\n")

	for _, tc := range tcs {
		_, p := tc.ans, tc.para
		fmt.Printf("【input】:%v       【output】:%v\n\n", p, twoSum2(p.nums, p.target))
	}
	fmt.Println()
}
