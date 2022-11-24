package sort

// 快速排序
// 先先整体有序，再局部有序的分治算法
func quickSort(nums []int) {
	if len(nums) == 0 {
		return
	}
	mid, length := nums[0], len(nums)
	l, r := 1, length-1
	for l < r {
		for l <= r && nums[l] < mid {
			l++
		}
		for l <= r && nums[r] > mid {
			r--
		}
		if l <= r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}
	nums[0], nums[l-1] = nums[l-1], nums[0]
	quickSort(nums[0:r])
	quickSort(nums[r+1:])

	return
}

func quickSort2(nums []int, start, end int) {
	if len(nums) == 0 || start >= end {
		return
	}
	pivot := nums[(start+end)/2]
	left, right := start, end
	for left <= right {
		for left <= right && nums[left] < pivot {
			left++
		}
		for left <= right && nums[right] > pivot {
			right--
		}
		if left <= right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}

	quickSort2(nums, start, right)
	quickSort2(nums, left, end)

	return
}

// 仅用l和mid实现
// 刷非递归实现
// go 切片底层实现
