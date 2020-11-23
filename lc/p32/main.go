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

func longestValidParentheses(s string) int {
	n := len(s)
	dp := make([]int, len(s))
	max := 0
	for i := 1; i < n; i++ {
		if s[i] == '(' {
			continue
		}
		tmp := dp[i-1]
		if s[i-tmp-1] == '(' {
			dp[i] = tmp + 2
			if i-tmp-2 >= 0 {
				dp[i] += dp[i-tmp-2]
			}
			if dp[i] > max {
				max = dp[i]
			}
		}
	}
	return max
}

func longestValidParentheses2(s string) int {
	n := len(s)
	res := make([]int, len(s))
	stack := make([]int, 0)
	max := 0
	for i := 0; i < n; i++ {
		if s[i] == '(' || len(stack) == 0 {
			stack = append(stack, i)
		} else {
			top := stack[len(stack)-1]
			if s[top] == '(' {
				stack = stack[:len(stack)-1]
				res[i] = res[top] + (i - top + 1)
				if top > 0 {
					res[i] += res[top-1]
				}
				if res[i] > max {
					max = res[i]
				}
			} else {
				stack = append(stack, i)
			}
		}
	}
	return max
}

func main() {
	fmt.Println("a")
}
