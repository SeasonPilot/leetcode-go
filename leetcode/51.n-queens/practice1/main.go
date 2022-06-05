package main

// 回溯 + 剪枝  方法一：基于集合的回溯
var solutions [][]string

func solveNQueens(n int) [][]string {
	solutions = [][]string{} // fixme: 在 LeetCode 中执行时会保存之前的结果，所以这里要清空

	queens := make([]int, n)
	for i := 0; i < n; i++ {
		queens[i] = -1
	}
	columns := make(map[int]bool)
	diagonals1, diagonals2 := make(map[int]bool), make(map[int]bool)

	backtrack(queens, n, 0, columns, diagonals1, diagonals2)

	return solutions
}

func backtrack(queens []int, n int, row int, columns, diagonals1, diagonals2 map[int]bool) {
	// terminator
	if row == n {
		board := generateBoard(queens)
		solutions = append(solutions, board)
		return
	}

	// process logic in current level
	// 分别放在不同的列
	for i := 0; i < n; i++ {
		// 剪枝
		if columns[i] {
			continue
		}
		diagonal1 := row - i // 行-列，唯一标识一条斜线？
		if diagonals1[diagonal1] {
			continue
		}
		diagonal2 := row + i
		if diagonals2[diagonal2] { // fixme: diagonals2[diagonal2]
			continue
		}

		// process logic in current level
		queens[row] = i
		columns[i] = true
		diagonals1[diagonal1], diagonals2[diagonal2] = true, true

		// drill down   // n 表示 列
		backtrack(queens, n, row+1, columns, diagonals1, diagonals2)

		// reverse
		queens[row] = -1
		// fixme: map 的 reverse 是删除 key。或者设置成 false 都行。
		columns[i] = false
		diagonals1[diagonal1], diagonals2[diagonal2] = false, false
		//delete(columns, i)
		//delete(diagonals1, diagonal1)
		//delete(diagonals2, diagonal2)
	}

}

func generateBoard(queens []int) []string {
	var board []string
	n := len(queens)

	for i := 0; i < n; i++ {
		//row := ""      //fixme: 无法通过string 的 index 赋值
		//var row []byte //fixme: 数组越界
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		row[queens[i]] = 'Q'
		board = append(board, string(row))
	}
	return board
}
