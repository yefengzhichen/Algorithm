package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type OrderedStream struct {
	arr    []string
	ptr    int
	length int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{arr: make([]string, n), ptr: 0, length: n}
}

func (this *OrderedStream) Insert(id int, value string) []string {
	this.arr[id-1] = value
	res := make([]string, 0)
	if id-1 == this.ptr {
		id -= 1
		for ; id < this.length && len(this.arr[id]) > 0; id++ {
			res = append(res, this.arr[id])
		}
		this.ptr = id
	}
	return res
}

func main() {
	fmt.Println("a")
}
