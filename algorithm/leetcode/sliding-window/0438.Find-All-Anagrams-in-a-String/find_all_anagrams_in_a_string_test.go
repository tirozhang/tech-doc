package _438_Find_All_Anagrams_in_a_String

import (
	"testing"
)

func Test_Problem438(t *testing.T) {
	t.Log(isAnagrams("cbaebabacd", "abc"))
	t.Log(findAnagrams("cbaebabacd", "abc"))
	t.Log(findAnagrams("abab", "ab"))

}
