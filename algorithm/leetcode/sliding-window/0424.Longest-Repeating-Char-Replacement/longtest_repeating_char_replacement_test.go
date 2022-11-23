package _424_Longest_Repeating_Char_Replacement

import "testing"

func Test_Problem424(t *testing.T) {
	t.Log(characterReplacement("AABABBA", 1))
	t.Log(characterReplacement("ABAA", 0))
	t.Log(characterReplacement2("AABABBA", 1))
	t.Log(characterReplacement2("ABAA", 0))
}
