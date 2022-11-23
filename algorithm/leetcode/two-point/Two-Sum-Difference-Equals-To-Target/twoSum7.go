package Two_Sum_Difference_Equals_To_Target

func TwoSum7(nums []int, target int) []int {
	// write your code here
	if len(nums) < 2 {
		return nil
	}
	r, length := 1, len(nums)
	if target < 0 {
		target = -target
	}
	for l, v := range nums {
		if r < l+1 {
			r = l + 1
		}
		for r < length && nums[r]-v < target {
			r++
		}
		if nums[r]-nums[l] == target {
			return []int{nums[l], nums[r]}
		}
		if r > length {
			break
		}

	}
	return nil
}

func TwoSum7Hash(nums []int, target int) []int {
	// write your code here
	if len(nums) < 2 {
		return nil
	}

	freq := make(map[int]int)
	if target < 0 {
		target = -target
	}
	for _, v := range nums {
		if _, ok := freq[v]; ok {
			return []int{v - target, v}
		}
		sum := target + v
		freq[sum] = 1
	}
	return nil
}

func TwoSum7BinarySearch(nums []int, target int) []int {
	// write your code here
	if len(nums) < 2 {
		return nil
	}
	if target < 0 {
		target = -target
	}
	for k, v := range nums {
		findK := binary_search(nums[k+1:], target+v)
		if findK >= 0 {
			return []int{v, target + v}
		}
	}
	return nil
}

func binary_search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	l, r := 0, len(nums)-1
	for l+1 < r {
		mid := (l + r) / 2
		if nums[mid] > target {
			r = mid
		} else {
			l = mid
		}
	}
	if nums[l] == target {
		return l
	}
	if nums[r] == target {
		return r
	}
	return -1
}
