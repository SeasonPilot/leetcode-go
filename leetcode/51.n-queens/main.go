package main

// 回溯 + 剪枝  方法一：基于集合的回溯
var solutions [][]string

func solveNQueens(n int) [][]string {
	solutions = [][]string{}
	// 初始化 queens 为-1
	queens := make([]int, n)
	for i := 0; i < n; i++ {
		queens[i] = -1
	}
	// 初始化
	columns := map[int]bool{}
	diagonals1, diagonals2 := map[int]bool{}, map[int]bool{}

	backtrack(queens, n, 0, columns, diagonals1, diagonals2)

	return solutions
}

// row 同时也是递归深度？
func backtrack(queens []int, n, row int, columns, diagonals1, diagonals2 map[int]bool) {
	// terminator
	if row == n { // 递归深度==总皇后个数
		board := generateBoard(queens, n)    // 生成其中一种结果
		solutions = append(solutions, board) // 添加到结果集中
		return
	}

	// 第 row 行时，皇后放在不同列时的情况?
	for i := 0; i < n; i++ {
		// 判断一个位置所在的列和两条斜线上是否已经有皇后。如果不能放置皇后则继续放下个位置(列).    剪枝
		if columns[i] {
			continue
		}
		diagonal1 := row - i // 行下标与列下标之差相等
		if diagonals1[diagonal1] {
			continue
		}
		diagonal2 := row + i // 行下标与列下标之和相等
		if diagonals2[diagonal2] {
			continue
		}

		// process logic in current level 放置皇后
		queens[row] = i
		columns[i] = true
		diagonals1[diagonal1], diagonals2[diagonal2] = true, true

		// drill down
		backtrack(queens, n, row+1, columns, diagonals1, diagonals2)

		// reverse
		queens[row] = -1
		delete(columns, i)
		delete(diagonals1, diagonal1)
		delete(diagonals2, diagonal2)
	}
}

func generateBoard(queens []int, n int) []string {
	board := []string{}
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		// 在该行放置皇后
		row[queens[i]] = 'Q'
		board = append(board, string(row))
	}
	return board
}
