package practice2

// DP 解法
/*
状态  子串长度 s[i..j]
选择  改变子串长度
状态转移方程   s[i..j] 是否是回文串  dp[i][j] == true
base case    s[i..j]长度 < 4      a    aa   aba
*/
func longestPalindrome(s string) string {
	n := len(s)
	var res string

	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	for j := 0; j < n; j++ {
		for i := j; i >= 0; i-- {

			if j-i+1 < 4 { // 当子串的长度为 2 或者 3 时，不用检查子串是否回文，只需要检查 s[i] 是否等于 s[j]。
				//dp[i][j] = (s[i] == s[j]) && j-i+1 < 4
				dp[i][j] = (s[i] == s[j])
			} else {
				dp[i][j] = (s[i] == s[j]) && dp[i+1][j-1]
			}

			// j-i<3  j-i+1<4 即子串长度len(s[i:j])<4
			//if dp[i][j] && (res == "" || len(res) < j-i+1) {
			if dp[i][j] && j-i+1 > len(res) {
				res = s[i : j+1] //fixme: s[i : j+1]
			}

		}
	}
	return res
}

/*
您好，这是必应。您的搜索输入是关于回文子串的问题。回文子串是指正着读和倒过来读一样的字符串，例如“aba”或“abba”。要判断一个子串是否是回文子串，有一种常用的方法是使用动态规划

动态规划的思路是，用一个二维数组 dp[i][j] 表示子串 s[i...j] 是否是回文子串，如果是，dp[i][j] 为 true，否则为 false。然后从小到大枚举子串的长度，根据状态转移方程更新 dp[i][j] 的值。状态转移方程如下：

- 如果 s[i] != s[j]，那么 dp[i][j] 一定为 false。
- 如果 s[i] == s[j]，那么还要看 s[i+1...j-1] 是否是回文子串，也就是 dp[i+1][j-1] 的值。如果是，那么 dp[i][j] 也为 true，否则为 false。
- 但是有一个特殊情况，就是当子串的长度为 2 或者 3 时，不用检查子串是否回文，只需要检查 s[i] 是否等于 s[j]。例如：wqw，只需要检查 首尾的两个字母 是否相同。

所以，状态转移方程可以写成：

```java
dp[i][j] = (s[i] == s[j]) && (j - i < 2 || dp[i+1][j-1])
```
*/
