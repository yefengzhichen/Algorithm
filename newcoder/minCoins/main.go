package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minCoins(coins []int, aim int) int {
	arr := make([]int, aim+1)
	for i := 1; i <= aim; i++ {
		cur := aim + 1
		for j := 0; j < len(coins); j++ {
			index := i - coins[j]
			if index >= 0 && arr[index] >= 0 {
				cur = min(cur, arr[index]+1)
			}
		}

		if cur == aim+1 {
			arr[i] = -1
		} else {
			arr[i] = cur
		}
	}
	return arr[aim]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func minCoins2(coins []int, aim int) int {
	arr := make([]int, aim+1)
	for i := 0; i < aim+1; i++ {
		arr[i] = aim + 1
	}
	for i := 0; i < len(coins); i++ {
		for j := 1; j < aim; j++ {
			if j-arr[i] >= 0 {
				arr[i] = min(arr[i], arr[j-arr[i]])
			}
		}
	}
	if arr[aim] == aim+1 {
		return -1
	}
	return arr[aim]
}

func main() {
	n, aim := 0, 0
	fmt.Scanln(&n, &aim)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	//fmt.Println(n, aim, arr)
	fmt.Println(minCoins(arr, aim))
}
