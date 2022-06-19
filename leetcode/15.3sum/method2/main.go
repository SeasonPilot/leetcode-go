package main

import "sort"

// https://mp.weixin.qq.com/s/fSyJVvggxHq28a0SdmZm6Q
// 方法一：排序 + 双指针  labuladong
func threeSum(nums []int) [][]int {
	return threeSumTarget(nums, 0)
}

/* 计算数组 nums 中所有和为 target 的三元组 */
func threeSumTarget(nums []int, target int) [][]int {
	// 数组得排个序
	sort.Ints(nums)

	n := len(nums)
	ret := make([][]int, 0)

	// 穷举 threeSum 的第一个数
	for i := 0; i < n; i++ {
		// 对 target - nums[i] 计算 twoSum
		tuples := twoSum(nums, i+1, target-nums[i])
		// 如果存在满足条件的二元组，再加上 nums[i] 就是结果三元组
		if tuples != nil {
			for _, tuple := range tuples {
				tuple = append(tuple, nums[i])
				ret = append(ret, tuple)
			}
		}

		// 跳过第一个数字重复的情况，否则会出现重复结果
		for i < n-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return ret
}

// 同 twoSum 题目  leetcode/1.two-sum/nsum-template/main.go:6
/* 从 nums[start] 开始，计算有序数组
 * nums 中所有和为 target 的二元组 */
func twoSum(nums []int, start, target int) [][]int {
	// nums 数组必须有序
	sort.Ints(nums)

	// 左指针改为从 start 开始，其他不变
	lo, hi := start, len(nums)-1
	var res = make([][]int, 0)

	for lo < hi {
		left := nums[lo]
		right := nums[hi]
		sum := nums[lo] + nums[hi]

		if sum < target {
			for lo < hi && nums[lo] == left {
				lo++
			}
		} else if sum > target {
			for lo < hi && nums[hi] == right {
				hi--
			}
		} else {
			res = append(res, []int{left, right})
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
