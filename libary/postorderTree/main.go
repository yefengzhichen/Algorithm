package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	prev := root
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Right == nil || root.Right == prev {
			res = append(res, root.Val)
			prev = root
			root = nil
		} else {
			stack = append(stack, root)
			root = root.Right
		}
	}
	return res
}

func postorderTraversal2(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		if cur == nil {
			continue
		}
		res = append(res, cur.Val)
		stack = append(stack, cur.Left)
		stack = append(stack, cur.Right)
	}
	n := len(res)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = res[n-i-1]
	}
	return arr
}

func main() {
	fmt.Println("a")
}
