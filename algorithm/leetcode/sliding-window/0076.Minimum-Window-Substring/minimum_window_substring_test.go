package _076_Minimum_Window_Substring

import "testing"

func Test_Problem76(t *testing.T) {
	t.Log(minWindow("ADOBECODEBANC", "ABC"))
	t.Log(minWindow("aa", "aa"))
	t.Log("-------")
	t.Log(minWindow2("ADOBECODEBANC", "ABC"))
	t.Log(minWindow2("aa", "aa"))
	t.Log(minWindow2("acbbaca", "aba"))

}
