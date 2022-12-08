package _127_Word_Ladder

import "leetcode/help_func"

func ladderLength(beginWord string, endWord string, wordList []string) int {

	visited := map[string]int{beginWord: 1}
	queue := []string{beginWord}
	distance := 0
	wordList = append(wordList, endWord)

	for len(queue) > 0 {
		length := len(queue)
		distance++
		for i := 0; i < length; i++ {
			word := queue[0]
			queue = queue[1:]
			nextWords := help_func.FindNextWords(word, wordList)
			for _, v := range nextWords {
				if v == endWord {
					return distance + 1
				}
				if _, ok := visited[v]; ok {
					continue
				}
				visited[v] = 1
				queue = append(queue, v)
			}
		}
	}

	return 0
}
