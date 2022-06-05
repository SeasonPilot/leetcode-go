package main

// 回溯
func solveSudoku(board [][]byte) {
	var lines, columns [9][9]bool // 二维数组不需要初始化？
	var block [3][3][9]bool
	var spaces [][2]int // 记录空格,第二个数组[2]表示坐标

	// 扫描棋盘
	for i, line := range board {
		for j, b := range line {
			if b == '.' {
				spaces = append(spaces, [2]int{i, j})
			} else {
				//
				digit := b - '1' // fixme: 1 应该是码点
				lines[i][digit] = true
				columns[j][digit] = true
				block[i/3][j/3][digit] = true
			}
		}
	}

	var dfs func(pos int) bool
	dfs = func(pos int) bool { // pos 为递归深度，空格序号
		// terminator
		if pos == len(spaces) {
			return true
		}
		// process logic in current
		i, j := spaces[pos][0], spaces[pos][1] // fixme: 获取当前空格坐标位置
		for k := byte(0); k < 9; k++ {         // 向空格中填写数字
			if !lines[i][k] && !columns[j][k] && !block[i/3][j/3][k] { //fixme: 这里是并且的关系
				lines[i][k] = true
				columns[j][k] = true
				block[i/3][j/3][k] = true
				board[i][j] = k + '1' // 把数字填入棋盘  // fixme: 1 应该是码点。 byte 可以和 rune 进行计算

				// drill down
				if dfs(pos + 1) {
					return true
				}

				// reverse
				lines[i][k] = false
				columns[j][k] = false
				block[i/3][j/3][k] = false
				//board[i][j] = '.' // 把.填入棋盘 // fixme: 好像不需要这行代码？ 已经设置为 false，是根据 false 来判断的。
			}
		}

		return false
	}

	dfs(0)
}
