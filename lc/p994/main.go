package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func orangesRotting(grid [][]int) int {
	if len(grid) < 1 || len(grid[0]) < 1 {
		return 0
	}
	arrx := make([]int, 0)
	arry := make([]int, 0)
	m, n := len(grid), len(grid[0])
	countNew := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				arrx = append(arrx, i)
				arry = append(arry, j)
			} else if grid[i][j] == 1 {
				countNew += 1
			}
		}
	}
	if len(arrx) == 0 && countNew == 0 {
		return 0
	}
	minute := -1
	for len(arrx) > 0 {
		nextx := make([]int, 0)
		nexty := make([]int, 0)
		minute += 1
		for i := 0; i < len(arrx); i++ {
			x := arrx[i]
			y := arry[i]
			if x-1 >= 0 && grid[x-1][y] == 1 {
				grid[x-1][y] = 2
				nextx = append(nextx, x-1)
				nexty = append(nexty, y)
			}
			if x+1 < m && grid[x+1][y] == 1 {
				grid[x+1][y] = 2
				nextx = append(nextx, x+1)
				nexty = append(nexty, y)
			}
			if y-1 >= 0 && grid[x][y-1] == 1 {
				grid[x][y-1] = 2
				nextx = append(nextx, x)
				nexty = append(nexty, y-1)
			}
			if y+1 < n && grid[x][y+1] == 1 {
				grid[x][y+1] = 2
				nextx = append(nextx, x)
				nexty = append(nexty, y+1)
			}
		}
		arrx = nextx
		arry = nexty
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}
	return minute
}

func orangesRotting2(grid [][]int) int {
	if len(grid) < 1 || len(grid[0]) < 1 {
		return 0
	}
	arrx := make([]int, 0)
	arry := make([]int, 0)
	m, n := len(grid), len(grid[0])
	countNew := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				arrx = append(arrx, i)
				arry = append(arry, j)
			} else if grid[i][j] == 1 {
				countNew += 1
			}
		}
	}
	if len(arrx) == 0 && countNew == 0 {
		return 0
	}
	minute := -1
	fx := []int{-1, 1, 0, 0}
	fy := []int{0, 0, -1, +1}
	for len(arrx) > 0 {
		nextx := make([]int, 0)
		nexty := make([]int, 0)
		minute += 1
		for i := 0; i < len(arrx); i++ {
			for j := 0; j < 4; j++ {
				x := arrx[i] + fx[j]
				y := arry[i] + fy[j]
				if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 1 {
					grid[x][y] = 2
					nextx = append(nextx, x)
					nexty = append(nexty, y)
				}
			}
		}
		arrx = nextx
		arry = nexty
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}
	return minute
}

func main() {
	fmt.Println("a")
}
