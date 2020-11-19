package main

import (
	"testing"
)

func TestA(t *testing.T) {
	tests := []struct {
		name string
		in   *ListNode
		k    int
		out  *ListNode
	}{
		{
			name: "1",
			in:   &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			k:    3,
			out:  &ListNode{3, &ListNode{2, &ListNode{1, &ListNode{4, &ListNode{5, nil}}}}},
		},
		{
			name: "2",
			in:   &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			k:    2,
			out:  &ListNode{2, &ListNode{1, &ListNode{4, &ListNode{3, &ListNode{5, nil}}}}},
		},
		{
			name: "3",
			in:   &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			k:    1,
			out:  &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
		},
	}
	for i, test := range tests {
		if i < 0 {
			continue
		}
		t.Run(test.name, func(t *testing.T) {
			res := reverseKGroup(test.in, test.k)
			for node1, node2 := test.out, res; node1 != nil && node2 != nil; node1, node2 = node1.Next, node2.Next {
				t.Errorf("want %d, res %d\n", node1.Val, node2.Val)
			}
			// if reflect.DeepEqual(res, test.out) {
			// t.Errorf("res %+v, want %+v", res, test.out)
			// }
		})
	}
}
