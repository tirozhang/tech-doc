package _127_Word_Ladder

import "testing"

func Test_WordLadder(t *testing.T) {
	t.Log(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}))
}
