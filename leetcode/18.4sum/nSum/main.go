package main

/*
	注意：调用这个函数之前一定要先给 nums 排序
	因为 nSum 是一个递归函数，如果在 nSum 函数里调用排序函数，那么每次递归都会进行没有必要的排序，效率会非常低。
*/
func nSumTarget(nums []int, n, start, target int) [][]int {
	sz := len(nums)
	res := make([][]int, 0)

	// 至少是 2Sum，且数组大小不应该小于 n
	if n < 2 || sz < n {
		return res
	}

	// 2Sum 是 base case
	if n == 2 {
		// 双指针那一套操作
		lo, hi := start, sz-1 // fixme: lo 从 start 开始

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

	} else {
		// n > 2 时，递归计算 (n-1)Sum 的结果
		for i := start; i < sz; i++ {
			sub := nSumTarget(nums, n-1, i+1, target-nums[i])

			if sub != nil {
				for _, tuple := range sub {
					// (n-1)Sum 加上 nums[i] 就是 nSum
					tuple = append(tuple, nums[i])
					res = append(res, tuple)
				}
			}

			for i < sz-1 && nums[i] == nums[i+1] {
				i++
			}
		}
	}

	return res
}
