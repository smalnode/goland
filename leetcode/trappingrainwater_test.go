package leetcode

import "testing"

func TestTrappingWater(t *testing.T) {
	testcases := []struct {
		height []int
		w      int
	}{
		{
			[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			6,
		},
		{
			[]int{0, 0},
			0,
		},
		{
			[]int{0, 0, 0, 0, 0, 0},
			0,
		},
		{
			[]int{10, 10, 10, 10, 10, 10},
			0,
		},
		{
			[]int{10, 8, 8, 9, 11, 10},
			5,
		},
	}

	for i, c := range testcases {
		got := trap(c.height)
		if got != c.w {
			t.Errorf("#%d height=%v, got %d, expected %d ", i, c.height, got, c.w)
		}
	}
}
