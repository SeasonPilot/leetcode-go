package main

import "fmt"

// 最终使用此方法
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	if m*n == 0 {
		return m + n
	}

	// 二维数组 第一个表示 word1 长度， 第二个表示 word2 长度，
	dp := make([][]int, m+1) // 字符串长度可以为 0，所以切片长度是 m + 1

	for i := range dp { // 初始化二维切片
		dp[i] = make([]int, n+1)
	}

	// base case
	for i := 0; i < m+1; i++ { // 填充第一列
		dp[i][0] = i // word1[i] 变成 word2[0], 删掉 word1[i], 需要 i 部操作
	}
	for j := 0; j < n+1; j++ { // 填充第一行
		dp[0][j] = j // word1[0] 变成 word2[j], 插入 word1[j]，需要 j 部操作
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {

			if word1[i-1] == word2[j-1] { // word1 第一个字符， 长度为 1，所以是 word1[i-1]， dp[i]
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = Min(
					dp[i-1][j-1]+1, // 替换； 最后一位， x 变为 y,或者 y 变 x
					dp[i-1][j]+1,   // 删除 x
					dp[i][j-1]+1,   // 删除 y
				)
			}

		}
	}

	return dp[m][n]
}

func Min(args ...int) int {
	min := args[0]
	for _, item := range args {
		if item < min {
			min = item
		}
	}
	return min
}

func main() {
	word1 := "horse"
	word2 := "ros"
	fmt.Println(minDistance(word1, word2))
}
