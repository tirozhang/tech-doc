package _003_Longest_Substring_Without_Repeating_Characters

func lengthOfLongestSubstring(s string) int {

	if len(s) == 0 {
		return 0
	}
	var freq [127]int
	result, left, right := 0, 0, -1

	for left < len(s) {
		if right+1 < len(s) && freq[s[right+1]] == 0 {
			freq[s[right+1]]++
			right++

		} else {
			freq[s[left]]--
			left++
		}
		result = max(result, right-left+1)
	}
	return result
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring2(s string) int {
	n, ans := len(s), 0
	for l, r := 0, 0; r < n; r++ {
		for r < n && check(s[l:r+1]) {
			if r-l+1 > ans {
				ans = r - l + 1
			}
			r++
		}
		for r < n && l <= r && !check(s[l:r+1]) {
			l++
		}
	}
	return ans
}

func check(s string) bool {
	var cnt [256]int
	for _, ch := range s {
		if cnt[ch] > 0 {
			return false
		}
		cnt[ch]++
	}
	return true
}
