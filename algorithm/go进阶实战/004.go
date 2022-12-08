package main

import (
	"fmt"
)

// 判断两个给定的字符串排序后是否一致
func main() {
	fmt.Println(isRegroup("tnt", "tnt"))
	fmt.Println(isRegroup("test", "test"))

}

func isRegroup(s1, s2 string) bool {
	r1, r2 := []rune(s1), []rune(s2)
	if len(r1) != len(r2) {
		return false
	}
	l := len(r1)
	for i := 0; i < l; i++ {
		if r1[i] != r1[l-i-1] {
			return false
		}
	}

	// 方法二
	//for _, v := range r1 {
	//	if strings.Count(s1, string(v)) != strings.Count(s2, string(v)) {
	//		return false
	//	}
	//}
	return true
}
