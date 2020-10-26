package main

import (
	"testing"
)

func TestA(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		aim  int
		out  int
	}{
		{
			name: "1",
			in:   []int{5, 2, 3},
			aim:  20,
			out:  4,
		},
		{
			name: "1",
			in:   []int{5, 2, 3},
			aim:  0,
			out:  0,
		},
		{
			name: "1",
			in:   []int{5, 3},
			aim:  2,
			out:  -1,
		},
	}
	for i, test := range tests {
		if i < 0 {
			continue
		}
		t.Run(test.name, func(t *testing.T) {
			res := minCoins(test.in, test.aim)
			if res != test.out {
				t.Errorf("res %d, want %d", res, test.out)
			}
		})
	}
}
