package main

func maxProfit(prices []int, fee int) int {
	n := len(prices)
	if n < 1 { // fixme: 边界条件
		return 0
	}
	//状态 天数、持有状态
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	//选择 buy sell rest
	//状态转移
	for i := 0; i < n; i++ {
		//base case
		if i-1 == -1 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue // fixme: continue
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee) // 卖出时要收手续费
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
