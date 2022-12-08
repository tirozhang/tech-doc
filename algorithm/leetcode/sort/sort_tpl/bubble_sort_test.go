package sort_tpl

import "testing"

func Test_BubbleSort(t *testing.T) {
	nums := []int{1, 3, 4, 2, 5}
	bubbleSort(nums)
	t.Log(nums)
}
