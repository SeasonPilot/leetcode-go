package main

func maxProfit(prices []int) int {
	// 状态
	n := len(prices)
	if n < 1 {
		return 0
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	// 选择 buy sell rest
	// 状态转移
	for i := 0; i < n; i++ {
		// base case
		if i == 0 {
			dp[0][0] = 0
			dp[0][1] = -prices[i] //fixme: -prices[i]
			continue
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], -prices[i]) // fixme: dp[i-1][0]-prices[i])，因为只能交易一次所以这里应该是 0-prices[i]
	}
	return dp[n-1][0]
}

func max(i, i2 int) int {
	if i < i2 {
		return i2
	}
	return i
}
