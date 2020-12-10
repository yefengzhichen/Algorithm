package main

import (
	"testing"
)

func TestA(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		out  int
	}{
		{
			name: "1",
			in:   []int{4, 5, 6, 7},
			out:  0,
		},
		{
			name: "1",
			in:   []int{7, 5, 6, 4},
			out:  5,
		},
		{
			name: "1",
			in:   []int{1, 3, 2, 3, 1},
			out:  4,
		},
	}
	for i, test := range tests {
		if i < 2 {
			continue
		}
		t.Run(test.name, func(t *testing.T) {
			res := reversePairs(test.in)
			if res != test.out {
				t.Errorf("res %d, want %d", res, test.out)
			}
		})
	}
}
