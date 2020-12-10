package main

import (
	"testing"
)

func TestA(t *testing.T) {
	tests := []struct {
		name string
		inS  string
		inT  string
		out  string
	}{
		{
			name: "1",
			inS:  "ABANC",
			inT:  "ABC",
			out:  "BANC",
		},
		{
			name: "1",
			inS:  "BANCBA",
			inT:  "ABC",
			out:  "CBA",
		},
		{
			name: "1",
			inS:  "ABC",
			inT:  "ABC",
			out:  "ABC",
		},
		{
			name: "1",
			inS:  "A",
			inT:  "AA",
			out:  "",
		},
	}
	for i, test := range tests {
		if i < 0 {
			continue
		}
		t.Run(test.name, func(t *testing.T) {
			res := minWindow(test.inS, test.inT)
			if res != test.out {
				t.Errorf("res %s, want %s", res, test.out)
			}
		})
	}
}
