package _076_Minimum_Window_Substring

func minWindow(s string, t string) string {
	if s == "" || t == "" {
		return ""
	}

	l, r, minString := 0, 0, ""
	length := len(s)
	freq, showTimes := makeFindMap(t)
	freqLength := len(freq)
	findStrLen := 0
	for l < length {
		if r < length && findStrLen < freqLength {
			if _, ok := freq[int32(s[r])]; ok {
				freq[int32(s[r])]++
				if freq[int32(s[r])] == showTimes[int32(s[r])] {
					findStrLen++
					if findStrLen == freqLength {
						if minString == "" || len(minString) > r-l+1 {
							minString = s[l : r+1]
						}
					}
				}
			}
			r++
		} else {
			if findStrLen < freqLength || l > r {
				break
			}

			if v, ok := freq[int32(s[l])]; ok {
				freq[int32(s[l])]--
				if v == showTimes[int32(s[l])] {
					if len(minString) > r-l {
						minString = s[l:r]
					}
					findStrLen--
				}
			}
			l++
		}
	}
	return minString
}

func makeFindMap(t string) (map[int32]int, map[int32]int) {
	freq := make(map[int32]int, len(t))
	showTimes := make(map[int32]int, len(t))
	for _, v := range t {
		freq[v] = 0
		showTimes[v]++
	}
	return freq, showTimes
}
