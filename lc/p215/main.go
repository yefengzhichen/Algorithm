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

func findKthLargest(nums []int, k int) int {
	arr := make([]int, k)
	for i := 0; i < k; i++ {
		arr[i] = nums[i]
	}
	for i := (k - 1) / 2; i >= 0; i-- {
		adjustHeap(arr, i, k)
	}
	for i := k; i < len(nums); i++ {
		if nums[i] <= arr[0] {
			continue
		}
		arr[0] = nums[i]
		adjustHeap(arr, 0, k)
	}
	return arr[0]
}

func adjustHeap(arr []int, index, n int) {
	child := 2*index + 1
	for child < n {
		if child+1 < n && arr[child+1] < arr[child] {
			child += 1
		}
		if arr[child] >= arr[index] {
			break
		}
		arr[child], arr[index] = arr[index], arr[child]
		index = child
		child = 2*index + 1
	}
}

func main() {
	fmt.Println("a")
}
