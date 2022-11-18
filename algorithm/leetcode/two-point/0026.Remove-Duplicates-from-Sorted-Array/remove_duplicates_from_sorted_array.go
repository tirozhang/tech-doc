package _026_Remove_Duplicates_from_Sorted_Array

func removeDuplicates(nums []int) int {
	n := 0
	for k, v := range nums {
		if k == 0 || v != nums[k-1] {
			nums[n] = v
			n++
		}
	}
	return n
}
