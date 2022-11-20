package _125_Valid_Palindrome

import "strings"

// isPalindrome 时间复杂度O(n),空间复杂度O(1)
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	r := []rune(s)
	i, j := 0, len(r)-1
	for i <= j {
		for i < j && !isChar(s[i]) {
			i++
		}

		for i < j && !isChar(s[j]) {
			j--
		}

		if s[i] != s[j] {
			return false
		}
		i++
		j--

	}
	return true
}

func isChar(c byte) bool {
	if ('a' <= c && c <= 'z') || ('0' <= c && c <= '9') {
		return true
	}
	return false
}
