package main

import "sort"

// labuladong
// https://mp.weixin.qq.com/s/fSyJVvggxHq28a0SdmZm6Q
func twoSum(nums []int, target int) []int {
	// nums 数组必须有序
	sort.Ints(nums)
	// 左右指针
	lo, hi := 0, len(nums)-1
	var res []int

	for lo < hi {
		// 记录索引 lo 和 hi 最初对应的值
		left := nums[lo]
		right := nums[hi]
		sum := nums[lo] + nums[hi]

		// 根据 sum 和 target 的比较，移动左右指针
		if sum < target {
			for lo < hi && nums[lo] == left {
				lo++
			}
		} else if sum > target {
			for lo < hi && nums[hi] == right {
				hi--
			}
		} else {
			res = append(res, left, right)
			// 跳过所有重复的元素,包括自身
			for lo < hi && nums[lo] == left {
				lo++
			}
			for lo < hi && nums[hi] == right {
				hi--
			}
		}
	}
	return res
}
