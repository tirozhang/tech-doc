package _026_Remove_Duplicates_from_Sorted_Array

import "sort"

//  删除有序数组中的重复项
func removeDuplicates(nums []int) int {
	n := 0
	for k, v := range nums {
		if k == 0 || v != nums[k-1] {
			nums[n] = v
			n++
		}
	}
	return n
}

//  删除无序数组中的重复项 时间复杂度O(n) 空间复杂度O(n)
func removeDuplicates2(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return length
	}
	l, r, freq := 0, length-1, make(map[int]int)
	for l <= r {
		if _, ok := freq[nums[l]]; !ok {
			freq[nums[l]] = 1
			l++
			continue
		}
		for _, ok := freq[r]; r > l && ok; r-- {
		}

		nums[l], nums[r] = nums[r], nums[l]
		freq[nums[l]] = 1
		l++
		r--
	}
	return len(freq)
}

//  删除无序数组中的重复项 时间复杂度O(nlogn) 空间复杂度O(1)
func removeDuplicates3(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return length
	}
	sort.Ints(nums)
	return removeDuplicates(nums)
}

// 同向双指针实现
func removeDuplicates4(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return length
	}
	l, r := 0, 1
	sort.Ints(nums)
	for l < r {
		for r < length && nums[r] == nums[l] {
			r++
		}
		if r >= length {
			break
		}
		nums[l+1] = nums[r]
		l++
		r++
	}
	return l + 1
}
