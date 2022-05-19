package method2

import "strings"

// https://leetcode.cn/problems/n-queens/solution/dai-ma-sui-xiang-lu-51-n-queenshui-su-fa-2k32/
var res [][]string

func isValid(board [][]string, row, col int) (res bool) {
	n := len(board)
	for i := 0; i < row; i++ {
		if board[i][col] == "Q" {
			return false
		}
	}
	for i := 0; i < n; i++ {
		if board[row][i] == "Q" {
			return false
		}
	}

	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == "Q" {
			return false
		}
	}
	for i, j := row, col; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == "Q" {
			return false
		}
	}
	return true
}

func backtrack(board [][]string, row int) {
	size := len(board)
	// terminator
	if row == size {
		temp := make([]string, size)
		for i := 0; i < size; i++ {
			temp[i] = strings.Join(board[i], "")
		}
		res = append(res, temp)
		return
	}

	for col := 0; col < size; col++ {
		if !isValid(board, row, col) { // 剪枝
			continue
		}
		// process logic in current level
		board[row][col] = "Q"
		// drill down
		backtrack(board, row+1)
		// reverse
		board[row][col] = "."
	}
}

func solveNQueens(n int) [][]string {
	res = [][]string{}

	// 初始化
	board := make([][]string, n)
	for i := 0; i < n; i++ {
		board[i] = make([]string, n)
	}
	// 都初始化为 "."
	for i := 0; i < n; i++ { // 行
		for j := 0; j < n; j++ {
			board[i][j] = "."
		}
	}

	backtrack(board, 0)

	return res
}
