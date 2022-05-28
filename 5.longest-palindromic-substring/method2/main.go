package method1

// DP
func longestPalindrome(s string) string {
	n := len(s)

	// 定义状态数组    // dp[i][j] 表示 s[i..j] 是否是回文串
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}

	var length = 1
	var begin = 0

	// 左右边界的关系是 i<=j  =》 j>=i
	// 动态方程
	// i,j 子串左右边界
	for j := 0; j < n; j++ {
		for i := j; i >= 0; i-- {
			dp[i][j] = s[i] == s[j] && (j-i < 3 || dp[i+1][j-1]) // fixme: dp[i+1][j-1] 放在最后可以防止数组越界
			if dp[i][j] && j-i+1 > length {
				length = j - i + 1 // fixme: length 和 begin 要同时变更
				begin = i
			}
		}
	}

	ret := s[begin : begin+length] // fixme: 切片操作左开右闭，不包含后面的字符

	return ret
}
