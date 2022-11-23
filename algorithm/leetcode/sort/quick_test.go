package sort

import "testing"

func Test_QuickSort(t *testing.T) {
	nums := []int{1, 3, 4, 2, 5}
	quickSort(nums)
	t.Log(nums)
}
