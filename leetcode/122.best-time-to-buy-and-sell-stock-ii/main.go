package main

func maxProfit(prices []int) int {
	n := len(prices)
	if n < 1 {
		return 0
	}
	// 状态 状态就是会变的量。 天数、持有股票数、持有状态
	dp := make([][]int, n)
	for i := range dp { // fixme: 初始化二维切片
		dp[i] = make([]int, 2)
	}
	// 选择 选择就是导致状态变化的动作。 buy, sell, rest
	// 状态转移
	for i := 0; i < n; i++ {
		if i-1 == -1 {
			dp[0][0] = 0
			dp[0][1] = -prices[i]
			continue // fixme:
		}
		// base case
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][0]-prices[i], dp[i-1][1])
	}
	return dp[n-1][0]
}

func max(i, i2 int) int {
	if i < i2 {
		return i2
	}
	return i
}
