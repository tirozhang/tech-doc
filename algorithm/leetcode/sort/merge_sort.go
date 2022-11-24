package sort

func mergeSort(nums []int, start, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2

	mergeSort(nums, start, mid)
	mergeSort(nums, mid+1, end)

	// merge
	left, right := start, mid+1
	var temp []int
	for left <= mid && right <= end {
		if nums[left] < nums[right] {
			temp = append(temp, nums[left])
			left++
		} else {
			temp = append(temp, nums[right])
			right++
		}
	}

	for left <= mid {
		temp = append(temp, nums[left])
		left++
	}

	for right <= end {
		temp = append(temp, nums[right])
		right++
	}

	for k, v := range temp {
		nums[k+start] = v
	}
}

// 原址排序
func mergeSort2(nums []int, start, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2

	mergeSort2(nums, start, mid)
	mergeSort2(nums, mid+1, end)

	// merge
	left, right, index := start, mid+1, mid+1

	for left <= right && right <= end {
		if nums[left] < nums[right] {
			for index < right {
				swap(nums, left, index, right)
				index = right
			}
			left++
		} else {
			right++
		}
	}
}
