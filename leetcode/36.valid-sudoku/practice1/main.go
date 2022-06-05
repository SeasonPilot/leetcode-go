package main

func isValidSudoku(board [][]byte) bool {
	// 记录用的容器
	rows, columns := [9][9]int{}, [9][9]int{} // 第一位表示 行号， 第二位表示 出现的数字
	subboxes := [3][3][9]int{}

	for i, row := range board {
		for j, c := range row {
			if c == '.' {
				continue
			}

			// 记录数字出现次数
			number := c - '1'
			rows[i][number]++
			columns[j][number]++ // 第一位表示 列号， 第二位表示 出现的数字
			subboxes[i/3][j/3][number]++

			// 判断
			if rows[i][number] > 1 || columns[j][number] > 1 || subboxes[i/3][j/3][number] > 1 {
				return false
			}
		}
	}
	return true
}
