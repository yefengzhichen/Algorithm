package main

import (
	"testing"
)

func TestA(t *testing.T) {
	tests := []struct {
		name string
		ina  []int
		inb  []int
		out  float64
	}{
		{
			name: "1",
			ina:  []int{1, 2, 8},
			inb:  []int{3, 4, 9},
			out:  3.5,
		},
		{
			name: "2",
			ina:  []int{2, 8},
			inb:  []int{3, 4, 9},
			out:  4.0,
		},
		{
			name: "3",
			ina:  []int{2, 3, 4},
			inb:  []int{1},
			out:  2.5,
		},
		{
			name: "3",
			ina:  []int{1, 2},
			inb:  []int{1, 2},
			out:  1.5,
		},
		{
			name: "3",
			ina:  []int{1},
			inb:  []int{2, 3, 4, 5, 6},
			out:  3.5,
		},
	}
	for i, test := range tests {
		if i < 3 {
			continue
		}
		t.Run(test.name, func(t *testing.T) {
			res := findMedianSortedArrays(test.ina, test.inb)
			if res != test.out {
				t.Errorf("res %f, want %f", res, test.out)
			}
		})
	}
}
