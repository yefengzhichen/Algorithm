package main

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func specialArray(nums []int) int {
	sort.Ints(nums)
	n, i := len(nums), 0
	last := 0
	for ; i < n && nums[i] < 0; i++ {
	}
	if i == n {
		return 0
	}
	for i < n {
		want := n - i
		if want > last && want <= nums[i] {
			return want
		}
		last = nums[i]
		for i += 1; i < n && nums[i] == last; i++ {
		}
	}
	return -1
}

func main() {
	fmt.Println("a")
}
