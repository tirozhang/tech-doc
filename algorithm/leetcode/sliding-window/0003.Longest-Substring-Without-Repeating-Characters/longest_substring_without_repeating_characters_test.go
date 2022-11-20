package _003_Longest_Substring_Without_Repeating_Characters

import (
	"fmt"
	"testing"
)

type question struct {
	para string
	ans  int
}

func Test_Problem3(t *testing.T) {
	qs := []question{
		{para: "abcdefgabcdefgh", ans: 7},
	}
	fmt.Printf("------------------------Leetcode Problem 3------------------------\n")

	for _, q := range qs {
		fmt.Printf("【input】:%v       【output】:%v\n", q.para, lengthOfLongestSubstring(q.para))
	}
}
