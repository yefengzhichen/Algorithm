package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	root := &ListNode{0, head}
	node := head
	for node.Next != nil {
		next := node.Next
		node.Next = next.Next
		next.Next = root.Next
		root.Next = next
	}
	return root.Next
}

func main() {
	fmt.Println("a")
}
