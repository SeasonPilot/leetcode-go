package main

//链接：https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii/solution/dong-tai-gui-hua-php-by-salmonl-13/
func maxProfit(prices []int) int {
	count := 2
	// 交易次数len为count + 1, 因为包括0
	// 不能把count写在2的位置，提示Go error: non-constant array bound, 数组的len不能是常量
	dp := make([][3][2]int, len(prices))

	for i := 0; i < len(prices); i++ {
		for j := count; j >= 1; j-- {
			if i == 0 {
				dp[i][j][0] = 0
				dp[i][j][1] = -prices[i]
				continue
			}

			dp[i][j][0] = maxInt(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			dp[i][j][1] = maxInt(dp[i-1][j-1][0]-prices[i], dp[i-1][j][1])
		}
	}

	return dp[len(prices)-1][count][0]
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
