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

func reversePairs(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}
	cnt, _ := process(nums, 0, n)
	return cnt
}

func process(nums []int, low, high int) (int, []int) {
	if low == high-1 {
		return 0, []int{nums[low]}
	}
	mid := low + (high-low)>>1
	lcnt, larr := process(nums, low, mid)
	rcnt, rarr := process(nums, mid, high)
	arr := make([]int, 0)
	i, j, res := 0, 0, 0
	for i < len(larr) && j < len(rarr) {
		if larr[i] <= rarr[j] {
			arr = append(arr, larr[i])
			i += 1
			res += j
		} else {
			arr = append(arr, rarr[j])
			j += 1
		}
	}
	// fmt.Println(arr, i, j, lcnt, rcnt, res)
	for i < len(larr) {
		arr = append(arr, larr[i])
		i += 1
		res += j
	}
	for j < len(rarr) {
		arr = append(arr, rarr[j])
		j += 1
	}
	return res + lcnt + rcnt, arr
}

func reversePairs2(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}
	arr := make([]int, n)
	arr[0] = nums[0]
	res := 0
	for i := 1; i < n; i++ {
		index := binarySearch(arr, i, nums[i])
		res += index
		for j := i; j > index; j-- {
			arr[j] = arr[j-1]
		}
		arr[index] = nums[i]
	}
	return res
}

func binarySearch(arr []int, n int, val int) int {
	low, high := 0, n-1
	for low <= high {
		mid := low + (high-low)>>1
		if arr[mid] > val {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if low < 0 {
		return 0
	}
	return low
}

func main() {
	fmt.Println("a")
}
