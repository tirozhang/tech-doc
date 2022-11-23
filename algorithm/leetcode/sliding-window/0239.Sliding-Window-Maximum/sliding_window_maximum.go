package _239_Sliding_Window_Maximum

// 方法一：暴力法
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k > len(nums) {
		return []int{}
	}
	l, r, length := 0, k-1, len(nums)
	var res []int

	for r < length {
		res = append(res, findMaxinum(nums[l:r+1]))
		r++
		l++
	}
	return res
}

func findMaxinum(subNums []int) int {
	max := subNums[0]
	for _, v := range subNums {
		if v > max {
			max = v
		}
	}
	return max
}

// 方法二 压缩重复计算次数
func maxSlidingWindow2(nums []int, k int) []int {
	if len(nums) == 0 || k > len(nums) {
		return []int{}
	}
	l, r, length := 0, k-1, len(nums)
	res := make([]int, length-k+1)

	for r < length {
		if l > 0 && nums[l-1] < res[l-1] && nums[r] < res[l-1] {
			res[l] = res[l-1]
		} else {
			res[l] = findMaxinum(nums[l : r+1])
		}
		r++
		l++
	}
	return res
}

// 方法三 单调队列
func maxSlidingWindow3(nums []int, k int) []int {

	if len(nums) == 0 || k > len(nums) {
		return []int{}
	}
	l, length := 0, len(nums)
	res := make([]int, length-k+1)
	var queue []int

	for l < length {
		if len(queue) > 0 && l >= k && queue[0] <= l-k {
			queue = queue[1:]
		}

		for len(queue) > 0 && nums[l] > nums[queue[len(queue)-1]] {
			queue = queue[0 : len(queue)-1]
		}
		queue = append(queue, l)

		if l >= k-1 {
			res[l-k+1] = nums[queue[0]]
		}
		l++
	}
	return res

}
