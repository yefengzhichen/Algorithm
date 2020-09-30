package main

import (
	"testing"
)

func TestA(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "pure letters",
			in:   "abc",
			out:  "abc",
		},
		{
			name: "normal",
			in:   "3[a]2[bc]",
			out:  "aaabcbc",
		},
		{
			name: "nested",
			in:   "2[a3[bc]]",
			out:  "abcbcbcabcbcbc",
		},
		{
			name: "big int",
			in:   "10[ab]",
			out:  "abababababababababab",
		},
		{
			name: "mixed",
			in:   "ab3[z]",
			out:  "abzzz",
		},
		{
			name: "upper case",
			in:   "B2[F]",
			out:  "BFF",
		},
	}
	for i, test := range tests {
		if i < 0 {
			continue
		}
		t.Run(test.name, func(t *testing.T) {
			res := decodeString(test.in)
			if res != test.out {
				t.Errorf("res %s, want %s", res, test.out)
			}
		})
	}
}
