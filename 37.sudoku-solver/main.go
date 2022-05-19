package main

// 回溯
func solveSudoku(board [][]byte) {
	var line, column [9][9]bool
	var block [3][3][9]bool
	var spaces [][2]int // 记录空格,第二个数组[2]表示坐标

	// 扫描棋盘
	for i, row := range board {
		for j, b := range row { // 如果该位置是一个空白格，那么我们将其加入一个用来存储空白格位置的列表中，方便后续的递归操作；
			if b == '.' {
				spaces = append(spaces, [2]int{i, j})
			} else { // 如果该位置是一个数字 x，那么我们需要将 line[i][x−1]，column[j][x−1] 以及 block[⌊i/3⌋][⌊j/3⌋][x−1] 均置为 True。
				digit := b - '1'
				line[i][digit] = true
				column[j][digit] = true
				block[i/3][j/3][digit] = true
			}
		}
	}

	// dfs define
	var dfs func(int) bool
	dfs = func(pos int) bool {
		// terminator
		if pos == len(spaces) {
			return true
		}

		i, j := spaces[pos][0], spaces[pos][1]
		for digit := byte(0); digit < 9; digit++ {
			if !line[i][digit] && !column[j][digit] && !block[i/3][j/3][digit] { // valid-sudoku。   必须均为 False
				// process logic in current level
				line[i][digit] = true
				column[j][digit] = true
				block[i/3][j/3][digit] = true
				board[i][j] = digit + '1'

				//drill down   并且继续对下一个空白格位置进行递归
				if dfs(pos + 1) {
					return true
				}

				// reverse  在回溯到当前递归层时，我们还要将上述的三个值重新置为 False。
				line[i][digit] = false
				column[j][digit] = false
				block[i/3][j/3][digit] = false
			}
		}

		return false
	}

	// function call
	dfs(0)
}
