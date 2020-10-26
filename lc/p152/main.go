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

func maxProduct(nums []int) int {
	last, res := 0, nums[0]
	nums = append(nums, 0)
	sort.Ints()
	for i, val := range nums {
		if val == 0 && i > last {
			if i != len(nums)-1 {
				res = max(res, 0)
			}
			n := i - last
			left := subProduct(nums[last:i])
			total := left[n-1]
			fmt.Println(left, n)
			if left[n-1] > 0 && n >= 1 {
				res = max(res, total)
			} else {
				for j := last; j < i; j++ {
					if nums[j] < 0 {
						if i-last == 1 {
							res = max(res, nums[j])
						} else {
							if j == last {
								res = max(res, total/nums[j])
							} else if j == i-1 {
								res = max(res, left[j-1])
							} else {
								cur := max(left[j-last-1], total/left[j-last-1]/nums[j])
								res = max(res, cur)
							}
						}
					}
				}
			}
			i += 1
			last = i
		} else {
			i += 1
		}
	}
	return res
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func subProduct(nums []int) (left []int) {
	for i, cur := 0, 1; i < len(nums); i++ {
		cur *= nums[i]
		left = append(left, cur)
	}
	return left
}

func main() {
	fmt.Println("a")
}
