package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMin(arr [][]int) int {
	row, col := len(arr), len(arr[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			tmp := 0
			if i-1 >= 0 {
				tmp = arr[i-1][j]
			}
			if j-1 > 0 && tmp > arr[i][j-1] {
				tmp = arr[i][j-1]
			}
			arr[i][j] += tmp
		}
	}
	return arr[row-1][col-1]
}

func getMin2(arr [][]int) int {
	row, col := len(arr), len(arr[0])
	var res [101]int
	for j := 0; j < col; j++ {
		res[j] = arr[0][j]
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			tmp := 0
			if i-1 >= 0 {
				tmp = res[j]
			}
			if j-1 > 0 && tmp > res[j-1] {
				tmp = res[j-1]
			}
			res[j] += tmp
		}
	}
	return res[col-1]
}

func main() {
	fmt.Println("a")
}
