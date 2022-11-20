package _915_Partition_Array

func PartitionArray(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	length := len(nums)
	l, r := 0, length-1
	for l < r {
		for l <= r && nums[l] < k {
			l++
		}
		for l <= r && nums[r] >= k {
			r--
		}
		if l <= r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}
	return l
}
