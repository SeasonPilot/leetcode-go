package main

// 最终使用 method2 的方法
// flood fill
func numIslands(grid [][]byte) int {
	var islandsNums int

	for y, row := range grid { // fixme: x,y 搞反了，列对应x，行对应y
		for x, b := range row {
			if b == '1' {
				islandsNums += dfs(grid, x, y)
			}
		}
	}

	return islandsNums
}

func dfs(grid [][]byte, x int, y int) int {
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}

	// terminator
	if grid[y][x] == '0' { // fixme: grid[x][y]  第一个 index 对应行；列对应x，行对应y
		return 0
	}

	// process logic in current level
	grid[y][x] = '0'
	for i := 0; i < len(dx); i++ {
		nx, ny := x+dx[i], y+dy[i]

		if nx < len(grid[0]) && nx >= 0 && ny < len(grid) && ny >= 0 { // 判断是否超出边界   // fixme: grid 行和列长度不同
			if grid[ny][nx] == '1' { // fixme: 这里是 nx

				// drill down
				dfs(grid, nx, ny)
			}
		}
	}

	return 1
}
