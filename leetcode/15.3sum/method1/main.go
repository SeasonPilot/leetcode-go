package main

import "sort"

// 方法一：排序 + 双指针
func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	n := len(nums)
	ret := make([][]int, 0)

	for i := 0; i < n-2; i++ {
		lo, hi := i+1, n-1 // fixme: 每次循环要重置 lo, hi

		for lo < hi {
			// 记录索引 lo 和 hi 最初对应的值
			left := nums[lo] // 存储下
			right := nums[hi]

			if -nums[i] > nums[lo]+nums[hi] { // fixme: 应该是 -nums[i]    0-nums[i] > nums[lo]+nums[hi]
				lo++
			} else if -nums[i] < nums[lo]+nums[hi] {
				hi--
			} else {
				ret = append(ret, []int{nums[i], nums[lo], nums[hi]})
				// 跳过所有重复的元素,包括自身
				for lo < hi && nums[lo] == left { // 添加结果后，要移动 左右 指针
					lo++
				}
				for lo < hi && nums[hi] == right {
					hi--
				}
			}
		}

		// 跳过第一个数字重复的情况，否则会出现重复结果
		for ; i < n-1 && nums[i] == nums[i+1]; i++ { // i < n-1 而不是 i < n， 是因为 nums[i+1] 数组会越界，且 i 是从 0 开始。
		}
	}

	return ret
}
