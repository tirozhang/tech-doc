package _209_Minimum_Size_Subarray_Sum

import "testing"

func Test_Problem209(t *testing.T) {

	t.Log(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))    //2
	t.Log(minSubArrayLen(15, []int{1, 2, 3, 4, 5}))      //5
	t.Log(minSubArrayLen(5, []int{2, 3, 1, 1, 1, 1, 1})) //2
	t.Log(minSubArrayLen(5, []int{1, 5, 6, 1, 7, 8}))    //2
	t.Log("-------")
	t.Log(minSubArrayLen2(7, []int{2, 3, 1, 2, 4, 3}))    //2
	t.Log(minSubArrayLen2(15, []int{1, 2, 3, 4, 5}))      //5
	t.Log(minSubArrayLen2(5, []int{2, 3, 1, 1, 1, 1, 1})) //2
	t.Log(minSubArrayLen2(4, []int{1, 5, 6, 1, 7, 8}))    //0
}
