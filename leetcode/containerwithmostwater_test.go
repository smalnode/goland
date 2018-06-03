package leetcode

import "testing"

func TestMaxArea(t *testing.T) {
	testcases := []struct {
		height []int
		area   int
	}{
		{
			[]int{1, 2},
			1,
		},
		{
			[]int{1, 2, 1, 3, 10, 10, 2, 1},
			10,
		},
		{
			[]int{1, 2, 1, 3, 10, 10, 3, 2, 1},
			12,
		},
		{
			[]int{0, 0},
			0,
		},
		{
			[]int{0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			9,
		},
	}
	for i, c := range testcases {
		got := maxArea(c.height)
		if got != c.area {
			t.Errorf("#%d height=%v, got %d, expected %d ", i, c.height, got, c.area)
		}
	}
}
