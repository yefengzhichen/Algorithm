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
			name: "1",
			in:   [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}},
			out:  4,
		},
		{
			name: "1",
			in:   [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}},
			out:  -1,
		},
		{
			name: "1",
			in:   [][]int{{0, 2}},
			out:  0,
		},
		{
			name: "1",
			in:   [][]int{{0}},
			out:  0,
		},
	}
	for i, test := range tests {
		if i < 0 {
			continue
		}
		t.Run(test.name, func(t *testing.T) {
			res := orangesRotting2(test.in)
			if res != test.out {
				t.Errorf("res %d, want %d", res, test.out)
			}
		})
	}
}
