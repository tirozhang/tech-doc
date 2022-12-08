package main

import (
	"fmt"
	"strings"
)

// 判断字符串中字符是否唯一
func main() {
	s := "abcdefghi"
	fmt.Println(isUnique(s))
	s = "aaaa"
	fmt.Println(isUniqueString(s))

}

func isUnique(astr string) bool {
	var m = make(map[byte]bool)
	for i := 0; i < len(astr); i++ {
		if m[astr[i]] {
			return false
		}
		m[astr[i]] = true
	}
	return true
}

func isUniqueString(s string) bool {
	if strings.Count(s, "") > 3000 {
		return false
	}

	for k, v := range s {

		// 方法一 strings.Count
		//if strings.Count(s, string(v)) > 1 {
		//	return false
		//}

		// 方法二 strings.Index
		if strings.Index(s, string(v)) != k {
			return false
		}
	}
	return true
}
