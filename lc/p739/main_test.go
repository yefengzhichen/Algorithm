package main

import (
	"reflect"
	"testing"
)

func TestA(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		out  []int
	}{
		{
			name: "1",
			in:   []int{73, 74, 75, 71, 69, 72, 76, 73},
			out:  []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
	}
	for i, test := range tests {
		if i < 0 {
			continue
		}
		t.Run(test.name, func(t *testing.T) {
			res := dailyTemperatures(test.in)
			if reflect.DeepEqual(res, test.out) {
				t.Errorf("res %d, want %d", res, test.out)
			}
		})
	}
}
