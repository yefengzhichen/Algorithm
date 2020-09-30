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
			name: "nil",
			in:   nil,
			out:  []int{},
		},
		{
			name: "one elem",
			in:   &TreeNode{1, nil, nil},
			out:  []int{1},
		},
		{
			name: "three elem",
			in:   &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}},
			out:  []int{1, 2, 3},
		},
		{
			name: "complete",
			in:   &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{3, &TreeNode{6, nil, nil}, &TreeNode{7, nil, nil}}},
			out:  []int{1, 2, 4, 5, 3, 6, 7},
		},
		{
			name: "right",
			in:   &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, &TreeNode{4, nil, nil}, &TreeNode{5, nil, &TreeNode{6, nil, nil}}}},
			out:  []int{1, 2, 3, 4, 5, 6},
		},
	}
	for i, test := range tests {
		if i < 0 {
			continue
		}
		t.Run(test.name, func(t *testing.T) {
			res := preorderTraversal(test.in)
			if !reflect.DeepEqual(res, test.out) {
				t.Errorf("res %v %d, want %v %d", res, len(res), test.out, len(test.out))
			}
		})
	}
}
