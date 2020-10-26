package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func canPartition(nums []int) bool {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	if sum%2 == 1 {
		return false
	}
	sum = sum / 2
	return process(nums, 0, len(nums)-1, sum)
}

func process(arr []int, low, high, val int) bool {
	if val < 0 {
		return false
	} else if val == 0 {
		return true
	}
	if low == high {
		if arr[low] == val {
			return true
		} else {
			return false
		}
	}
	return process(arr, low+1, high, val-arr[low]) || process(arr, low+1, high, val)
}

func main() {
	fmt.Println("a")
}
