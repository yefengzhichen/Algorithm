package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	dp := make([]int, len(nums)+1)
	dp[0] = 1
	res := 1
	for i := 1; i < len(nums); i++ {
		max := 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && max < dp[j]+1 {
				max = dp[j] + 1
			}
		}
		dp[i] = max
		if max > res {
			res = max
		}
	}
	return res
}
func lengthOfLIS2(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	var stack []int
	stack = append(stack, nums[0])
	for i := 1; i < len(nums); i++ {
		if nums[i] > stack[len(stack)-1] {
			stack = append(stack, nums[i])
		} else {
			index := firstGreatOrEqual(stack, nums[i])
			stack[index] = nums[i]
		}
	}
	return len(stack)
}

func firstGreatOrEqual(nums []int, val int) int {
	low, high := 0, len(nums)-1
	mid := 0
	for low <= high {
		mid = low + ((high - low) >> 1)
		if nums[mid] >= val {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	if low < len(nums) {
		return low
	}
	return -1
}

func main() {
	mid := 2 + (8-2)>>1
	fmt.Println(mid)
}
