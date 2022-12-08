package help_func

// Swap 原址交换nums[left:index] nums[index:right]
func Swap(nums []int, left, index, right int) {
	Reverse(nums[left:index])
	Reverse(nums[index:right])
	Reverse(nums[left:right])

}

// Reverse 原址翻转
func Reverse(arr []int) {
	for l, r := 0, len(arr)-1; l < r; l++ {
		arr[l], arr[r] = arr[r], arr[l]
		r--
	}
}
