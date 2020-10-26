package main

import (
	"testing"
)

func TestA(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		out  bool
	}{
		{
			name: "1",
			in:   []int{},
			out:  true,
		},
		{
			name: "2",
			in:   []int{1, 3, 2, 6, 5},
			out:  true,
		},
		{
			name: "3",
			in:   []int{1, 6, 3, 2, 5},
			out:  false,
		},
		{
			name: "3",
			in:   []int{1, 2, 5, 10, 6, 9, 4, 3},
			out:  false,
		},
	}
	for i, test := range tests {
		if i < 0 {
			continue
		}
		t.Run(test.name, func(t *testing.T) {
			res := verifyPostorder(test.in)
			if res != test.out {
				t.Errorf("res %v, want %v", res, test.out)
			}
		})
	}
}
