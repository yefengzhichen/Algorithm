package main

import (
	"testing"
)

func TestA(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		ink  int
		out  int
	}{
		{
			name: "1",
			in:   []int{3, 2, 3, 1, 2, 4, 5, 5, 6, 7, 7, 8, 2, 3, 1, 1, 1, 10, 11, 5, 6, 2, 4, 7, 8, 5, 6},
			ink:  20,
			out:  2,
		},
		{
			name: "1",
			in:   []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			ink:  4,
			out:  4,
		},
	}
	for i, test := range tests {
		if i < 0 {
			continue
		}
		t.Run(test.name, func(t *testing.T) {
			res := findKthLargest(test.in, test.ink)
			if res != test.out {
				t.Errorf("res %d, want %d", res, test.out)
			}
		})
	}
}
