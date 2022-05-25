package main

import "fmt"

func numIslands(grid [][]byte) int {
	var islands int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				sink(grid, i, j)
				islands++
				fmt.Printf("%c\n", grid)
			}
		}
	}

	return islands
}

func sink(grid [][]byte, i int, j int) {
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}
	//terminator
	if grid[i][j] == '0' {
		return
	}

	//process logic in current level
	grid[i][j] = '0'
	for k := 0; k < len(dx); k++ {
		x, y := i+dx[k], j+dy[k]
		if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[i]) {
			if grid[x][y] == '1' { // fixme: grid[x][y] 方括号内的值为新的 index

				//drill down
				sink(grid, x, y)
			}
		}
	}
	//reverse
}
