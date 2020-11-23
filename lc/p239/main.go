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

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	stack := make([]int, 0)
	if n < k {
		return stack
	}
	res := make([]int, 0)
	for i := 0; i < n; i++ {
		for len(stack) > 0 && i-stack[0] >= k {
			stack = stack[1:len(stack)]
		}
		for len(stack) > 0 && nums[i] >= nums[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
		if i >= k-1 {
			fmt.Println(nums)
			res = append(res, nums[stack[0]])
		}
	}
	return res
}

func main() {
	fmt.Println("a")
}
