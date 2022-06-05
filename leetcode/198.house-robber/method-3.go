package main

func rob3(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	first := nums[0]
	if n == 1 {
		return first
	}
	second := max(first, nums[1])

	for i := 2; i < n; i++ {
		//dp = max(first+nums[i], second) // 这里不是 nums[2]
		//first = second
		//second = dp

		// 优化后 变量平移
		first, second = second, max(first+nums[i], second)
	}

	return second
}
