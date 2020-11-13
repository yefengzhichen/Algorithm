package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func firstGreatOrEqual(arr []int, val int) int {
	low, high := 0, len(arr)-1
	mid := 0
	for low <= high {
		mid = low + ((high - low) >> 1)
		if arr[mid] >= val {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	if low < len(arr) {
		return low
	}
	return -1
}

func firstGreat(arr []int, val int) int {
	low, high := 0, len(arr)-1
	mid := 0
	for low <= high {
		mid = low + ((high - low) >> 1)
		if arr[mid] > val {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	if low < len(arr) {
		return low
	}
	return -1
}

//二分查找等于某个值
func binarySearch(arr []int, val int) int {
	low, high := 0, len(arr)-1
	mid := 0
	for low <= high {
		mid = low + ((high - low) >> 1)
		if arr[mid] == val {
			return mid
		} else if arr[mid] > val {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func main() {
	mid := 2 + (8-2)>>1
	fmt.Println(mid)
}
