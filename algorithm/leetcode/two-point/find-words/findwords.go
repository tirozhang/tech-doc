package find_words

/**
 * @param str: the string
 * @param dict: the dictionary
 * @return: return words which  are subsequences of the string
 */
func FindWords(str string, dict []string) []string {
	// write your code here.
	str2index := [128]int{}
	res := []string{}
	for index, v := range str {
		str2index[v] = index
	}
	for _, v := range dict {
		if check(str2index, v) {
			res = append(res, v)
		}
	}
	return res
}

func check(str2index [128]int, s string) bool {
	start := 0
	for _, ch := range s {
		if str2index[ch] == 0 && str2index[ch] <= start {
			return false
		}
		start = str2index[ch]
		// str2index[ch]--
	}
	return true

}
