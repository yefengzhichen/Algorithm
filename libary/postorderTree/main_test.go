package main

import (
	"reflect"
	"testing"
)

func TestA(t *testing.T) {
	tests := []struct {
		name string
		in   *TreeNode
		out  []int
	}{
		{
			name: "1",
			in:   &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{3, &TreeNode{6, nil, nil}, &TreeNode{7, nil, nil}}},
			out:  []int{4, 5, 2, 6, 7, 3, 1},
		},
	}
	for i, test := range tests {
		if i < 0 {
			continue
		}
		t.Run(test.name, func(t *testing.T) {
			res := postorderTraversal2(test.in)
			if reflect.DeepEqual(test.out, res) {
				t.Errorf("res %v, want %v", res, test.out)
			}
		})
	}
}
