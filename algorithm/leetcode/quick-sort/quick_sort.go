package quick_sort

// 快速排序
func quickSort(nums []int) {
	if len(nums) <= 1 {
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

// 仅用l和mid实现
// 刷非递归实现
// go 切片底层实现
