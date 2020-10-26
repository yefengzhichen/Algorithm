package main

import (
	"testing"
)

func TestA(t *testing.T) {
	tests := []struct {
		name string
		in   [][]int
		aim  int
		out  int
	}{
		{
			name: "1",
			in:   [][]int{{1, 2, 3}, {1, 1, 1}},
			out:  4,
		},
		{
			name: "2",
			in:   [][]int{{1, 2, 3}, {1, 2, 1}, {1, 1, 2}},
			out:  5,
		},
		{
			name: "2",
			in:   [][]int{{613, 0, 93, 463}, {101, 369, 112, 255}, {42, 67, 86, 543}, {485, 452, 393, 461}},
			out:  1758,
		},
	}
	for i, test := range tests {
		if i < 0 {
			continue
		}
		t.Run(test.name, func(t *testing.T) {
			res := getMin2(test.in)
			if res != test.out {
				t.Errorf("res %d, want %d", res, test.out)
			}
		})
	}
}
