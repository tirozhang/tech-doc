package _209_Minimum_Size_Subarray_Sum

// 前缀和
func minSubArrayLen(target int, nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	preSum := getPrefixSum(nums)
	return twoNum(preSum, target)
}
func minSubArrayLen2(target int, nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	preSum := getPrefixSum(nums)
	return twoNumEqual(preSum, target)
}

// 有序数组双指针 求 >=target
func twoNum(nums []int, target int) int {
	length := len(nums)
	l, r, minLen := 0, 1, 0
	for l < length {
		for r < length && nums[r]-nums[l] < target {
			r++
		}
		if r == length {
			break
		}

		if minLen == 0 || r-l < minLen {
			minLen = r - l
		}
		l++
	}
	return minLen
}

// 哈希表 求==target
func twoNumEqual(nums []int, target int) int {
	minLen := 0
	m := make(map[int]int, len(nums))
	for end, v := range nums {
		if left, ok := m[v]; ok {
			if minLen == 0 || minLen > end-left {
				minLen = end - left
			}
		}
		m[v+target] = end
	}
	return minLen
}

func getPrefixSum(nums []int) []int {
	preSum := []int{0}
	for k, v := range nums {
		preSum = append(preSum, preSum[k]+v)
	}
	return preSum
}
