package sort

// 原址交换nums[left:index] nums[index:right]
func swap(nums []int, left, index, right int) {
	reverse(nums[left:index])
	reverse(nums[index:right])
	reverse(nums[left:right])

}

// 原址翻转
func reverse(arr []int) {
	for l, r := 0, len(arr)-1; l < r; l++ {
		arr[l], arr[r] = arr[r], arr[l]
		r--
	}
}
