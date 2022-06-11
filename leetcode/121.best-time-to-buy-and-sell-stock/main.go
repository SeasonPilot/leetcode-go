package main

func maxProfit(prices []int) int {
	// 状态： 状态就是会变的量。 第一个是天数，第二个是允许交易的最大次数，第三个是当前的持有状态
	n := len(prices) // 一共是 1到n 天,左移后是 0 <= i <= n - 1

	if n < 1 {
		return 0
	}
	// 明确 DP 数组的定义,状态的个数决定了DP数组的维数。根据这个定义，我们可以推导出base case以及我们最终想计算的结果
	dp := make([][]int, n)
	for i := range dp { // 初始化二维切片
		dp[i] = make([]int, 2)
	}

	// 选择： 选择就是导致状态变化的动作。  买入、卖出、无操作， 用 buy, sell, rest 表示这三种选择。
	// 状态迁移 根据我们的 选择 去写出状态转移的逻辑。
	for i := 0; i < n; i++ {
		// base case
		if i-1 == -1 {
			dp[0][0] = 0
			dp[0][1] = -prices[i] // 第一天买入
			continue
		}
		// reset, i-1 买入，i 卖出
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]) // 今天选择 reset ，今天选择 sell //fixme: prices数组越界了，天数是从 1 开始，prices 是从 0 开始
		// 继续持有， i 天买入
		dp[i][1] = max(dp[i-1][1], -prices[i]) // 今天选择 reset，今天选择 buy
	}
	return dp[n-1][0] // 最大收益为最后一天卖出后的收益
}

func max(i, i2 int) int {
	if i < i2 {
		return i2
	}
	return i
}

/*
	DP 方程是对称的

	dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][1][1] + prices[i])
	dp[i][1][1] = max(dp[i-1][1][1], dp[i-1][0][0] - prices[i])

	dp[i][k][0 or 1]
	0 <= i <= n - 1, 1 <= k <= K
	n 为天数，大 K 为交易数的上限，0 和 1 代表是否持有股票。
	此问题共 n × K × 2 种状态，全部穷举就能搞定。
*/
