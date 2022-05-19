package main

// 使用一维数组 定义状态数组
func rob1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	dp := make([]int, n)
	dp[0] = nums[0]

	if n < 2 {
		return dp[n-1]
	}
	// 因为这里要访问第二个元素，所以要提前判断 n 是否是 <2
	dp[1] = max(nums[0], nums[1])

	/*
	   	如果房屋数量大于两间，应该如何计算能够偷窃到的最高总金额呢？对于第 k(k>2) 间房屋，有两个选项：

	      偷窃第 k 间房屋，那么就不能偷窃第 k−1 间房屋，偷窃总金额为前 k−2 间房屋的最高总金额与第 k 间房屋的金额之和。

	      不偷窃第 k 间房屋，偷窃总金额为前 k−1 间房屋的最高总金额。
	*/
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[n-1]
}
