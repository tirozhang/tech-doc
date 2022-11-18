package _125_Valid_Palindrome

import "strings"

// isPalindrome 时间复杂度O(n),空间复杂度O(1)
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	isValid, left, right := isValidPalindrome(s, 0, len(s)-1)
	if isValid {
		return true
	}
	isValidLeft, _, _ := isValidPalindrome(s, left+1, right)
	isValidRight, _, _ := isValidPalindrome(s, left, right-1)

	return isValidLeft || isValidRight
}

func isValidPalindrome(s string, left, right int) (bool, int, int) {
	i, j := left, right
	for i < j {
		if s[i] != s[j] {
			return false, i, j
		}
		i++
		j--
	}
	return true, i, j
}
