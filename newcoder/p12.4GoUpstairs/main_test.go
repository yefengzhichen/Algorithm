package main

import (
	"testing"
)

func TestA(t *testing.T) {
	tests := []struct {
		name string
		in   int
		out  int
	}{
		{
			name: "1",
			in:   1,
			out:  1,
		},
		{
			name: "2",
			in:   3,
			out:  3,
		},
		{
			name: "3",
			in:   4,
			out:  5,
		},
	}
	for i, test := range tests {
		if i < 0 {
			continue
		}
		t.Run(test.name, func(t *testing.T) {
			res := countWays(test.in)
			if res != test.out {
				t.Errorf("res %d, want %d", res, test.out)
			}
		})
	}
}
