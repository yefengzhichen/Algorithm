package main

import "fmt"

func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	row, col := len(grid), len(grid[0])
	res, cur := 0, 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			cur = process(grid, i, j, row, col)
			if res < cur {
				res = cur
			}
		}
	}
	return res
}

func process(grid [][]int, i, j int, row, col int) int {
	if grid[i][j] == 0 || grid[i][j] == -1 {
		return 0
	}
	count := 1
	grid[i][j] = -1
	if i-1 >= 0 {
		count += process(grid, i-1, j, row, col)
	}
	if j-1 >= 0 {
		count += process(grid, i, j-1, row, col)
	}
	if i+1 < row {
		count += process(grid, i+1, j, row, col)
	}
	if j+1 < col {
		count += process(grid, i, j+1, row, col)
	}
	return count
}

func main() {
	fmt.Println("a")
}
