package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isEvenOddTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	odd := 0
	start := []int{-1, 2000000}
	arr := []*TreeNode{root}
	for len(arr) > 0 {
		var next []*TreeNode
		last := start[odd]
		for _, node := range arr {
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
			if odd == 0 && (node.Val%2 == 0 || node.Val <= last) {
				return false
			}
			if odd == 1 && (node.Val%2 == 1 || node.Val >= last) {
				return false
			}
			last = node.Val
		}
		odd = 1 - odd
		arr = next
	}
	return true
}

func main() {
	fmt.Println("a")
}
