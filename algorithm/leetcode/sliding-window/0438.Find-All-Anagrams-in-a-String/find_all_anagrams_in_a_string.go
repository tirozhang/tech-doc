package _438_Find_All_Anagrams_in_a_String

import (
	"strings"
)

// 超出时间限制
func findAnagrams(s string, p string) []int {
	if len(s) == 0 || len(s) < len(p) {
		return nil
	}
	l, r := 0, len(p)-1
	length := len(s)
	var res []int

	for l < length {
		if r < length && isAnagrams(s[l:r+1], p) {
			res = append(res, l)
		}
		r++
		l++
	}
	return res
}

func isAnagrams(s, p string) bool {
	for _, v := range p {
		if strings.Count(p, string(v)) != strings.Count(s, string(v)) {
			return false
		}
	}
	return true
}

func findAnagrams2(s string, p string) []int {
	var freq [256]int
	result := []int{}
	if len(s) == 0 || len(s) < len(p) {
		return nil
	}
	for i := 0; i < len(p); i++ {
		freq[p[i]-'a']++
	}

	left, right, count := 0, 0, len(p)
	for right < len(s) {
		if freq[s[right]-'a'] >= 1 {
			count--
		}
		freq[s[right]-'a']--
		right++
		if count == 0 {
			result = append(result, left)
		}
		if right-left == len(p) {
			if freq[s[left]-'a'] >= 0 {
				count++
			}
			freq[s[left]-'a']++
			left++
		}

	}

	return result
}
