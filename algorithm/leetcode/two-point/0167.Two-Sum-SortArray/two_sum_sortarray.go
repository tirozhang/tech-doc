package _167_Two_Sum_SortArray

func twoSumSortArray(nums []int, target int) []int {

	for i, j := 0, len(nums)-1; i < j; {
		if nums[i]+nums[j] < target {
			i++
		}

		if nums[i]+nums[j] > target {
			j--
		}

		if nums[i]+nums[j] == target {
			return []int{i + 1, j + 1}
		}
	}
	return nil
}

func twoSumSortArray2(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		sum := numbers[l] + numbers[r]
		switch {
		case sum > target:
			r--
		case sum < target:
			l++
		case sum == target:
			return []int{l + 1, r + 1}
		}
	}
	return []int{-1, -1}
}
