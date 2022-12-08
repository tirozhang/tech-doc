package sort_tpl

func selectSort(nums []int) {
	index, length := 0, len(nums)
	for index < length {
		findIndex := index
		for start := index + 1; start < length; start++ {
			if nums[start] < nums[findIndex] {
				findIndex = start
			}
		}
		nums[findIndex], nums[index] = nums[index], nums[findIndex]
		index++
	}
}
