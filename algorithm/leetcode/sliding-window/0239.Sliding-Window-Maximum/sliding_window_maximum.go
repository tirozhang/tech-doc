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

/*
type SortArray []int

func (s *SortArray) push(num int) {}
func (s *SortArray) push(num int) {}

// 方法三 有序数组
func maxSlidingWindow3(nums []int, k int) []int {

	if len(nums) == 0 || k > len(nums) {
		return []int{}
	}
	l, r, length := 0, k-1, len(nums)
	res := make([]int, length-k+1)
	sortArr := make([]int, k)

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

*/
