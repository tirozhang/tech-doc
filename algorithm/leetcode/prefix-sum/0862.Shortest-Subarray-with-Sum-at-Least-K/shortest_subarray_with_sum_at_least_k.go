package _862_Shortest_Subarray_with_Sum_at_Least_K

// 子数组的和>=k 最小子数组
func shortestSubarray(A []int, K int) int {
	res, prefixSum := len(A)+1, getPrefixSum(A)
	// deque 中保存递增的 prefixSum 下标
	deque := []int{}
	for i := range prefixSum {
		// 下面这个循环希望能找到 [deque[0], i] 区间内累加和 >= K，如果找到了就更新答案
		for len(deque) > 0 && prefixSum[i]-prefixSum[deque[0]] >= K {
			length := i - deque[0]
			if res > length {
				res = length
			}
			// 找到第一个 deque[0] 能满足条件以后，就移除它，因为它是最短长度的子序列了
			deque = deque[1:]
		}
		// 下面这个循环希望能保证 prefixSum[deque[i]] 递增
		for len(deque) > 0 && prefixSum[i] <= prefixSum[deque[len(deque)-1]] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
	}
	if res <= len(A) {
		return res
	}
	return -1
}

func getPrefixSum(nums []int) []int {
	prefixSum := make([]int, len(nums)+1)
	for k, v := range nums {
		prefixSum[k+1] = prefixSum[k] + v
	}
	return prefixSum
}
