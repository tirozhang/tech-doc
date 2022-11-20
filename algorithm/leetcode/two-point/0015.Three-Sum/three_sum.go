package _015_Three_Sum

import (
	"sort"
	"strconv"
)

// O(n^2)
func threeSum(nums []int) [][]int {
	sumMap := make(map[int]int)
	isFind := make(map[string]int)
	var res [][]int
	for k, v := range nums {
		sumMap[-v] = k
	}
	l, length := 0, len(nums)
	for l < length {
		r := l + 1
		for r < length {
			if r == l {
				continue
			}
			if index, ok := sumMap[nums[l]+nums[r]]; ok && index != l && index != r {
				find := []int{nums[l], nums[r], -nums[l] - nums[r]}
				sort.Ints(find)
				hashKey := strconv.Itoa(find[0]) + strconv.Itoa(find[1]) //robin key
				if _, ok := isFind[hashKey]; !ok {
					isFind[hashKey] = 1
					res = append(res, find)
				}
			}
			r++
		}
		l++
	}
	return res
}

// 双指针 + 排序
func threeSum2(nums []int) [][]int {
	sort.Ints(nums)
	result, index, addNum, length := make([][]int, 0), 0, 0, len(nums)
	for index = 1; index < length-1; index++ {
		start, end := 0, length-1
		if index > 1 && nums[index] == nums[index-1] {
			start = index - 1
		}
		for start < index && end > index {
			if start > 0 && nums[start] == nums[start-1] {
				start++
				continue
			}
			if end < length-1 && nums[end] == nums[end+1] {
				end--
				continue
			}
			addNum = nums[start] + nums[end] + nums[index]
			if addNum == 0 {
				result = append(result, []int{nums[start], nums[index], nums[end]})
				start++
				end--
			} else if addNum > 0 {
				end--
			} else {
				start++
			}
		}
	}
	return result
}

// twoSum+sort
func threeSum3(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for k, v := range nums {
		if k != 0 && v == nums[k-1] {
			continue
		}
		findTwoNum := twoSumSortArray(nums[k+1:], -v)
		for _, r := range findTwoNum {
			r = append(r, v)
			res = append(res, r)
		}
	}
	return res
}

func twoSumSortArray(numbers []int, target int) [][]int {
	l, r := 0, len(numbers)-1
	var res [][]int
	for l < r {
		if l > 0 && numbers[l] == numbers[l-1] {
			l++
			continue
		}
		sum := numbers[l] + numbers[r]
		switch {
		case sum > target:
			r--
		case sum < target:
			l++
		case sum == target:
			res = append(res, []int{numbers[l], numbers[r]})
			r--
			l++
		}
	}
	return res
}
