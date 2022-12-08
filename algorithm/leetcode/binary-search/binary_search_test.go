package binary_search

import "testing"

func Test_BinarySearch(t *testing.T) {
	t.Log(binarySearch([]int{1, 1, 1}, 0, 2, 1))
	t.Log(binarySearch([]int{1, 2, 2, 4, 5}, 0, 4, 2))
	t.Log(binarySearch([]int{1, 1, 1, 1, 2}, 0, 4, 2))

	t.Log(binarySearch2([]int{1, 1, 1}, 1))
	t.Log(binarySearch2([]int{1, 2, 2, 4, 5}, 2))
	t.Log(binarySearch2([]int{1, 1, 1, 1, 2}, 2))

	t.Log(binarySearch3([]int{1, 1, 1}, 1))
	t.Log(binarySearch3([]int{1, 2, 2, 4, 5}, 2))
	t.Log(binarySearch3([]int{1, 1, 1, 2, 2}, 2))

}
