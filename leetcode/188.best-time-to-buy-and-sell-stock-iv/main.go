package main

func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n < 1 {
		return 0
	}
	// 状态 天数、交易次数、持有状态
	dp := make([][][2]int, n)
	for i := range dp {
		dp[i] = make([][2]int, k+1)
	}
	// 选择 buy sell reset
	// 状态转移
	for i := 0; i < n; i++ {
		for c := 1; c < k+1; c++ {
			// base case
			if i-1 == -1 {
				dp[i][c][0] = 0 // fixme: dp[i][c][0] index 一定要是变量，不能是 0. 含义：第 0 天不管交易次数是多少不持有股票的收益都是 0
				dp[i][c][1] = -prices[i]
				continue // fixme:
			}
			dp[i][c][0] = max(dp[i-1][c][0], dp[i-1][c][1]+prices[i])
			dp[i][c][1] = max(dp[i-1][c-1][0]-prices[i], dp[i-1][c][1]) // 状态转移方程中交易以买入为准，买入后交易增加一次，卖出交易次数不变。
		}
	}
	return dp[n-1][k][0]
}

func max(i, i2 int) int {
	if i < i2 {
		return i2
	}
	return i
}
