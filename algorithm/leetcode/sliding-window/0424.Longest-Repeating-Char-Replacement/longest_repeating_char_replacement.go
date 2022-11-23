package _424_Longest_Repeating_Char_Replacement

func characterReplacement(s string, k int) int {
	if len(s) == 0 {
		return 0
	}
	n, ans := len(s), 0
	for l, r := 0, 0; r < n; r++ {
		for r < n && check(s[l:r+1]) <= k {
			if r-l+1 > ans {
				ans = r - l + 1
			}
			r++
		}
		for r < n && l < r && check(s[l:r+1]) > k {
			l++
		}

	}
	return ans
}

func check(s string) int {
	var cnt [256]int
	maxRepeat := 1
	for _, ch := range s {
		cnt[ch]++
		if maxRepeat < cnt[ch] {
			maxRepeat = cnt[ch]
		}
	}
	return len(s) - maxRepeat
}

func characterReplacement2(s string, k int) int {
	if len(s) == 0 {
		return 0
	}
	n, ans, maxRepeat := len(s), 0, 1
	var cnt [128]int

	for l, r := 0, 0; l < n; l++ {
		for r < n && r-l-maxRepeat <= k { // r:l-1的最大长度
			cnt[s[r]]++
			if cnt[s[r]] > maxRepeat {
				maxRepeat = cnt[s[r]]
			}
			ans = maxNum(ans, r-l)
			r++
		}
		if r-l-maxRepeat <= k {
			ans = maxNum(ans, r-l)
		}

		cnt[s[l]]--
		maxRepeat = max(cnt)
	}

	return ans
}

func max(arr [128]int) int {
	res := 0
	for _, v := range arr {
		if res < v {
			res = v
		}
	}
	return res
}

func maxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}
