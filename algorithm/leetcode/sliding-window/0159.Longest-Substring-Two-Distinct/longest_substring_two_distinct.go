package _159_Longest_Substring_Two_Distinct

func lengthOfLongestSubstringTwoDistinct(s string) int {
	if len(s) == 0 {
		return 0
	}
	n, ans := len(s), 0
	var cnt [128]int
	for l, r := 0, 0; l < n; l++ {
		for r < n && check(cnt) <= 2 {
			cnt[s[r]]++
			ans = max(ans, r-l)
			r++
		}
		if check(cnt) <= 2 {
			ans = max(ans, r-l)
		}
		cnt[s[l]]--
	}
	return ans
}
func check(cnt [128]int) int {
	counter := 0
	for _, v := range cnt {
		if v > 0 {
			counter += 1
		}
	}
	return counter
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
