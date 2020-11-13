package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		lastMax := 0
		for j := 0; j < n; j++ {
			if i-1 >= 0 && j-1 >= 0 {
				lastMax = max(lastMax, dp[i-1][j-1])
			}
			if text1[i] == text2[j] {
				dp[i][j] = lastMax + 1
			} else {
				if i-1 >= 0 {
					lastMax = max(lastMax, dp[i-1][j])
				}
				dp[i][j] = lastMax
			}
		}
		fmt.Println(lastMax, dp[i])
	}
	return dp[m-1][n-1]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	fmt.Println("a")
}
