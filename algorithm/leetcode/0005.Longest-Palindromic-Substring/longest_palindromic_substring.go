package _005_Longest_Palindromic_Substring

// 方法一 暴力破解 O(n^3)
func longestPalindrome(s string) string {
	mString, le := "", len(s)
	for i := 0; i < le; i++ {
		for j := i + 1; j < le; j++ {
			l, r := i, j
			for l <= r && s[l] == s[r] {
				l++
				r--
			}
			if l >= r && len(mString) < j-i+1 {
				mString = s[i : j+1]
			}
		}
	}
	return mString
}

// 方法二 暴力破解 O(n^3)
func longestPalindrome2(s string) string {
	if s == "" || len(s) == 1 {
		return s
	}
	mString, length := s[0:1], len(s)
	for left := 0; left < length; left++ {
		for right := left + 1; right < length; right++ {
			if isPalindrome(s, left, right) && len(mString) < right-left+1 {
				mString = s[left : right+1]
			}
		}
	}
	return mString
}

func isPalindrome(s string, l, r int) bool {
	for l <= r && s[l] == s[r] {
		l++
		r--
	}
	return l >= r
}

// 方法三 基于中心线枚举的算法
func longestPalindrome3(s string) string {
	if s == "" || len(s) == 1 {
		return s
	}
	mString, length := s[0:1], len(s)
	for mid := 0; mid < length; mid++ {
		mString = maxString(getPalindromeFrom(s, mid, mid), mString)
		mString = maxString(getPalindromeFrom(s, mid, mid+1), mString)
	}
	return mString
}

func getPalindromeFrom(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}

func maxString(s1, s2 string) string {
	if len(s1) > len(s2) {
		return s1
	}
	return s2
}
