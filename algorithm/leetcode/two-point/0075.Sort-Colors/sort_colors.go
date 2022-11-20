package _075_Sort_Colors

import _915_Partition_Array "leetcode/two-point/0915.Partition-Array"

func sortColors(nums []int) {
	if len(nums) < 2 {
		return
	}
	l, r, i := 0, len(nums)-1, 0
	for l <= r {

		for i <= r && nums[i] == 1 {
			i++
		}
		if i > r {
			return
		}

		if nums[i] > 1 {
			nums[r], nums[i] = nums[i], nums[r]
			r--
		} else if nums[i] < 1 {
			nums[l], nums[i] = nums[i], nums[l]
			l++
			i++
		}
	}
}

func sortColors2(nums []int) {
	if len(nums) < 2 {
		return
	}
	l := _915_Partition_Array.PartitionArray(nums, 1)
	_915_Partition_Array.PartitionArray(nums[l:], 2)
}

func sortColors3(nums []int) {
	if len(nums) < 2 {
		return
	}
	l, r, i := 0, len(nums)-1, 0
	for i <= r {
		switch {
		case nums[i] > 1:
			nums[r], nums[i] = nums[i], nums[r]
			r--
		case nums[i] < 1:
			nums[l], nums[i] = nums[i], nums[l]
			l++
			i++
		default:
			i++
		}
	}
}
