package win_sum

import "testing"

func Test_Winsum(t *testing.T) {
	t.Log(WinSum([]int{1, 2, 7, 7, 2}, 3))

	t.Log(WinSum2([]int{1, 2, 7, 7, 2}, 3))
}
