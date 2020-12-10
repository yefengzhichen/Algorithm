package main

import (
	"testing"
)

func TestA(t *testing.T) {
	tests := []struct {
		name string
		ins  string
		inp  string
		out  bool
	}{
		{
			name: "1",
			ins:  "",
			inp:  "a*",
			out:  true,
		},
		{
			name: "2",
			ins:  "",
			inp:  "",
			out:  true,
		},
		{
			name: "3",
			ins:  "bab",
			inp:  ".*",
			out:  true,
		},
		{
			name: "4",
			ins:  "aaab",
			inp:  "a*",
			out:  false,
		},
	}
	for i, test := range tests {
		if i < 0 {
			continue
		}
		t.Run(test.name, func(t *testing.T) {
			res := isMatch(test.ins, test.inp)
			if res != test.out {
				t.Errorf("name %s, res %v, want %v", test.name, res, test.out)
			}
		})
	}
}
