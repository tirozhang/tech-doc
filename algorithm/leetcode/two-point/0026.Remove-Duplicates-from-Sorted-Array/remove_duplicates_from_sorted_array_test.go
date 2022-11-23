package _026_Remove_Duplicates_from_Sorted_Array

import (
	"fmt"
	"testing"
)

func Test_Problem26(t *testing.T) {
	arr := []int{1, 2, 2, 3, 3}
	fmt.Println(removeDuplicates(arr), arr)
	arr = []int{1, 3, 1, 2, 0, 2}
	fmt.Println(removeDuplicates2(arr), arr)
	arr = []int{1, 3, 1, 2, 0, 2}
	fmt.Println(removeDuplicates3(arr), arr)
	arr = []int{1, 3, 1, 2, 0, 2}
	fmt.Println(removeDuplicates4(arr), arr)
}
