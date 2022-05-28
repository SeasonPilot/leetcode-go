package main

func longestPalindrome(s string) string {
	// 分治
	// 定义状态数组
	ret, dp := "", make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	for j := 0; j < len(s); j++ {
		for i := j; i >= 0; i-- {
			// DP方程
			//dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
			dp[i][j] = s[i] == s[j] && (j-i < 3 || dp[i+1][j-1])
			// fixme: 边界条件 (j-1) - (i+1) +1 < 2  子串长度小于 2,即 [i+1][j-1] 的长度 为 1 或 0 时不用检查是否回文
			//  边界条件：s[i,j] 长度为 2 或者 3 时，不用检查子串是否回文，只需要检查 s[i] 是否等于 s[j]。例如：wqw，只需要检查 首尾的两个字母 是否相同

			if dp[i][j] && (ret == "" || j-i+1 > len(ret)) {
				ret = s[i : j+1] // 更新结果
			}
		}
	}
	return ret
}
