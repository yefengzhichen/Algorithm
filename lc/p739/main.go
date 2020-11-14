package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dailyTemperatures(T []int) []int {
	n := len(T)
	res := make([]int, n)
	if n < 1 {
		return res
	}
	stack := make([]int, n)
	stack[0] = 0
	last := 0
	for i := 1; i < n; i++ {
		j := last
		for ; j >= 0 && T[i] > T[stack[j]]; j-- {
			res[stack[j]] = i - stack[j]
		}
		stack[j+1] = i
		last = j + 1
	}
	return res
}

func dailyTemperatures2(T []int) []int {
	arr := make([]int, 0)
	n := len(T)
	if n < 1 {
		return arr
	}
	stack := make([]int, 0)
	stack = append(stack, n-1)
	arr = append(arr, 0)
	for i := n - 2; i >= 0; i-- {
		// top := stack[len(stack)-1]
		for len(stack) > 0 && T[i] >= T[stack[len(stack)-1]] {
			stack = stack[0 : len(stack)-1]
		}
		cur := 0
		if len(stack) > 0 {
			cur = stack[len(stack)-1] - i
		}
		stack = append(stack, i)
		arr = append(arr, cur)
	}
	res := make([]int, len(arr))
	for i := n - 1; i >= 0; i-- {
		res[n-1-i] = arr[i]
	}
	return res
}

func main() {
	fmt.Println("a")
}
