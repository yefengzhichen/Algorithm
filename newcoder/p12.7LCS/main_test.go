package main

import (
	"testing"
)

func TestA(t *testing.T) {
	tests := []struct {
		name string
		ina  string
		inb  string
		out  int
	}{
		{
			name: "1",
			ina:  "1A2C3D46",
			inb:  "B1D23CDB",
			out:  4,
		},
		{
			name: "2",
			ina:  "2222",
			inb:  "2222",
			out:  4,
		},
		{
			name: "3",
			ina:  "abcde",
			inb:  "ace",
			out:  3,
		},
		{
			name: "4",
			ina:  "abc",
			inb:  "abc",
			out:  3,
		},
		{
			name: "5",
			ina:  "abc",
			inb:  "def",
			out:  0,
		},
		{
			name: "6",
			ina:  "abcde",
			inb:  "acer",
			out:  3,
		},
	}
	for i, test := range tests {
		if i < 0 {
			continue
		}
		t.Run(test.name, func(t *testing.T) {
			res := longestCommonSubsequence(test.ina, test.inb)
			if res != test.out {
				t.Errorf("res %d, want %d", res, test.out)
			}
		})
	}
}
