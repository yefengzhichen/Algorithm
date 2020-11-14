package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	Val int
	X   int
	Y   int
}

func kthSmallest(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	heap := make([]Node, m)
	for i := 0; i < m; i++ {
		heap[i] = Node{matrix[i][0], i, 0}
	}
	heapInit(heap)
	arr := make([]int, m*n)
	index := 0
	for len(heap) > 0 {
		cur := heap[0]
		arr[index] = cur.Val
		index += 1
		if cur.Y+1 < n {
			heap[0] = Node{matrix[cur.X][cur.Y+1], cur.X, cur.Y + 1}
			heapAdjust(heap, 0, len(heap))
		} else {
			heap[0] = heap[len(heap)-1]
			heap = heap[0 : len(heap)-1]
			heapAdjust(heap, 0, len(heap))
		}
	}
	// fmt.Println(arr)
	return arr[k-1]
}

func heapAdjust(arr []Node, start, end int) {
	for cur := start; 2*cur+1 < end; {
		child := 2*cur + 1
		if child+1 < end && arr[child+1].Val < arr[child].Val {
			child += 1
		}
		if arr[child].Val > arr[cur].Val {
			break
		}
		arr[cur], arr[child] = arr[child], arr[cur]
		cur = child
	}
}

func heapInit(arr []Node) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		heapAdjust(arr, i, len(arr))
	}
}

func main() {
	fmt.Println("a")
}
