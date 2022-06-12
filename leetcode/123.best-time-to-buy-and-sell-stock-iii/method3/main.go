package main

import "math"

func maxProfit(prices []int) int {
	//dp[i][k][0],i=-1,k=0
	dp_i_0 := []int{0, 0, 0}
	//dp[i][k][1],i=-1,k=0
	dp_i_1 := []int{math.MinInt64, math.MinInt64, math.MinInt64}
	for _, v := range prices {
		for k := 1; k <= 2; k++ {
			if dp_i_1[k]+v > dp_i_0[k] {
				dp_i_0[k] = dp_i_1[k] + v
			}
			if dp_i_0[k-1]-v > dp_i_1[k] {
				dp_i_1[k] = dp_i_0[k-1] - v
			}
		}
	}
	return dp_i_0[2]
}

//链接：https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii/solution/dong-tai-gui-hua-zhuang-tai-ji-oke-wai-kong-jian-b/

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := make([][3][2]int, len(prices))

	// k = 2

	for i := 0; i < len(prices); i++ {
		// k可以++，也可以--。还是统一++
		for k := 1; k <= 2; k++ {
			if i == 0 {
				dp[i][k][0] = 0
				dp[i][k][1] = -prices[i]
				continue
			}
			// 为什么不是k-1   dp[i-1][k][1]+prices[i]
			dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
		}
	}
	return dp[len(prices)-1][2][0]
}
func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}

//链接：https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii/solution/by-xiaosheng-xs-333e/
