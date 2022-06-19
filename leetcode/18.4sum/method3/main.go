package main

import "sort"

// https://mp.weixin.qq.com/s/fSyJVvggxHq28a0SdmZm6Q
// 方法：排序 + 双指针  合并 labuladong 方法的多个函数
func fourSum(nums []int, target int) [][]int {
	// 排序
	sort.Ints(nums)

	n := len(nums)
	ret := make([][]int, 0)

	// 穷举 fourSum 的第一个数
	for i := 0; i < n-3; i++ {

		for j := i + 1; j < n-2; j++ {
			// 对 target - nums[i] - nums[j] 计算 twoSum

			lo, hi := j+1, len(nums)-1

			for lo < hi {
				left := nums[lo]
				right := nums[hi]
				sum := nums[lo] + nums[hi]

				if sum < target-nums[i]-nums[j] {
					for lo < hi && nums[lo] == left {
						lo++
					}
				} else if sum > target-nums[i]-nums[j] {
					for lo < hi && nums[hi] == right {
						hi--
					}
				} else {
					ret = append(ret, []int{nums[i], nums[j], left, right})
					// 跳过所有重复的元素,包括自身
					for lo < hi && nums[lo] == left {
						lo++
					}
					for lo < hi && nums[hi] == right {
						hi--
					}
				}
			}

			// 跳过第一个数字重复的情况，否则会出现重复结果
			for j < n-2 && nums[j] == nums[j+1] {
				j++
			}
		}

		// fourSum 的第一个数不能重复
		for i < n-3 && nums[i] == nums[i+1] {
			i++
		}
	}
	return ret
}
