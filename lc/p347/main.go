package main

import "fmt"

func initHeap(arr []int, n int) {
	for i := n/2 - 1; i >= 0; i-- {
		adjustHeap(arr, i, n)
	}
}

func adjustHeap(arr []int, start int, end int) {
	for i := start; 2*i+1 < end; {
		child := 2*i + 1
		if child+1 < end && arr[child] > arr[child+1] {
			child = child + 1
		}
		if arr[i] <= arr[child] {
			break
		}
		tmp := arr[i]
		arr[i] = arr[child]
		arr[child] = tmp
		i = child
	}
}

func topKFrequent(nums []int, k int) []int {
	var res []int
	if len(nums) < 1 {
		return res
	}
	m := make(map[int]int, 0)
	for _, num := range nums {
		m[num] += 1
	}
	var valSet []int
	for _, val := range m {
		valSet = append(valSet, val)
	}
	initHeap(valSet, k)
	for i := k; i < len(valSet); i++ {
		if valSet[0] < valSet[i] {
			tmp := valSet[0]
			valSet[0] = valSet[i]
			valSet[i] = tmp
			adjustHeap(valSet, 0, k)
		}
	}
	for key, val := range m {
		if val >= valSet[0] {
			res = append(res, key)
		}
	}
	return res
}

func main() {
	fmt.Println("a")
}
