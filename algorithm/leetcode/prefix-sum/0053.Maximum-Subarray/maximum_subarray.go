package _053_Maximum_Subarray

// åįžå
func maxSubArray(nums []int) int {
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
