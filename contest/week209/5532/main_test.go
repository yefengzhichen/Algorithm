package main

import (
	"testing"
)

func TestA(t *testing.T) {
	tests := []struct {
		name string
		in   *TreeNode
		out  bool
	}{
		{
			name: "1",
			in:   &TreeNode{5, &TreeNode{9, &TreeNode{3, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{1, &TreeNode{7, nil, nil}, nil}},
			out:  false,
		},
	}
	for i, test := range tests {
		if i < 0 {
			continue
		}
		t.Run(test.name, func(t *testing.T) {
			res := isEvenOddTree(test.in)
			if res != test.out {
				t.Errorf("res %v, want %v", res, test.out)
			}
		})
	}
}
