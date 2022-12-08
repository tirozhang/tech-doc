package help_func

// FindNextWords 找到词典中和s仅一个字母不一样的单词  单词接龙题目
func FindNextWords(s string, wordList []string) []string {
	nextWords := []string{}
	for _, word := range wordList {
		diff := 0
		for k, v := range []rune(s) {
			if v != rune(word[k]) {
				diff++
			}
		}
		if diff == 1 {
			nextWords = append(nextWords, word)
		}
	}
	return nextWords
}
