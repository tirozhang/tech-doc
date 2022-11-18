package _125_Valid_Palindrome

import (
	"fmt"
	"testing"
)

func Test_Problem125(t *testing.T) {
	tcs := []struct {
		s   string
		ans bool
	}{
		{
			"",
			true,
		},
		{
			"，，，，，，",
			true,
		},
		{
			"0p",
			false,
		},

		{
			"0",
			true,
		},

		{
			"race a car",
			false,
		},

		{
			"A man, a plan, a canal: Panama",
			true,
		},
	}
	fmt.Printf("------------------------Leetcode Problem 125------------------------\n")

	for _, tc := range tcs {
		fmt.Printf("【input】:%v       【output】:%v\n", tc, isPalindrome(tc.s))
	}
	fmt.Println()
}
