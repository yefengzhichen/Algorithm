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

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	root := &ListNode{0, head}
	node, num := head, k
	for ; node.Next != nil && num > 1; num-- {
		node = node.Next
	}
	if num != 1 {
		return head
	}
	next := node.Next
	node.Next = nil
	for node = head; node.Next != nil; {
		tmp := node.Next
		node.Next = tmp.Next
		tmp.Next = root.Next
		root.Next = tmp
	}
	node.Next = reverseKGroup(next, k)
	return root.Next
}

func reverseKGroup2(head *ListNode, k int) *ListNode {
	count := 0
	for cur := head; cur != nil; cur = cur.Next {
		count += 1
	}
	num := count / k
	tmp := &ListNode{0, head}
	root := tmp
	next, cur := head, head
	for i := 0; i < num; i++ {
		for j := 1; j < k; j++ {
			cur = cur.Next
		}
		next = cur.Next
		cur.Next = nil
		root = reverse(root)
		root.Next = next
		cur = next
	}
	return tmp.Next
}

func reverse(root *ListNode) *ListNode {
	cur := root.Next
	for cur.Next != nil {
		next := cur.Next
		cur.Next = next.Next
		next.Next = root.Next
		root.Next = next
	}
	return cur
}

func main() {
	fmt.Println("a")
}
