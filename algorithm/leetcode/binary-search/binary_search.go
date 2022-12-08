package binary_search

// 递归实现
func binarySearch(nums []int, start, end, target int) int {

	if start > end {
		return -1
	}

	mid := (start + end) / 2
	if nums[mid] == target {
		return mid
	}

	if nums[mid] > target {
		return binarySearch(nums, start, mid-1, target)
	}

	return binarySearch(nums, mid+1, end, target)
}

// 非递归实现
func binarySearch2(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	start, end := 0, len(nums)-1
	for start < end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	if nums[start] == target {
		return start
	}
	return -1
}

// 非递归实现 查找first target
func binarySearch3(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	start, end := 0, len(nums)-1
	for start+1 < end {
		mid := (start + end) / 2 //1、可能偏左或右导致死循环 2、start + (end-start)/2 防止溢出
		if nums[mid] == target {
			end = mid
		}
		if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	if nums[end] == target {
		return end
	}
	if nums[start] == target {
		return start
	}
	return -1
}
