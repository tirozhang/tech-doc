package _348_Zero_Filled_Subarray

func zeroFilledSubarray(nums []int) int {
	r, length, counter := 0, len(nums)-1, 0
	for l, v := range nums {
		if v != 0 || (v == 0 && l < r) {
			continue
		}
		if l > r {
			r = l + 1
		}
		for r <= length && nums[r] == 0 {
			r++
		}
		counter += (r - l) * (r - l + 1) / 2
	}
	return counter
}

func zeroFilledSubarray2(nums []int) int {
	r, answer, length := 1, 0, len(nums)
	if length == 0 {
		return 0
	}
	for l, v := range nums {
		if v != 0 {
			continue
		}
		if r < l+1 {
			r = l + 1
		}
		for r < length && nums[r] == 0 {
			r++
		}
		answer += r - l
	}
	return answer
}
