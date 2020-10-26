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
			in:   []int{1, 2, 4},
			aim:  3,
			out:  2,
		},
		{
			name: "2",
			in:   []int{1, 2, 4},
			aim:  4,
			out:  4,
		},
		{
			name: "3",
			in:   []int{1, 3, 6},
			aim:  13,
			out:  9,
		},
	}
	for i, test := range tests {
		if i < 0 {
			continue
		}
		t.Run(test.name, func(t *testing.T) {
			res := coins3(test.in, test.aim)
			if res != test.out {
				t.Errorf("res %d, want %d", res, test.out)
			}
		})
	}
}
