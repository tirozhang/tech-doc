package _915_Partition_Array

import "testing"

func Test_Problem915(t *testing.T) {
	t.Log(PartitionArray([]int{3, 2, 2, 1}, 2))
	t.Log(PartitionArray([]int{}, 2))
	t.Log(PartitionArray([]int{7, 7, 9, 8, 6, 6, 8, 7, 9, 8, 6, 6}, 10))

}
