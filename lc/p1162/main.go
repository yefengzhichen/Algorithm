package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDistance(grid [][]int) int {
	ax := make([]int, 0)
	ay := make([]int, 0)
	count, m, n := 0, len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				ax = append(ax, i)
				ay = append(ay, j)
			} else {
				count += 1
			}
		}
	}
	if len(ax) == 0 || count == 0 {
		return -1
	}
	res := 0
	px := []int{-1, 1, 0, 0}
	py := []int{0, 0, -1, 1}
	for len(ax) > 0 {
		nextx := make([]int, 0)
		nexty := make([]int, 0)
		res += 1
		for i := 0; i < len(ax); i++ {
			for j := 0; j < 4; j++ {
				curx := ax[i] + px[j]
				cury := ay[i] + py[j]
				if curx >= 0 && curx < m && cury >= 0 && cury < n && grid[curx][cury] == 0 {
					nextx = append(nextx, curx)
					nexty = append(nexty, cury)
					grid[curx][cury] = 1
				}
			}
		}
		ax = nextx
		ay = nexty
	}
	return res
}

func main() {
	fmt.Println("a")
}
