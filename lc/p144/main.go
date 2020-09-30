package main

import (
	"fmt"
	"reflect"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		if cur == nil {
			continue
		}
		res = append(res, cur.Val)
		stack = append(stack, cur.Right)
		stack = append(stack, cur.Left)
	}
	return res
}
func main() {
	fmt.Println("a")
	var res []int
	b := []int{}
	fmt.Println(res == nil)
	fmt.Println(b == nil)
	fmt.Println(reflect.DeepEqual(res, b))
}
