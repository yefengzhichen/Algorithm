package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minOperations2(nums []int, x int) int {
	n := len(nums)
	left := make([]int, n+1)
	left[0] = 0
	for i := 1; i < n; i++ {
		left[i] = left[i-1] + nums[i-1]
	}
	right := make([]int, n+1)
	right[0] = 0
	for i := 1; i < n; i++ {
		right[i] = right[i-1] + nums[n-i]
	}
	//fmt.Println(left, right)
	min := int(math.Pow(10, 9)) + 1
	for i := 0; i < n; i++ {
		for j := 0; j < n-i+1; j++ {
			cur := left[i] + right[j]
			//fmt.Println(i, j, cur, min)
			if cur == x && (i+j) < min {
				min = i + j
			} else if cur > x {
				break
			}
		}
	}
	if min == int(math.Pow(10, 9))+1 {
		return -1
	}
	return min
}

func minOperations(nums []int, x int) int {
	n := len(nums)
	left := make([]int, n+1)
	left[0] = 0
	for i := 1; i < n; i++ {
		left[i] = left[i-1] + nums[i-1]
	}
	right := make([]int, n+1)
	right[0] = 0
	for i := 1; i < n; i++ {
		right[i] = right[i-1] + nums[n-i]
	}
	//fmt.Println(left, right)
	min := int(math.Pow(10, 9)) + 1
	i, j := n, 0
	for i >= 0 {
		cur := left[i] + right[j]
		if cur == x {
			if i+j < min {
				min = i + j
			}
			i -= 1
			j += 1
		} else if cur < x && i+j < n {
			j += 1
		} else {
			i -= 1
		}
	}
	if min == int(math.Pow(10, 9))+1 {
		return -1
	}
	return min
}

func main() {
	fmt.Println("a")
}
