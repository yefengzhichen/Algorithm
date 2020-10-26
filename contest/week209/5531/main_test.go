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
			in:   []int{3, 5},
			out:  2,
		},
		{
			name: "2",
			in:   []int{0, 4, 3, 0, 4},
			out:  3,
		},
		{
			name: "3",
			in:   []int{0, 0},
			out:  -1,
		},
		{
			name: "4",
			in:   []int{3, 6, 7, 7, 0},
			out:  -1,
		},
	}
	for i, test := range tests {
		if i < 0 {
			continue
		}
		t.Run(test.name, func(t *testing.T) {
			res := specialArray(test.in)
			if res != test.out {
				t.Errorf("res %d, want %d", res, test.out)
			}
		})
	}
}
