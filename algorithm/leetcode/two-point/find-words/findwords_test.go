package find_words

import "testing"

func Test_Findwords(t *testing.T) {
	t.Log(FindWords("bcogtadsjofisdhklasdj", []string{"book", "code", "tag"}))
}
