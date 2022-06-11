package main

func maxProfit(prices []int) int {
	n := len(prices)
	if n < 1 {
		return 0
	}
	// 状态 天数、持有状态、冷冻期
	// 选择 buy sell rest
	// 状态转移
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 2)
	}

	for i := 0; i < n; i++ {
		// base case 1
		if i-1 == -1 { //
			dp[0][0] = 0
			dp[0][1] = -prices[i]
			continue
		}
		// base case 2
		if i-2 == -1 { //
			dp[1][0] = max(dp[0][0], dp[0][1]+prices[i])
			dp[1][1] = max(dp[0][0]-prices[i], dp[0][1])
			continue
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-2][0]-prices[i], dp[i-1][1]) // 每次 sell 之后要等一天才能继续交易
	}
	return dp[n-1][0]
}

func max(i, i2 int) int {
	if i < i2 {
		return i2
	}
	return i
}
