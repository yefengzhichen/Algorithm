package main

import (
	"testing"
)

func TestA(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		out  int
	}{
		{
			name: "1",
			in:   []int{1, 4, 2, 1, 3},
			out:  3,
		},
		{
			name: "2",
			in:   []int{1, 2, 5, 3, 4},
			out:  4,
		},
		{
			name: "3",
			in:   []int{395, 132, 276, 31, 612, 103, 209, 105, 214, 541, 454, 87, 600, 385, 345, 393, 45, 154, 70, 101, 468, 496, 253, 181, 162, 605, 425, 588, 74, 261, 155, 58, 549, 378, 535, 217, 213, 35, 564, 204, 193, 301, 78, 470, 140, 566, 315, 162, 471, 80, 451, 208, 402, 80, 224, 375, 279, 567, 272, 39, 495, 622, 256, 396, 452, 141, 344, 586, 310, 506, 348, 481, 388, 599, 412, 105, 75, 338, 71, 149, 19, 317, 23, 8, 592, 452, 624, 395, 412, 12, 303, 207, 491, 466, 238, 94, 538, 478, 163, 624, 308, 271, 18, 417, 209, 83, 18, 113, 169, 521, 539, 242, 36, 180, 429, 360, 203, 164, 580, 198, 98, 119, 157, 249, 609, 93, 323, 592, 105, 573, 243, 132, 25, 208, 505, 141, 454, 83, 199, 279, 464, 96, 285, 239, 24, 299, 484, 562, 410, 285, 421, 280, 63, 288, 502, 503, 55, 615, 395, 115, 560, 218, 165, 224, 536, 556, 201, 573, 167, 248, 541, 539, 35, 112, 56, 326, 138, 362, 91, 14, 531, 539, 291, 497, 570, 171, 615, 318, 586, 354, 462, 31, 199, 297, 589, 86, 257, 618, 591, 59, 532, 199, 302, 195, 587, 51, 87, 504, 62, 403, 513, 33, 86, 166, 576, 51, 201, 254, 343, 422, 388, 604, 305, 511, 388, 403, 564, 534, 466, 423, 42, 92, 146, 435, 613, 92, 239, 455, 614, 332, 176, 218, 60, 432, 584, 205, 323, 170, 320},
			out:  31,
		},
		{
			name: "4", //测试前面已有最大
			in:   []int{1, 2, 3, 4, 5, 2},
			out:  5,
		},
		{
			name: "5", //测试前面已有最大
			in:   []int{0, 9, 10, 2, 11},
			out:  4,
		},
		{
			name: "6",
			in:   []int{2, 2},
			out:  1,
		},
	}
	for i, test := range tests {
		if i < 0 {
			continue
		}
		t.Run(test.name, func(t *testing.T) {
			res := lengthOfLIS(test.in)
			if res != test.out {
				t.Errorf("res %d, want %d", res, test.out)
			}
		})
	}
}
