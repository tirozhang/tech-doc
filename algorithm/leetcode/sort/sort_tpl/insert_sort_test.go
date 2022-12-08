package sort_tpl

import "testing"

func Test_InsertSort(t *testing.T) {
	nums := []int{1, 3, 4, 2, 5}
	insertSort(nums)
	t.Log(nums)
}
