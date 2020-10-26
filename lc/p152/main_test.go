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
			in:   []int{2, 3, -2, 4},
			out:  6,
		},
		{
			name: "2",
			in:   []int{-2, 0, -1},
			out:  0,
		},
		{
			name: "3",
			in:   []int{-2},
			out:  -2,
		},
		{
			name: "4",
			in:   []int{0, 2},
			out:  2,
		},
	}
	for i, test := range tests {
		if i < 3 {
			continue
		}
		t.Run(test.name, func(t *testing.T) {
			res := maxProduct(test.in)
			if res != test.out {
				t.Errorf("res %d, want %d", res, test.out)
			}
		})
	}
}
