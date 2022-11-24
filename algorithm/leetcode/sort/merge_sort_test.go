package sort

import "testing"

func Test_MergeSort(t *testing.T) {
	nums := []int{1, 3, 4, 2, 5}
	mergeSort(nums, 0, 4)
	t.Log(nums)
	nums = []int{1, 3, 4, 2, 5}
	mergeSort2(nums, 0, 4)
	t.Log(nums)
}
