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

//两种不同的dp思路
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]bool, m+1)
	}
	dp[0][0] = true
	for i := 1; i < n+1; i++ {
		if p[i-1] == '*' {
			dp[i][0] = dp[i-1][0]
		}
		for j := 1; j < m+1; j++ {
			if p[i-1] == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			} else if p[i-1] == '?' || p[i-1] == s[j-1] {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[n][m]
}

func isMatch2(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 1; i < n+1; i++ {
		if p[i-1] != '*' {
			break
		}
		dp[0][i] = true
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			} else if p[j-1] == '?' || p[j-1] == s[i-1] {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[m][n]
}

func main() {
	fmt.Println("a")
}
