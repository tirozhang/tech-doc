package sort

import "testing"

func Test_QuickSort(t *testing.T) {
	nums := []int{1, 3, 4, 2, 5}
	quickSort2(nums, 0, 4)
	t.Log(nums)
	nums = []int{5, 2, 1, 1, 1}
	quickSort(nums)
	t.Log(nums)
	nums = []int{1}
	quickSort(nums)
	t.Log(nums)
	nums = []int{5, 4, 3}
	quickSort(nums)
	t.Log(nums)
	nums = []int{}
	quickSort(nums)
	t.Log(nums)
}
