package sort

import "testing"

func Test_SelectSort(t *testing.T) {
	nums := []int{1, 3, 4, 2, 5}
	selectSort(nums)
	t.Log(nums)
}
