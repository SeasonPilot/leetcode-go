package main

import "sort"

//链接：https://leetcode.cn/problems/3sum-closest/solution/gu-ding-yi-ge-shu-zai-shuang-zhi-zhen-shun-bian-fu/
// 排序 + 双指针
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[len(nums)-1]

	for i := 0; i < len(nums)-2; i++ {
		n1 := nums[i]
		l, r := i+1, len(nums)-1
		for l < r {
			n2, n3 := nums[l], nums[r]
			sum := n1 + n2 + n3
			if sum > target {
				r--
			} else {
				l++
			}
			if abs(sum-target) < abs(res-target) {
				res = sum
			}
		}
	}
	return res
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// practice1  错误版本 ❌   排序 + 双指针
func threeSumClosest1(nums []int, target int) int {
	/*
		排序
		固定一个数，双指针

		先将数组从小到大排序，便于微调 sum 的大小。
		从左到右遍历，先固定一个数，剩下的部分，用头尾双指针扫描
		如果 sum 大于目标值，就右指针左移，使 sum 变小，否则左指针右移，sum 变大。
		再看 abs(sum - target) 是否比之前更小了，如果是，将当前 sum 更新给 res

		遍历结束，就有了最接近目标值的 sum
	*/

	sort.Ints(nums)
	n := len(nums)
	sum := nums[0] + nums[1] + nums[2] // 不能使用任意值，要使用数组中的值。  初始化

	for i := 0; i < n-2; i++ {
		n1 := nums[i]
		l, r := nums[i+1], nums[n-1]   // l r 是index, 即左右指针
		for j := i + 1; j < n-1; j++ { // 双指针通过 ++ 或者 -- 移动，不用循环
			for l < r {
				res := n1 + nums[l] + nums[r]
				if res == target {
					return res
				} else if abs(sum-target) > abs(res-target) { // 绝对值小的为结果
					l++
				} else {
					r++
				}
			}
		}
	}
	return sum
}

// practice2   排序 + 双指针
func threeSumClosest2(nums []int, target int) int {
	/*
		1. 排序
		2. 固定一个数，双指针
	*/
	sort.Ints(nums)
	n := len(nums)
	res := nums[0] + nums[1] + nums[n-1] // 初始化，要用数组里面的值

	for i := 0; i < n-2; i++ {
		n1 := nums[i]
		l, r := i+1, n-1 // 初始化左右指针
		for l < r {
			sum := n1 + nums[l] + nums[r] // 临时变量，存储三数之和
			if sum < target {
				l++
			} else {
				r--
			}
			if abs(sum-target) < abs(res-target) { // 如果新结果更优则更新
				res = sum
			}
		}
	}
	return res
}
