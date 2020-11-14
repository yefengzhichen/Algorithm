package main

import (
	"testing"
)

func TestA(t *testing.T) {
	tests := []struct {
		name string
		in   [][]int
		k    int
		out  int
	}{
		{
			name: "1",
			in:   [][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}},
			k:    8,
			out:  13,
		},
	}
	for i, test := range tests {
		if i < 0 {
			continue
		}
		t.Run(test.name, func(t *testing.T) {
			res := kthSmallest(test.in, test.k)
			if res != test.out {
				t.Errorf("res %d, want %d", res, test.out)
			}
		})
	}
}
