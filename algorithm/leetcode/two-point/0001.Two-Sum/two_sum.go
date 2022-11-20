package _001_Two_Sum

import (
	"sort"
)

// hash表实现 O(n)
func twoSum(nums []int, target int) []int {
	// 通过hash表来存储遍历过的数字
	m := make(map[int]int, len(nums))
	for k, num := range nums {
		another := target - num
		if i, ok := m[another]; ok {
			return []int{i, k}
		}
		m[num] = k
	}
	return []int{-1, -1}
}

// 双指针实现 O(nlogn)
func twoSum2(nums []int, target int) []int {
	notFound := []int{-1, -1}
	if len(nums) == 0 {
		return notFound
	}

	indexSite := make(map[int]int, len(nums))
	for k, v := range nums {
		indexSite[v] = k
	}

	sort.Ints(nums)
	l, r := 0, len(nums)-1
	for l < r {
		sum := nums[l] + nums[r]
		switch {
		case sum > target:
			r--
		case sum < target:
			l++
		case sum == target:
			return []int{indexSite[nums[l]], indexSite[nums[r]]}
		}
	}

	return notFound
}
