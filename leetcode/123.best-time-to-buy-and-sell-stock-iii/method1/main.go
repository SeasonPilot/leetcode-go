package method1

func maxProfit(prices []int) int {
	n := len(prices)
	max_k := 2 // 交易次数 len 为 max_k + 1, 因为包括 0

	if n < 1 {
		return 0
	}
	// 状态 天数、交易次数、持有状态
	dp := make([][][]int, n) // dp[i][k][0]表示截止到第i天进行了k次交易，目前不持有股票；dp[i][k][1]表示截止到第i天进行了k次交易，目前持有股票.
	for i := range dp {
		dp[i] = make([][]int, max_k+1) // 交易次数 len 为 max_k + 1, 因为包括 0
		for i2 := range dp[i] {
			dp[i][i2] = make([]int, 2)
		}
	}
	// 选择 buy sell rest

	// 状态转移
	for i := 0; i < n; i++ { // 0 <= i <= n - 1
		for k := 1; k < max_k+1; k++ { // 1 <= k <= max_k; max_k 为交易数的上限
			// base case
			if i-1 == -1 {
				dp[i][k][0] = 0
				dp[i][k][1] = -prices[i]
				continue
			}
			//dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			//dp[i][k][1] = max(dp[i-1][k+1][0]-prices[i], dp[i-1][k][1]) // 最终结果是错误的。 第i天的k一定是大于等于第i-1天的，即第i天的交易次数一定是大于等于第i-1天的，即交易次数是递增的。 // k+1 会导致数组越界  卖出时交易次数+1。

			// 1、当天进行了k次交易，且不持有股票的最大收益为两种情况取大者：一、前一天交易了k次本身也不持有股票；二、前一天进行了k次交易，并持有股票，今天卖了。交易发生在当天，所以前一天的交易次数是k, 不是k - 1。
			// reset， 卖出
			dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			// 2、当天进行了k次交易，并持有股票的最大收益为两种情况取大者：一、前一天交易了k次本身持有股票；二、前一天进行了最多k - 1次交易，并不持有股票，当天买入了。当天还能买入，说明之前交易次数没有达到k次。
			// 买入， rest
			// 注意：个人理解，状态转移方程中交易以买入为准，买入后交易增加一次，卖出交易次数不变。
			dp[i][k][1] = max(dp[i-1][k-1][0]-prices[i], dp[i-1][k][1]) // 交易次数是递增的。 第i天交易次数是k,那么第i-1天交易次数就是k-1。得留一次交易次数给第i天  // fixme: k-1 只能写在这
		}
	}
	return dp[n-1][max_k][0]

}

// 参考： https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii/solution/dong-tai-gui-hua-php-by-salmonl-13/

func max(i, i2 int) int {
	if i < i2 {
		return i2
	}
	return i
}
