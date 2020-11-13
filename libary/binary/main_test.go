package main

import (
	"testing"
)

func TestFirstGreatOrEqual(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		val  int
		out  int
	}{
		{
			name: "1",
			in:   []int{1, 2, 3, 3, 3, 5},
			val:  3,
			out:  2,
		},
		{
			name: "2",
			in:   []int{1, 2, 3, 3, 3, 5},
			val:  4,
			out:  5,
		},
		{
			name: "3",
			in:   []int{1, 2, 2, 2, 4, 8, 10},
			val:  2,
			out:  1,
		},
	}
	for i, test := range tests {
		if i < 0 {
			continue
		}
		t.Run(test.name, func(t *testing.T) {
			res := firstGreatOrEqual(test.in, test.val)
			if res != test.out {
				t.Errorf("res %d, want %d", res, test.out)
			}
		})
	}
}

func TestFirstGreat(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		val  int
		out  int
	}{
		{
			name: "1",
			in:   []int{1, 2, 3, 3, 3, 5},
			val:  3,
			out:  5,
		},
		{
			name: "2",
			in:   []int{1, 2, 3, 3, 3, 5},
			val:  2,
			out:  2,
		},
		{
			name: "3",
			in:   []int{1, 2, 2, 2, 4, 8, 10},
			val:  2,
			out:  4,
		},
	}
	for i, test := range tests {
		if i < 0 {
			continue
		}
		t.Run(test.name, func(t *testing.T) {
			res := firstGreat(test.in, test.val)
			if res != test.out {
				t.Errorf("res %d, want %d", res, test.out)
			}
		})
	}
}
