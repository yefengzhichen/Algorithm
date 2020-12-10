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

func isMatch(s string, p string) bool {
	m, n := len(p), len(s)
	dp := make([][]bool, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 2; i < m+1; i += 2 {
		if p[i-1] == '*' {
			dp[i][0] = true
		}
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if p[i-1] == '*' && i >= 2 {
				dp[i][j] = dp[i][j-1] && (p[i-2] == '.' || p[i-2] == s[j-1])
			} else if p[i-1] == '.' || p[i-1] == s[j-1] {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[m][n]
}

func main() {
	fmt.Println("a")
}
