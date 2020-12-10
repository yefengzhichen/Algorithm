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

func translateNum(num int) int {
	x, y := 1, 1
	for num > 0 {
		tmp := num % 100
		if tmp >= 10 && tmp <= 25 {
			y, x = x+y, y
		} else {
			y, x = y, y
		}
		num = num / 10
	}
	return y
}

func main() {
	fmt.Println("a")
}
