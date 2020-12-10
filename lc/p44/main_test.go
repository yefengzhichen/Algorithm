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
			ins:  "aaaab",
			inp:  "aa*",
			out:  true,
		},
		{
			name: "1",
			ins:  "bab",
			inp:  "aa*",
			out:  false,
		},
		{
			name: "1",
			ins:  "bab",
			inp:  "b?b",
			out:  true,
		},
		{
			name: "1",
			ins:  "bab",
			inp:  "",
			out:  false,
		},
		{
			name: "1",
			ins:  "",
			inp:  "",
			out:  true,
		},
		{
			name: "1",
			ins:  "",
			inp:  "***",
			out:  true,
		},
	}
	for i, test := range tests {
		if i < 0 {
			continue
		}
		t.Run(test.name, func(t *testing.T) {
			res := isMatch2(test.ins, test.inp)
			if res != test.out {
				t.Errorf("res %v, want %v", res, test.out)
			}
		})
	}
}
