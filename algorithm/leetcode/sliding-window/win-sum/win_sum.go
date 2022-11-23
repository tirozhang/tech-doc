package win_sum

// WinSum 滑动窗口实现 单指针实现
func WinSum(nums []int, k int) []int {
	// write your code here
	if len(nums) < k || k == 0 {
		return []int{}
	}
	l, sum, length := 1, 0, len(nums)
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	ans := []int{sum}
	for l <= length-k {
		ans = append(ans, ans[l-1]-nums[l-1]+nums[l+k-1])
		l++
	}
	return ans
}

func WinSum2(nums []int, k int) []int {
	r, length := 0, len(nums)
	if length < k || k == 0 {
		return []int{}
	}
	var ans []int
	sum := 0
	for l, v := range nums {
		for r < length && r < l+k {
			sum += nums[r]
			r++
		}
		if r-l == k {
			ans = append(ans, sum)
			sum = sum - v
		}
	}
	return ans
}
