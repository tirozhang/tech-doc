package _352_Max_Size_Subarray_Sum

func maxSubArrayLen(nums []int, k int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	preSum := getPrefixSum(nums)
	return twoNumEqual(preSum, k)
}

func twoNumEqual(nums []int, target int) int {
	maxLen := 0
	m := make(map[int]int, len(nums))
	for end, v := range nums {
		if left, ok := m[v]; ok {
			if maxLen < end-left {
				maxLen = end - left
			}
		}
		if _, ok := m[v+target]; !ok {
			m[v+target] = end
		}
	}
	return maxLen
}
func getPrefixSum(nums []int) []int {
	preSum := []int{0}
	for k, v := range nums {
		preSum = append(preSum, preSum[k]+v)
	}
	return preSum
}
