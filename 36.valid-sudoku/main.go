package main

// LeetCode
func isValidSudoku(board [][]byte) bool {
	var rows, columns [9][9]int // 计数
	var subboxes [3][3][9]int

	for i, row := range board {
		for j, c := range row {
			if c == '.' {
				continue
			}

			index := c - '1'
			rows[i][index]++
			columns[j][index]++
			subboxes[i/3][j/3][index]++ // [i/3][j/3] 表示 subboxes 的 坐标

			if rows[i][index] > 1 || columns[j][index] > 1 || subboxes[i/3][j/3][index] > 1 {
				return false
			}
		}
	}
	return true
}

// method2 by me
func isValidSudoku2(board [][]byte) bool {
	//var rows, columns, block []map[byte]int  //fixme: 如果要直接访问索引，需要对切片进行初始化

	n := len(board)
	rows, columns, block := make([]map[byte]int, n), make([]map[byte]int, n), make([]map[byte]int, n)

	for i := range rows {
		rows[i] = make(map[byte]int)
	}

	for i := range columns {
		columns[i] = make(map[byte]int)
	}

	for i := range block {
		block[i] = make(map[byte]int)
	}

	//ma := map[byte]int{} //fixme: map 是引用类型，所以 rows, columns, block 的值都指向同一个地址( ma 的地址)
	//rows[0] = ma
	//columns[0] = ma
	//block[0] = ma

	//fixme: rows 已经初始化为长度为 n 的切片。append 是追加到切片最后面，不会改变 index 0 的值。
	// 如果切片长度为 0 则 append 后 index0 是有值的。
	//rows = append(rows, ma)
	//columns = append(columns, ma)
	//block = append(block, ma)

	for i, row := range board {
		for j, c := range row {
			if c == '.' {
				continue
			}

			//
			rows[i][c]++
			columns[j][c]++
			block[(i/3)*3+j/3][c]++

			//
			if rows[i][c] > 1 || columns[j][c] > 1 || block[(i/3)*3+j/3][c] > 1 {
				return false
			}
		}
	}
	return true
}

// method3 by me
func isValidSudoku3(board [][]byte) bool {
	// 初始化 map 类型的 slice ; different from method2
	n := len(board)
	rows, columns, block := make([]map[byte]bool, n), make([]map[byte]bool, n), make([]map[byte]bool, n)

	for i := range rows {
		rows[i] = make(map[byte]bool)
	}

	for i := range columns {
		columns[i] = make(map[byte]bool)
	}

	for i := range block {
		block[i] = make(map[byte]bool)
	}

	for i, row := range board {
		for j, c := range row {
			if c == '.' {
				continue
			}

			// different from method2
			if rows[i][c] || columns[j][c] || block[(i/3)*3+j/3][c] {
				return false
			}
			//
			rows[i][c] = true
			columns[j][c] = true
			block[(i/3)*3+j/3][c] = true

		}
	}
	return true
}
