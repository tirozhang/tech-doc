package _283_Move_Zeroes

import "testing"

func Test_Problem283(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	t.Log(nums)
}
