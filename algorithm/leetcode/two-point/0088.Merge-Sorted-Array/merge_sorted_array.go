package _088_Merge_Sorted_Array

/*
输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
输出：[1,2,2,3,5,6]
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	size := m + n
	for i := size - 1; m > 0 && n > 0; i-- {
		if nums1[m-1] > nums2[n-1] {
			nums1[i] = nums1[m-1]
			m--
		} else {
			nums1[i] = nums2[n-1]
			n--
		}
	}
	for ; n > 0; n-- {
		nums1[n-1] = nums2[n-1]
	}
}

func merge2(nums1 []int, m int, nums2 []int, n int) {
	size := m + n
	for m > 0 && n > 0 {
		if nums1[m-1] > nums2[n-1] {
			nums1[size-1] = nums1[m-1]
			m--
		} else {
			nums1[size-1] = nums2[n-1]
			n--
		}
		size--
	}
	for n > 0 {
		nums1[n-1] = nums2[n-1]
		n--
	}
}
