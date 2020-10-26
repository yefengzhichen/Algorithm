package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func coins(arr []int, aim int) int {
	m := len(arr)
	var dp [1001][1001]int
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for j := 0; j <= aim; j++ {
		if j%arr[0] == 0 {
			dp[0][j] = 1
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j <= aim; j++ {
			for k := 0; j-k*arr[i] >= 0; k++ {
				dp[i][j] += dp[i-1][j-k*arr[i]]
			}
		}
		fmt.Println(dp[i][0 : aim+1])
	}
	return dp[m-1][aim]
}

func coins2(arr []int, aim int) int {
	m := len(arr)
	var dp [1001][1001]int
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for j := 0; j <= aim; j++ {
		if j%arr[0] == 0 {
			dp[0][j] = 1
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j <= aim; j++ {
			dp[i][j] = dp[i-1][j]
			if j-arr[i] >= 0 {
				dp[i][j] += dp[i][j-arr[i]]
			}
		}
		fmt.Println(dp[i][0 : aim+1])
	}
	return dp[m-1][aim]
}

func coins3(arr []int, aim int) int {
	m := len(arr)
	var dp [1001]int
	dp[0] = 1
	for i := 0; i < m; i++ {
		for j := arr[i]; j <= aim; j++ {
			dp[j] += dp[j-arr[i]]
		}
	}
	return dp[aim]
}

func main() {
	fmt.Println("a")
}
