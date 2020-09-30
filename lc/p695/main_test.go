package main

import (
	"testing"
)

func TestA(t *testing.T) {
	tests := []struct {
		name string
		in   [][]int
		out  int
	}{
		{
			name: "nil",
			in:   [][]int{},
			out:  0,
		},
		{
			name: "row",
			in:   [][]int{{0, 0, 1, 1, 0}},
			out:  2,
		},
		{
			name: "col",
			in:   [][]int{{0}, {1}, {1}, {0}},
			out:  2,
		},
		{
			name: "normal",
			in: [][]int{
				{0, 0, 1, 1, 0, 0},
				{0, 0, 1, 0, 1, 0},
				{0, 0, 1, 0, 0, 0},
				{0, 0, 0, 1, 1, 0},
				{0, 0, 0, 1, 0, 0},
			},
			out: 4,
		},
	}
	for i, test := range tests {
		if i < 0 {
			continue
		}
		t.Run(test.name, func(t *testing.T) {
			res := maxAreaOfIsland(test.in)
			if res != test.out {
				t.Errorf("res %d, want %d", res, test.out)
			}
		})
	}
}
