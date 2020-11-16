package main

import (
	"testing"
)

func TestA(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		k    int
		out  int
	}{
		{
			name: "1",
			in:   []int{3, 2, 20, 1, 1, 3},
			k:    10,
			out:  5,
		},
	}
	for i, test := range tests {
		if i < 0 {
			continue
		}
		t.Run(test.name, func(t *testing.T) {
			res := minOperations(test.in, test.k)
			if res != test.out {
				t.Errorf("res %d, want %d", res, test.out)
			}
		})
	}
}
