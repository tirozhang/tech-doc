package _053_Maximum_Subarray

// 暴力解法
func maxSubArray(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	maxSum := nums[0]
	for l, _ := range nums {
		r, sum := l, 0
		for r < length {
			sum += nums[r]
			if sum > maxSum {
				maxSum = sum
			}
			r++
		}
	}
	return maxSum
}

// 滑动窗口 O(n)
func maxSubArray2(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	l, r, maxSum := 0, 0, nums[0]
	for l < length {
		sum := 0
		for r < length && sum >= 0 {
			sum += nums[r]
			if sum > maxSum {
				maxSum = sum
			}
			r++
		}
		l = r
	}
	return maxSum
}

// 前缀和
func maxSubArray3(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	preSum, minPreSum, maxSum := 0, 0, nums[0]
	for _, v := range nums {
		preSum += v
		if maxSum < preSum-minPreSum {
			maxSum = preSum - minPreSum
		}
		if preSum < minPreSum {
			minPreSum = preSum
		}
	}

	return maxSum
}
