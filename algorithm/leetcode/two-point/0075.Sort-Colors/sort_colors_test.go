package _075_Sort_Colors

import "testing"

func Test_Problem75(t *testing.T) {
	q := []int{2, 0, 2, 1, 1, 0}
	sortColors2(q)
	t.Log(q)
	q = []int{0, 0}
	sortColors2(q)
	t.Log(q)
	q = []int{0, 1}
	sortColors2(q)
	t.Log(q)
}
