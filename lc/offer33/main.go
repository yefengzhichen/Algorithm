package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func verifyPostorder(postorder []int) bool {
	return process(postorder, 0, len(postorder)-1)
}
func process(arr []int, low, high int) bool {
	if low >= high-1 {
		return true
	}
	i := high - 1
	for ; i >= low && arr[i] > arr[high]; i-- {
	}
	mid := i
	for ; i >= low; i-- {
		if arr[i] > arr[high] {
			return false
		}
	}
	return process(arr, low, mid) && process(arr, mid+1, high-1)
}

func main() {
	fmt.Println("a")
}
