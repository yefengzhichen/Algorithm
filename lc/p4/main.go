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

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if (m+n)%2 == 0 {
		left := getKth(nums1, 0, m, nums2, 0, n, (m+n+1)/2)
		right := getKth(nums1, 0, m, nums2, 0, n, (m+n+2)/2)
		fmt.Println(left, right)
		return float64(left+right) / 2.0
	}
	tmp := getKth(nums1, 0, m, nums2, 0, n, (m+n+1)/2)
	return float64(tmp)
}

func getKth(a []int, i int, m int, b []int, j int, n int, k int) int {
	fmt.Println(i, m, j, n, k)
	if m == i {
		return b[j+k-1]
	}
	if n == j {
		return a[i+k-1]
	}
	if k == 1 {
		return min(a[i], b[j])
	}
	half := k / 2
	pa := min(i+half, m) - 1
	pb := min(j+half, n) - 1
	// fmt.Println(a, b)
	// fmt.Println(pa, a[i+pa-1], pb, b[j+pb-1])
	if a[pa] <= b[pb] {
		return getKth(a, pa+1, m, b, j, n, k-(pa-i+1))
	}
	return getKth(a, i, m, b, pb+1, n, k-(pb-j+1))
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println("a")
}
