package sort

func insertSort(nums []int) {
	start, length := 1, len(nums)
	for start < length {
		index := 0
		for index <= start {
			if nums[start] >= nums[index] {
				index++
			} else {
				reverse(nums[index:start])
				reverse(nums[index : start+1])
			}
		}
		start++
	}
}
