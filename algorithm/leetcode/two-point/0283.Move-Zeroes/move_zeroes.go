package _283_Move_Zeroes

func moveZeroes(nums []int) {
	n := 0
	for k := range nums {
		if nums[k] != 0 {
			nums[n] = nums[k]
			n++
		}
	}
	for i := n; i < len(nums); i++ {
		nums[i] = 0
	}
	return
}
