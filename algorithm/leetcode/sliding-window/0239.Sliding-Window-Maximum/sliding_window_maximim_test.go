package _239_Sliding_Window_Maximum

import "testing"

func Test_Problem239(t *testing.T) {
	t.Log(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	t.Log(maxSlidingWindow([]int{1, -1}, 1))
	t.Log(maxSlidingWindow3([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	t.Log(maxSlidingWindow3([]int{1, -1}, 1))

}
