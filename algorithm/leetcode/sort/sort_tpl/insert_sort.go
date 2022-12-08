package sort_tpl

import "leetcode/help_func"

func insertSort(nums []int) {
	start, length := 1, len(nums)
	for start < length {
		index := 0
		for index <= start {
			if nums[start] >= nums[index] {
				index++
			} else {
				help_func.Reverse(nums[index:start])
				help_func.Reverse(nums[index : start+1])
			}
		}
		start++
	}
}
