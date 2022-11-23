package _053_Maximum_Subarray

import "testing"

func Test_Problem53(t *testing.T) {
	t.Log(maxSubArray([]int{-1}))                            //0
	t.Log(maxSubArray([]int{}))                              //0
	t.Log(maxSubArray([]int{1, 1}))                          //2
	t.Log(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) //6

}
