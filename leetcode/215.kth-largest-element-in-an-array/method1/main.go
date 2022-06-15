package method1

import (
	"sort"
)

// 使用内置函数
func findKthLargest(nums []int, k int) int {
	//sort.SliceStable(nums, func(i, j int) bool {
	//	if nums[i] > nums[j] { // fixme:  这里不是 i < j
	//		return true
	//	}
	//	return false
	//})

	// 优化后
	sort.SliceStable(nums, func(i, j int) bool { return nums[i] > nums[j] })

	return nums[k-1]
}
