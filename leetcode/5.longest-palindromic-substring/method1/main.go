package method1

// 最终使用此方法
// DP
func longestPalindrome(s string) string {
	n := len(s)
	var res string

	// 定义状态数组    // dp[i][j] 表示 s[i..j] 是否是回文串
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}

	// 左右边界的关系是 i<=j
	// i,j 子串左右边界   递推顺序: [0,0] [1,1][0,1] [2,2][1,2][0,2]
	// 注意：在状态转移方程中，我们是从长度较短的字符串向长度较长的字符串进行转移的，因此一定要注意动态规划的循环顺序。
	// 动态方程
	for j := 0; j < n; j++ { // 右边界向右移动,所以右边界初始位置在最左边
		for i := j; i >= 0; i-- { // 左边界向左移动，逐渐扩大字符串
			dp[i][j] = s[i] == s[j] && (j-i < 3 || dp[i+1][j-1]) // fixme: dp[i+1][j-1] 放在最后可以防止数组越界
			if dp[i][j] && (res == "" || j-i+1 > len(res)) {
				res = s[i : j+1]
			}
		}
	}

	return res
}
