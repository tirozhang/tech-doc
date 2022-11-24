package _076_Minimum_Window_Substring

func minWindow(s string, t string) string {
	if s == "" || t == "" {
		return ""
	}
	minString, length := "", len(s)
	targetCounter := makeFindMap(t)
	var counter [128]int

	for l, r := 0, 0; l < length; l++ {
		for r < length && !check(counter, targetCounter) {
			counter[s[r]]++
			r++
		}
		if check(counter, targetCounter) {
			minString = min(minString, s[l:r])
		}
		counter[s[l]]--
	}
	return minString
}

func minWindow2(s string, t string) string {
	if s == "" || t == "" {
		return ""
	}
	minString, length, equalCount := "", len(s), len(t)
	targetCounter := makeFindMap(t)
	matchedChars := 0
	var counter [128]int

	for l, r := 0, 0; l < length; l++ {
		for r < length && matchedChars != equalCount {
			counter[s[r]]++
			if targetCounter[s[r]] > 0 && counter[s[r]] <= targetCounter[s[r]] {
				matchedChars += 1
			}
			r++
		}
		if matchedChars == equalCount {
			minString = min(minString, s[l:r])
		}
		counter[s[l]]--
		if counter[s[l]] < targetCounter[s[l]] {
			matchedChars--
		}
	}
	return minString
}

func makeFindMap(t string) [128]int {
	var str2counter [128]int
	for _, v := range t {
		str2counter[v]++
	}
	return str2counter
}

func check(source, target [128]int) bool {
	for k, targetV := range target {
		if source[k] < targetV {
			return false
		}
	}
	return true
}

func min(a, b string) string {
	if len(a) == 0 || len(b) < len(a) {
		return b
	}
	return a
}
