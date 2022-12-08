package _179_maxConcat

import "testing"

func Test_maxConcat(t *testing.T) {
	t.Log(maxConcat([]int{123, 123}))  // 123123
	t.Log(maxConcat([]int{1212, 12}))  // 121212
	t.Log(maxConcat([]int{3333, 33}))  // 333333
	t.Log(maxConcat([]int{123, 14}))   // 14123
	t.Log(maxConcat([]int{12345, 12})) // 1234512
	t.Log(maxConcat([]int{12121, 12})) // 1212121
	t.Log(maxConcat([]int{12121, 12})) // 1212121
}
