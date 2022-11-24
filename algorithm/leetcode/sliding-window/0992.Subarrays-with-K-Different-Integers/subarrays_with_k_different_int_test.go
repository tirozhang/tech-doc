package _992_Subarrays_with_K_Different_Integers

import "testing"

func Test_Problem992(t *testing.T) {
	t.Log(subarraysWithKDistinct2([]int{1, 2, 1, 2, 3}, 2)) //7
	t.Log(subarraysWithKDistinct2([]int{1, 2}, 1))          //2
	t.Log(subarraysWithKDistinct2([]int{2, 1, 1, 1, 2}, 1)) //8
	t.Log(subarraysWithKDistinct2([]int{2, 1, 2, 1, 2}, 2)) //16

	t.Log(subarraysWithKDistinct([]int{1, 2, 1, 2, 3}, 2)) //7
	t.Log(subarraysWithKDistinct([]int{1, 2}, 1))          //2
	t.Log(subarraysWithKDistinct([]int{2, 1, 1, 1, 2}, 1)) //8
	t.Log(subarraysWithKDistinct([]int{2, 1, 2, 1, 2}, 2)) //16

}
