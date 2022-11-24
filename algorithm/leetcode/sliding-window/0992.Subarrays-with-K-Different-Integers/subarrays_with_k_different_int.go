package _992_Subarrays_with_K_Different_Integers

func subarraysWithKDistinct(nums []int, k int) int {
	length := len(nums)
	if length < k {
		return 0
	}
	subArrayNum := 0
	return subArrayNum
}

func subarraysWithKDistinct3(nums []int, k int) int {
	length := len(nums)
	if length < k {
		return 0
	}
	subArrayNum := 0
	counter := make(map[int]int, k+1)
	counter[nums[0]]++

	for l, r := 0, 1; l < length-k+1; l++ {
		for r < length && len(counter) <= k {
			if len(counter) == k {
				subArrayNum += sumSubArray(nums[l:r], k)
			}
			counter[nums[r]]++
			r++
		}

		if len(counter) == k {
			subArrayNum += sumSubArray(nums[l:r], k)
			return subArrayNum
		}

		counter[nums[l]]--
		if counter[nums[l]] == 0 {
			delete(counter, nums[l])
		}

	}
	return subArrayNum
}

func sumSubArray(nums []int, k int) int {
	//fmt.Println(nums)
	counter := make(map[int]int, k+1)
	for r := len(nums) - 1; r >= 0; r-- {
		counter[nums[r]]++
		if len(counter) == k {
			return r + 1
		}
	}
	return 0
}

// 暴力算法
func subarraysWithKDistinct2(nums []int, k int) int {

	length := len(nums)
	if length < k {
		return 0
	}
	subArrayNum := 0
	for l := 0; l < length-k+1; l++ {
		counter := make(map[int]int, k+1)
		for r := l; r < length; r++ {
			counter[nums[r]]++
			if len(counter) == k {
				subArrayNum++
			}
			if len(counter) > k {
				break
			}
		}
	}
	return subArrayNum
}
